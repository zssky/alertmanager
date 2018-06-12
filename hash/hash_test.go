package hash

import (
	"testing"
)

func TestUintHashOrderIndependent(t *testing.T) {
	a := []uint64{1, 2, 3, 4, 5, 6}
	b := []uint64{3, 4, 2, 5, 1, 6}
	h1 := Uint64(a)
	h2 := Uint64(b)
	if h1 != h2 {
		t.Fatalf("hashing is order independent: %d != %d", h1, h2)
	}
}
