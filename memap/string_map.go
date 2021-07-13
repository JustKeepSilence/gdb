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

type StringMemConcurrentMap []*StringMemConcurrentMapShared

type RmHisDbStringItems struct {
	GroupName  string
	TimeStamps []int32  // timeStamps
	Values     []string // values
}

type RmHisDbStringItem struct {
	TimeStamp int32
	Value     string
}

// StringMemConcurrentMapShared A "thread" safe string to anything map.
type StringMemConcurrentMapShared struct {
	items        map[string]RmHisDbStringItems
	sync.RWMutex // Read Write mutex, guards access to internal map.
}

func NewString() StringMemConcurrentMap {
	m := make(StringMemConcurrentMap, shardCount)
	for i := 0; i < shardCount; i++ {
		m[i] = &StringMemConcurrentMapShared{items: make(map[string]RmHisDbStringItems)}
	}
	return m
}

// getShard returns shard under given key
func (m StringMemConcurrentMap) getShard(key string) *StringMemConcurrentMapShared {
	return m[uint(fnv32(key))%uint(shardCount)]
}

// Upsert Insert or Update - updates existing element or inserts a new one using UpsertCb
func (m StringMemConcurrentMap) Upsert(key, groupName string, value RmHisDbStringItem) {
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
		tmp := RmHisDbStringItems{TimeStamps: []int32{value.TimeStamp}, Values: []string{value.Value}, GroupName: groupName}
		shard.items[key] = tmp
	}
}

// Remove removes an element from the map.
func (m StringMemConcurrentMap) Remove(key string) {
	// Try to get shard.
	shard := m.getShard(key)
	shard.Lock()
	delete(shard.items, key)
	shard.Unlock()
}

// StringTuple Used by the Iter & IterBuffered functions to wrap two variables together over a channel,
type StringTuple struct {
	Key string
	Val RmHisDbStringItems
}

// IterBuffered returns a buffered iterator which could be used in a for range loop.
func (m StringMemConcurrentMap) IterBuffered() <-chan StringTuple {
	chans := floatSnapshot(m)
	total := 0
	for _, c := range chans {
		total += cap(c)
	}
	ch := make(chan StringTuple, total)
	go fanStringIn(chans, ch)
	return ch
}

// Clear removes all items from map.
func (m StringMemConcurrentMap) Clear() {
	for item := range m.IterBuffered() {
		m.Remove(item.Key)
	}
}

// Returns a array of channels that contains elements in each shard,
// which likely takes a floatSnapshot of `m`.
// It returns once the size of each buffered channel is determined,
// before all the channels are populated using goroutines.
func floatSnapshot(m StringMemConcurrentMap) (chans []chan StringTuple) {
	chans = make([]chan StringTuple, shardCount)
	wg := sync.WaitGroup{}
	wg.Add(shardCount)
	// Foreach shard.
	for index, shard := range m {
		go func(index int, shard *StringMemConcurrentMapShared) {
			// Foreach key, value pair.
			shard.RLock()
			chans[index] = make(chan StringTuple, len(shard.items))
			wg.Done()
			for key, val := range shard.items {
				chans[index] <- StringTuple{key, val}
			}
			shard.RUnlock()
			close(chans[index])
		}(index, shard)
	}
	wg.Wait()
	return chans
}

// fanStringIn reads elements from channels `chans` into channel `out`
func fanStringIn(chans []chan StringTuple, out chan StringTuple) {
	wg := sync.WaitGroup{}
	wg.Add(len(chans))
	for _, ch := range chans {
		go func(ch chan StringTuple) {
			for t := range ch {
				out <- t
			}
			wg.Done()
		}(ch)
	}
	wg.Wait()
	close(out)
}

func (m StringMemConcurrentMap) Get(key string) (RmHisDbStringItems, bool) {
	shard := m.getShard(key)
	shard.RLock()
	// Get item from shard.
	val, ok := shard.items[key]
	shard.RUnlock()
	return val, ok
}

func (m StringMemConcurrentMap) Set(key string, v RmHisDbStringItems) {
	shard := m.getShard(key)
	shard.Lock()
	defer shard.Unlock()
	shard.items[key] = v
}
