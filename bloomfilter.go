package bloomfilter

import (
	"hash"
	"hash/fnv"
	"math"
)

type BloomFilter struct {
	bitmap []bool
	k      int
	n      int
	m      int
	hashfn hash.Hash64
}

func NewBloomFilter(numHashFuncs, bfSize int) *BloomFilter {
	bf := new(BloomFilter)
	bf.bitmap = make([]bool, bfSize)
	bf.k, bf.m = numHashFuncs, bfSize
	bf.n = 0
	bf.hashfn = fnv.New64()
	return bf
}

func (bf *BloomFilter) getHash(b []byte) (uint32, uint32) {
	bf.hashfn.Reset()
	bf.hashfn.Write(b)
	hash64 := bf.hashfn.Sum64()
	h1 := uint32(hash64 & ((1 << 32) - 1))
	h2 := uint32(hash64 >> 32)
	return h1, h2
}

func (bf *BloomFilter) Add(e []byte) {
	h1, h2 := bf.getHash(e)
	for i := 0; i < bf.k; i++ {
		ind := (h1 + uint32(i)*h2) % uint32(bf.m)
		bf.bitmap[ind] = true
	}
	bf.n++
}

func (bf *BloomFilter) Check(x []byte) bool {
	h1, h2 := bf.getHash(x)
	result := true
	for i := 0; i < bf.k; i++ {
		ind := (h1 + uint32(i)*h2) % uint32(bf.m)
		result = result && bf.bitmap[ind]
	}
	return result
}

func (bf *BloomFilter) FalsePositiveRate() float64 {
	return math.Pow((1 - math.Exp(-float64(bf.k*bf.n)/
		float64(bf.m))), float64(bf.k))
}
