package pgen

import (
	"math"
	"math/rand"
)

// A Gen is a deterministic generator of random-like numbers.
type Gen int64

func New(i int64) Gen {
	return Gen(i)
}

// Int64 generates the random number for index i. It returns integers in the range [0,2^63-1].
func (g Gen) Int64(i int64) int64 {
	if i < 0 {
		panic("invalid argument to Int64")
	}

	v := g.Uint64(i)
	// XOR-fold to 63 bits
	return int64((v >> 63) ^ (v & ((1 << 63) - 1)))
}

// Int32 generates the random number for index i. It returns integers in the range [0,2^31-1].
func (g Gen) Int32(i int64) int32 {
	if i < 0 {
		panic("invalid argument to Int64")
	}

	v := g.Uint64(i)
	// XOR-fold to 31 bits
	return int32((v >> 31) ^ (v & ((1 << 31) - 1)))
}

// Int64 generates the random number for index i and converts it to a float64 in the range [0,1].
func (g Gen) Float64(i int64) float64 {
	if i < 0 {
		panic("invalid argument to Float64")
	}
	return float64(g.Int64(i)) / math.MaxInt64
}

// Intn generates the random number for index i. It returns integers in the range [0,n).
func (g Gen) Intn(i int64, n int) int {
	if i <= 0 {
		panic("invalid argument to Int64")
	}

	un := uint64(n)
	mask := umask(un)
	v := g.Uint64(i)

	for v&mask >= un {
		v = hash(v, un)
	}
	return int(v & mask)
}

func (g Gen) Rand(i int64) *rand.Rand {
	seed := g.Int64(i)
	return rand.New(rand.NewSource(seed))
}

func (g Gen) Uint64(i int64) uint64 {
	s := hash(offset64, uint64(g))
	return hash(s, uint64(i))
}

// hash hashes the octets of v into s using the FNV-1a hash algorithm
// TODO: investigate http://xoroshiro.di.unimi.it/
func hash(s uint64, v uint64) uint64 {
	s ^= uint64(v & 0xff)
	s *= prime64

	s ^= uint64((v >> 8) & 0xff)
	s *= prime64

	s ^= uint64((v >> 16) & 0xff)
	s *= prime64

	s ^= uint64((v >> 24) & 0xff)
	s *= prime64

	s ^= uint64((v >> 32) & 0xff)
	s *= prime64

	s ^= uint64((v >> 40) & 0xff)
	s *= prime64

	s ^= uint64((v >> 48) & 0xff)
	s *= prime64

	s ^= uint64(v >> 56)
	s *= prime64

	return s
}

const (
	offset64 = 14695981039346656037
	prime64  = 1099511628211
)

// umask creates a mask equal to one less than the next power of 2 above n
func umask(n uint64) uint64 {
	mask := uint64((1 << 64) - 1)
	t := uint64(1 << 63)
	for j := 0; j < 64; j++ {
		if n&t == t {
			break
		}
		mask ^= t
		t >>= 1
	}
	return mask
}
