//This is free and unencumbered software released into the public domain.

//Package ints implements integer utility functions
package ints

//RoundUpPowerOf2 return the rounded up power of 2 of n
//from: https://graphics.stanford.edu/~seander/bithacks.html#RoundUpPowerOf2
func RoundUpPowerOf2(n int) int {
	n--
	n |= n >> 1
	n |= n >> 2
	n |= n >> 4
	n |= n >> 8
	n |= n >> 16
	n++
	return n
}
