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

type IntMemConcurrentMap []*IntMemConcurrentMapShared

type RmHisDbIntItems struct {
	GroupName  string
	TimeStamps []int32 // timeStamps
	Values     []int32 // values
}

type RmHisDbIntItem struct {
	TimeStamp int32
	Value     int32
}

// IntMemConcurrentMapShared A "thread" safe string to anything map.
type IntMemConcurrentMapShared struct {
	items        map[string]RmHisDbIntItems
	sync.RWMutex // Read Write mutex, guards access to internal map.
}

func NewInt() IntMemConcurrentMap {
	m := make(IntMemConcurrentMap, shardCount)
	for i := 0; i < shardCount; i++ {
		m[i] = &IntMemConcurrentMapShared{items: make(map[string]RmHisDbIntItems)}
	}
	return m
}

// getShard returns shard under given key
func (m IntMemConcurrentMap) getShard(key string) *IntMemConcurrentMapShared {
	return m[uint(fnv32(key))%uint(shardCount)]
}

// Upsert Insert or Update - updates existing element or inserts a new one using UpsertCb
func (m IntMemConcurrentMap) Upsert(key, groupName string, value RmHisDbIntItem) {
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
		tmp := RmHisDbIntItems{TimeStamps: []int32{value.TimeStamp}, Values: []int32{value.Value}, GroupName: groupName}
		shard.items[key] = tmp
	}
}

// Remove removes an element from the map.
func (m IntMemConcurrentMap) Remove(key string) {
	// Try to get shard.
	shard := m.getShard(key)
	shard.Lock()
	delete(shard.items, key)
	shard.Unlock()
}

// IntTuple Used by the Iter & IterBuffered functions to wrap two variables together over a channel,
type IntTuple struct {
	Key string
	Val RmHisDbIntItems
}

// IterBuffered returns a buffered iterator which could be used in a for range loop.
func (m IntMemConcurrentMap) IterBuffered() <-chan IntTuple {
	chans := intSnapshot(m)
	total := 0
	for _, c := range chans {
		total += cap(c)
	}
	ch := make(chan IntTuple, total)
	go fanIntIn(chans, ch)
	return ch
}

// Clear removes all items from map.
func (m IntMemConcurrentMap) Clear() {
	for item := range m.IterBuffered() {
		m.Remove(item.Key)
	}
}

// Returns a array of channels that contains elements in each shard,
// which likely takes a snapshot of `m`.
// It returns once the size of each buffered channel is determined,
// before all the channels are populated using goroutines.
func intSnapshot(m IntMemConcurrentMap) (chans []chan IntTuple) {
	chans = make([]chan IntTuple, shardCount)
	wg := sync.WaitGroup{}
	wg.Add(shardCount)
	// Foreach shard.
	for index, shard := range m {
		go func(index int, shard *IntMemConcurrentMapShared) {
			// Foreach key, value pair.
			shard.RLock()
			chans[index] = make(chan IntTuple, len(shard.items))
			wg.Done()
			for key, val := range shard.items {
				chans[index] <- IntTuple{key, val}
			}
			shard.RUnlock()
			close(chans[index])
		}(index, shard)
	}
	wg.Wait()
	return chans
}

// fanIntIn reads elements from channels `chans` into channel `out`
func fanIntIn(chans []chan IntTuple, out chan IntTuple) {
	wg := sync.WaitGroup{}
	wg.Add(len(chans))
	for _, ch := range chans {
		go func(ch chan IntTuple) {
			for t := range ch {
				out <- t
			}
			wg.Done()
		}(ch)
	}
	wg.Wait()
	close(out)
}

func (m IntMemConcurrentMap) Get(key string) (RmHisDbIntItems, bool) {
	shard := m.getShard(key)
	shard.RLock()
	// Get item from shard.
	val, ok := shard.items[key]
	shard.RUnlock()
	return val, ok
}

func (m IntMemConcurrentMap) Set(key string, v RmHisDbIntItems) {
	shard := m.getShard(key)
	shard.Lock()
	defer shard.Unlock()
	shard.items[key] = v
}
