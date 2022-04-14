package bloomfilter

import (
	"github.com/twmb/murmur3"
)

// SimpleBloomFilter is a filter that uses only a single murmur3 hash function.
type SimpleBloomfilter struct {
	Bit_array []bool
	Size      int
}

func NewSimpleBloomfilter(size int) SimpleBloomfilter {
	arr := make([]bool, size)
	return SimpleBloomfilter{
		Bit_array: arr,
		Size:      size,
	}
}

func (bf *SimpleBloomfilter) Check(s string) bool {
	h := bf.Hash(s)
	return bf.Bit_array[h]
}

func (bf *SimpleBloomfilter) Insert(s string) {
	h := bf.Hash(s)
	bf.Bit_array[h] = true
}

func (bf *SimpleBloomfilter) Hash(s string) uint32 {
	return murmur3.Sum32([]byte(s)) % uint32(bf.Size)
}
