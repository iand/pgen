/*
  This is free and unencumbered software released into the public domain. For more
  information, see <http://unlicense.org/> or the accompanying UNLICENSE file.
*/

package pgen

import (
	"testing"
)

func TestUmask(t *testing.T) {
	testCases := []struct {
		n, mask uint64
	}{
		{n: 0, mask: 0},
		{n: 22, mask: 31},
		{n: 122112, mask: 131071},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			mask := umask(tc.n)
			if mask != tc.mask {
				t.Errorf("got %b (%[1]d), wanted %b (%[2]d)", mask, tc.mask)
			}
		})
	}
}
