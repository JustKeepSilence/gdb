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

type BoolMemConcurrentMap []*BoolMemConcurrentMapShared

type RmHisDbBoolItems struct {
	GroupName  string
	TimeStamps []int32 // timeStamps
	Values     []bool  // values
}

type RmHisDbBoolItem struct {
	TimeStamp int32
	Value     bool
}

// BoolMemConcurrentMapShared A "thread" safe string to anything map.
type BoolMemConcurrentMapShared struct {
	items        map[string]RmHisDbBoolItems
	sync.RWMutex // Read Write mutex, guards access to internal map.
}

func NewBool() BoolMemConcurrentMap {
	m := make(BoolMemConcurrentMap, shardCount)
	for i := 0; i < shardCount; i++ {
		m[i] = &BoolMemConcurrentMapShared{items: make(map[string]RmHisDbBoolItems)}
	}
	return m
}

// getShard returns shard under given key
func (m BoolMemConcurrentMap) getShard(key string) *BoolMemConcurrentMapShared {
	return m[uint(fnv32(key))%uint(shardCount)]
}

// Upsert Insert or Update - updates existing element or inserts a new one using UpsertCb
func (m BoolMemConcurrentMap) Upsert(key, groupName string, value RmHisDbBoolItem) {
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
		tmp := RmHisDbBoolItems{TimeStamps: []int32{value.TimeStamp}, Values: []bool{value.Value}, GroupName: groupName}
		shard.items[key] = tmp
	}
}

// Remove removes an element from the map.
func (m BoolMemConcurrentMap) Remove(key string) {
	// Try to get shard.
	shard := m.getShard(key)
	shard.Lock()
	delete(shard.items, key)
	shard.Unlock()
}

// BoolTuple Used by the Iter & IterBuffered functions to wrap two variables together over a channel,
type BoolTuple struct {
	Key string
	Val RmHisDbBoolItems
}

// IterBuffered returns a buffered iterator which could be used in a for range loop.
func (m BoolMemConcurrentMap) IterBuffered() <-chan BoolTuple {
	chans := boolSnapshot(m)
	total := 0
	for _, c := range chans {
		total += cap(c)
	}
	ch := make(chan BoolTuple, total)
	go fanBoolIn(chans, ch)
	return ch
}

// Clear removes all items from map.
func (m BoolMemConcurrentMap) Clear() {
	for item := range m.IterBuffered() {
		m.Remove(item.Key)
	}
}

// Returns a array of channels that contains elements in each shard,
// which likely takes a boolSnapshot of `m`.
// It returns once the size of each buffered channel is determined,
// before all the channels are populated using goroutines.
func boolSnapshot(m BoolMemConcurrentMap) (chans []chan BoolTuple) {
	chans = make([]chan BoolTuple, shardCount)
	wg := sync.WaitGroup{}
	wg.Add(shardCount)
	// Foreach shard.
	for index, shard := range m {
		go func(index int, shard *BoolMemConcurrentMapShared) {
			// Foreach key, value pair.
			shard.RLock()
			chans[index] = make(chan BoolTuple, len(shard.items))
			wg.Done()
			for key, val := range shard.items {
				chans[index] <- BoolTuple{key, val}
			}
			shard.RUnlock()
			close(chans[index])
		}(index, shard)
	}
	wg.Wait()
	return chans
}

// fanBoolIn reads elements from channels `chans` into channel `out`
func fanBoolIn(chans []chan BoolTuple, out chan BoolTuple) {
	wg := sync.WaitGroup{}
	wg.Add(len(chans))
	for _, ch := range chans {
		go func(ch chan BoolTuple) {
			for t := range ch {
				out <- t
			}
			wg.Done()
		}(ch)
	}
	wg.Wait()
	close(out)
}

func (m BoolMemConcurrentMap) Get(key string) (RmHisDbBoolItems, bool) {
	shard := m.getShard(key)
	shard.RLock()
	// Get item from shard.
	val, ok := shard.items[key]
	shard.RUnlock()
	return val, ok
}

func (m BoolMemConcurrentMap) Set(key string, v RmHisDbBoolItems) {
	shard := m.getShard(key)
	shard.Lock()
	defer shard.Unlock()
	shard.items[key] = v
}
