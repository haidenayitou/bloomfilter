package bloomfilter

import (
	"testing"
)

func TestBasic(t *testing.T) {
	bf := NewBloomFilter(3, 100)
	d1, d2 := []byte("hello"), []byte("jello")
	bf.Add(d1)
	bf.Add(d2)
	if !bf.Check(d1) {
		t.Errorf("d1 should be present in the BloomFilter")
	}
	if !bf.Check(d2) {
		t.Errorf("d2 should be present in the BloomFilter")
	}
}
