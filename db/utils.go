/*
createTime: 2021/6/7
creator: JustKeepSilence
github: https://github.com/JustKeepSilence
goVersion: 1.16
*/

package db

import (
	"github.com/shopspring/decimal"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"unsafe"
)

// common functions for gdb

// convert strings to byte to avoid mem alloc
// note: return byte should not be changed
func convertStringToByte(s string) []byte {
	var b []byte
	strHeader := (*reflect.StringHeader)(unsafe.Pointer(&s))
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sliceHeader.Data = strHeader.Data
	sliceHeader.Len = strHeader.Len
	sliceHeader.Cap = strHeader.Len
	return b
}

// convert string slice to interface slice
func convertStringToInterface(s ...string) []interface{} {
	v := make([]interface{}, len(s))
	for _, s2 := range s {
		v = append(v, s2)
	}
	return v
}

// get directory size of given path
func dirSize(path string) (float64, error) {
	var size int64
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			size += 0
		} else {
			if !info.IsDir() {
				size += info.Size()
			}
		}
		return nil
	})
	r, _ := decimal.NewFromFloat(float64(size) / (1024 * 1024 * 1024)).Round(2).Float64()
	return r, err
}

// get startTime index of st in a, include index, you can use ts[index:] to get slice of ts
//
// if st <= a[0] ==> 0
//
// if st > a[len(a) - 1] ==> -1, that is not included st in slice a
//
// if st == a[len(a) - 1] ==> len(a) - 1
//
// if a[0]...st...a[len(a) - 1]  ==> sort.Search(a[i]>= st) ==> index
func getStartIndex(st int32, a []int32) int {
	if st <= a[0] {
		return 0 // st    a[0]......a[len(a) - 1]
	}
	if st == a[len(a)-1] {
		return len(a) - 1
	}
	if st > a[len(a)-1] {
		return -1 // a[0]......a[len(a) - 1]   st
	}
	return sort.Search(len(a), func(i int) bool {
		return a[i] >= st // a[0], a[1], st, a[2], a[3]....a[len(a) - 1]
	})
}

// get endTime index of et in a, not include index, you can use ts[:index] to get slice of ts
//
// if et <= a[0] ==> -1, that is not included et in slice a
//
// if et == a[len(a) - 1] ==> slice index is len(a) - 2 ==> len(a) - 1
//
// if et > a[len(a) - 1]  ==> len(a)
//
// a[0]...et....a[len(a) - 1] ==> sort.Search(et <= a[i]) ==> index
func getEndIndex(et int32, a []int32) int {
	if et <= a[0] {
		return -1 // not include et, et a[0].....a[len(a) - 1]
	}
	if et == a[len(a)-1] {
		return len(a) - 1 // index is len(a) - 2
	}
	if et > a[len(a)-1] {
		return len(a) // a[0].....a[len(a) - 1], et, index is len(a) - 1
	}
	return sort.Search(len(a), func(i int) bool {
		return et <= a[i]
	})
}

func getMiddleIndex(x int32, a []int32) (result int) {
	mid := len(a) / 2
	switch {
	case len(a) == 0:
		result = -1
	case a[mid] > x:
		result = getMiddleIndex(x, a[:mid])
	case a[mid] < x:
		result = getMiddleIndex(x, a[mid+1:])
		if result != -1 {
			result += mid + 1
		}
	default:
		result = mid
	}
	return
}

// get startIndex and endIndex of st and et corresponding to slice a
//
// you can use a[startIndex:endIndex] to get slice data of ts a
//
// if st == -1 or et == -1 , that is a not include [st, et)
func getIndex(st, et int32, a []int32) (int, int) {
	return getStartIndex(st, a), getEndIndex(et, a)
}
