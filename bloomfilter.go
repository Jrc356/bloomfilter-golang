package bloomfilter

import (
	"github.com/hoangtq219/murmur3"
)

type Bloomfilter struct {
	bit_array []bool
	size      int
}

func (bf Bloomfilter) New(arr_size int) Bloomfilter {
	var arr []bool
	return Bloomfilter{
		bit_array: arr,
		size:      arr_size,
	}
}

func (bf *Bloomfilter) Check(s string) bool {
	h := bf.Hash(s)
	return bf.bit_array[h]
}

func (bf *Bloomfilter) Insert(s string) {
	h := bf.Hash(s)
	bf.bit_array[h] = true
}

func (bf *Bloomfilter) Hash(s string) int {
	// TODO: change seed to something variable?
	// TODO: perhaps not the hashing function to use
	// or at least need to think about how to make the hash non negative without clashing much
	// TLDR; more research needed
	return murmur3.HashString(-1, s).AsInt() % bf.size
}
