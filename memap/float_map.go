/*
createTime: 2021/6/20
creator: JustKeepSilence
github: https://github.com/JustKeepSilence
goVersion: 1.16
*/

package memap

import (
	"sync"
)

type FloatMemConcurrentMap []*FloatMemConcurrentMapShared

type RmHisDbFloatItems struct {
	GroupName  string
	TimeStamps []int32   // timeStamps
	Values     []float32 // values
}

type RmHisDbFloatItem struct {
	TimeStamp int32
	Value     float32
}

// FloatMemConcurrentMapShared A "thread" safe string to anything map.
type FloatMemConcurrentMapShared struct {
	items        map[string]RmHisDbFloatItems
	sync.RWMutex // Read Write mutex, guards access to internal map.
}

func NewFloat() FloatMemConcurrentMap {
	m := make(FloatMemConcurrentMap, shardCount)
	for i := 0; i < shardCount; i++ {
		m[i] = &FloatMemConcurrentMapShared{items: make(map[string]RmHisDbFloatItems)}
	}
	return m
}

// getShard returns shard under given key
func (m FloatMemConcurrentMap) getShard(key string) *FloatMemConcurrentMapShared {
	return m[uint(fnv32(key))%uint(shardCount)]
}

// Upsert Insert or Update - updates existing element or inserts a new one using UpsertCb
func (m FloatMemConcurrentMap) Upsert(key, groupName string, value RmHisDbFloatItem) {
	shard := m.getShard(key)
	shard.Lock()
	defer shard.Unlock()
	v, ok := shard.items[key]
	if ok {
		// key exist
		v.Values = append(v.Values, value.Value)
		v.TimeStamps = append(v.TimeStamps, value.TimeStamp)
		shard.items[key] = v
	} else {
		tmp := RmHisDbFloatItems{TimeStamps: []int32{value.TimeStamp}, Values: []float32{value.Value}, GroupName: groupName}
		shard.items[key] = tmp
	}
}

// Remove removes an element from the map.
func (m FloatMemConcurrentMap) Remove(key string) {
	// Try to get shard.
	shard := m.getShard(key)
	shard.Lock()
	delete(shard.items, key)
	shard.Unlock()
}

// FloatTuple Used by the Iter & IterBuffered functions to wrap two variables together over a channel,
type FloatTuple struct {
	Key string
	Val RmHisDbFloatItems
}

// IterBuffered returns a buffered iterator which could be used in a for range loop.
func (m FloatMemConcurrentMap) IterBuffered() <-chan FloatTuple {
	chans := snapshot(m)
	total := 0
	for _, c := range chans {
		total += cap(c)
	}
	ch := make(chan FloatTuple, total)
	go fanFloatIn(chans, ch)
	return ch
}

// Clear removes all items from map.
func (m FloatMemConcurrentMap) Clear() {
	for item := range m.IterBuffered() {
		m.Remove(item.Key)
	}
}

// Returns a array of channels that contains elements in each shard,
// which likely takes a snapshot of `m`.
// It returns once the size of each buffered channel is determined,
// before all the channels are populated using goroutines.
func snapshot(m FloatMemConcurrentMap) (chans []chan FloatTuple) {
	chans = make([]chan FloatTuple, shardCount)
	wg := sync.WaitGroup{}
	wg.Add(shardCount)
	// Foreach shard.
	for index, shard := range m {
		go func(index int, shard *FloatMemConcurrentMapShared) {
			// Foreach key, value pair.
			shard.RLock()
			chans[index] = make(chan FloatTuple, len(shard.items))
			wg.Done()
			for key, val := range shard.items {
				chans[index] <- FloatTuple{key, val}
			}
			shard.RUnlock()
			close(chans[index])
		}(index, shard)
	}
	wg.Wait()
	return chans
}

// fanFloatIn reads elements from channels `chans` into channel `out`
func fanFloatIn(chans []chan FloatTuple, out chan FloatTuple) {
	wg := sync.WaitGroup{}
	wg.Add(len(chans))
	for _, ch := range chans {
		go func(ch chan FloatTuple) {
			for t := range ch {
				out <- t
			}
			wg.Done()
		}(ch)
	}
	wg.Wait()
	close(out)
}

func (m FloatMemConcurrentMap) Get(key string) (RmHisDbFloatItems, bool) {
	shard := m.getShard(key)
	shard.RLock()
	// Get item from shard.
	val, ok := shard.items[key]
	shard.RUnlock()
	return val, ok
}

func (m FloatMemConcurrentMap) Set(key string, v RmHisDbFloatItems) {
	shard := m.getShard(key)
	shard.Lock()
	defer shard.Unlock()
	shard.items[key] = v
}
