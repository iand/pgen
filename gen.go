/*
  This is free and unencumbered software released into the public domain. For more
  information, see <http://unlicense.org/> or the accompanying UNLICENSE file.
*/

// A package for deterministic generation of random-like numbers.
package pgen

import (
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

// Float64 generates the random number for index i and converts it to a float64 in the range [0,1).
func (g Gen) Float64(i int64) float64 {
	if i < 0 {
		panic("invalid argument to Float64")
	}
	// There are only 2^53 representable floating point numbers in [0,1] so to get an unbiased
	// distribution we need to pick one number from 2^53 and scale it to the range 0,1.
	// See http://lemire.me/blog/2017/02/28/how-many-floating-point-numbers-are-in-the-interval-01/
	// for interesting discussion of this phenomena.
	return float64(g.Intn(i, 1<<53)) / (1 << 53)
}

// Float32 generates the random number for index i and converts it to a float32 in the range [0,1).
func (g Gen) Float32(i int64) float32 {
	if i < 0 {
		panic("invalid argument to Float32")
	}
	return float32(g.Intn(i, 1<<24)) / (1 << 24)
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

// Rand returns a random number generator for index i.
func (g Gen) Rand(i int64) *rand.Rand {
	seed := g.Int64(i)
	return rand.New(rand.NewSource(seed))
}

// Uint64 generates the random number for index i. It returns integers in the range [0,2^64-1].
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
