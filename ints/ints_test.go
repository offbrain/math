package ints_test

import (
	"testing"

	"github.com/offbrain/math/ints"
)

var power2Tests = []struct {
	n        int // input
	expected int // expected result
}{
	{3, 4},
	{5, 8},
	{9, 16},
	{17, 32},
	{33, 64},
}

func TestRoundUpPowerOf2(t *testing.T) {
	for _, tt := range power2Tests {
		actual := ints.RoundUpPowerOf2(tt.n)
		if actual != tt.expected {
			t.Errorf("RoundUpPowerOf2(%d): expected %d, actual %d", tt.n, tt.expected, actual)
		}
	}
}
