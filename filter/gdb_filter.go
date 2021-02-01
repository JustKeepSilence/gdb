/*
creatTime: 2021/2/1
creator: JustKeepSilence
github: https://github.com/JustKeepSilence
goVersion: 1.15.3
*/

package filter

// gdb counting bloom filter see: https://pkg.go.dev/github.com/syndtr/goleveldb@v1.0.0/leveldb/filter#FilterGenerator
// and https://github.com/tylertreat/BoomFilters

type gdbBloomFilter int

func (gdbBloomFilter) Name() string {
	return "gdbBloomFilter"
}

func (g gdbBloomFilter) Contains(filter, key []byte) bool {

}
