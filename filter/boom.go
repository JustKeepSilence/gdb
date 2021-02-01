/*
creatTime: 2021/2/1
creator: JustKeepSilence
github: https://github.com/JustKeepSilence
goVersion: 1.15.3
*/

package filter

import (
	"encoding/binary"
	"hash"
	"math"
)

const fillRatio = 0.5

// Filter is a probabilistic data structure which is used to test the
// membership of an element in a set.
type Filter interface {
	// Test will test for membership of the data and returns true if it is a
	// member, false if not.
	Test([]byte) bool

	// Add will add the data to the Bloom filter. It returns the filter to
	// allow for chaining.
	Add([]byte) Filter

	// TestAndAdd is equivalent to calling Test followed by Add. It returns
	// true if the data is a member, false if not.
	TestAndAdd([]byte) bool
}

// OptimalM calculates the optimal Bloom filter size, m, based on the number of
// items and the desired rate of false positives.
func OptimalM(n uint, fpRate float64) uint {
	return uint(math.Ceil(float64(n) / ((math.Log(fillRatio) *
		math.Log(1-fillRatio)) / math.Abs(math.Log(fpRate)))))
}

// OptimalK calculates the optimal number of hash functions to use for a Bloom
// filter based on the desired rate of false positives.
func OptimalK(fpRate float64) uint {
	return uint(math.Ceil(math.Log2(1 / fpRate)))
}

// hashKernel returns the upper and lower base hash values from which the k
// hashes are derived.
func hashKernel(data []byte, hash hash.Hash64) (uint32, uint32) {
	hash.Write(data)
	sum := hash.Sum(nil)
	hash.Reset()
	return binary.BigEndian.Uint32(sum[4:8]), binary.BigEndian.Uint32(sum[0:4])
}
