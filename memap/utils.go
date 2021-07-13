/*
createTime: 2021/6/20
creator: JustKeepSilence
github: https://github.com/JustKeepSilence
goVersion: 1.16
*/

package memap

var shardCount = 32

func fnv32(key string) uint32 {
	hash := uint32(2166136261)
	const prime32 = uint32(16777619)
	for i := 0; i < len(key); i++ {
		hash *= prime32
		hash ^= uint32(key[i])
	}
	return hash
}

func fnvInt32(key int32) uint32 {
	hash := uint32(2166136261)
	const prime32 = uint32(16777619)
	hash *= prime32
	hash ^= uint32(key)
	return hash
}
