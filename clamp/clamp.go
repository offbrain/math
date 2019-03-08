//This is free and unencumbered software released into the public domain.

//Package clamp implements clamping methods
//Clamping is the process of limiting numbers to specific ranges
package clamp

import (
	"image"
)

//Float32 returns the nearest n in the range [min, max]
func Float32(n, min, max float32) float32 {
	if n < min {
		n = min
	} else if n > max {
		n = max
	}
	return n
}

//Float64 returns the nearest n in the range [min, max]
func Float64(n, min, max float64) float64 {
	if n < min {
		n = min
	} else if n > max {
		n = max
	}
	return n
}

//Int returns the nearest n in the range [min, max]
func Int(n, min, max int) int {
	if n < min {
		n = min
	} else if n > max {
		n = max
	}
	return n
}

//Uint32 returns the nearest n in the range [min, max]
func Uint32(n, min, max uint32) uint32 {
	if n < min {
		n = min
	} else if n > max {
		n = max
	}
	return n
}

//Uint64 returns the nearest n in the range [min, max]
func Uint64(n, min, max uint64) uint64 {
	if n < min {
		n = min
	} else if n > max {
		n = max
	}
	return n
}

//Point returns the nearest Point the Rectangle.
//The Rectangle contains the points with Min.X <= X < Max.X, Min.Y <= Y < Max.Y.
func Point(p image.Point, r image.Rectangle) image.Point {
	if p.X < r.Min.X {
		p.X = r.Min.X
	} else if p.X >= r.Max.X {
		p.X = r.Max.X - 1
	}
	if p.Y < r.Min.Y {
		p.Y = r.Min.Y
	} else if p.Y >= r.Max.Y {
		p.Y = r.Max.Y - 1
	}
	return p
}
