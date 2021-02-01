/*
creatTime: 2021/2/1
creator: JustKeepSilence
github: https://github.com/JustKeepSilence
goVersion: 1.15.3
*/

package compare

import "bytes"

// gdb compare, which is consistent with byte_compare in gp leveldb
// the main function is to encrypt the database

type gdbBytesComparer struct{}

func (gdbBytesComparer) Compare(a, b []byte) int {
	return bytes.Compare(a, b)
}

func (gdbBytesComparer) Name() string {
	return "leveldb.GdbBytesCompare"
}

func (gdbBytesComparer) Separator(dst, a, b []byte) []byte {
	i, n := 0, len(a)
	if n > len(b) {
		n = len(b)
	}
	for ; i < n && a[i] == b[i]; i++ {
	}
	if i >= n {
		// Do not shorten if one string is a prefix of the other
	} else if c := a[i]; c < 0xff && c+1 < b[i] {
		dst = append(dst, a[:i+1]...)
		dst[len(dst)-1]++
		return dst
	}
	return nil
}

func (gdbBytesComparer) Successor(dst, b []byte) []byte {
	for i, c := range b {
		if c != 0xff {
			dst = append(dst, b[:i+1]...)
			dst[len(dst)-1]++
			return dst
		}
	}
	return nil
}

var GdbComparer = gdbBytesComparer{}
