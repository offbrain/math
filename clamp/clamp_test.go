package clamp_test

import (
	"image"
	"testing"

	"github.com/offbrain/math/clamp"
)

func TestFloat32(t *testing.T) {
	if clamp.Float32(5, 1, 10) != 5 {
		t.Error("Incorrect result for number in range")
	}
	if clamp.Float32(-1, 1, 10) != 1 {
		t.Error("Incorrect result for number below range")
	}
	if clamp.Float32(11, 1, 10) != 10 {
		t.Error("Incorrect result for number above range")
	}
	if clamp.Float32(-50, -20, -15) != -20 {
		t.Error("Incorrect result for negative number")
	}
}

func TestFloat64(t *testing.T) {
	if clamp.Float64(5, 1, 10) != 5 {
		t.Error("Incorrect result for number in range")
	}
	if clamp.Float64(-1, 1, 10) != 1 {
		t.Error("Incorrect result for number below range")
	}
	if clamp.Float64(11, 1, 10) != 10 {
		t.Error("Incorrect result for number above range")
	}
	if clamp.Float64(-50, -20, -15) != -20 {
		t.Error("Incorrect result for negative number")
	}
}

func TestInt(t *testing.T) {
	if clamp.Int(5, 1, 10) != 5 {
		t.Error("Incorrect result for number in range")
	}
	if clamp.Int(-1, 1, 10) != 1 {
		t.Error("Incorrect result for number below range")
	}
	if clamp.Int(11, 1, 10) != 10 {
		t.Error("Incorrect result for number above range")
	}
	if clamp.Int(-50, -20, -15) != -20 {
		t.Error("Incorrect result for negative number")
	}
}

func TestUint32(t *testing.T) {
	if clamp.Uint32(5, 1, 10) != 5 {
		t.Error("Incorrect result for number in range")
	}
	if clamp.Uint32(0, 1, 10) != 1 {
		t.Error("Incorrect result for number below range")
	}
	if clamp.Uint32(11, 1, 10) != 10 {
		t.Error("Incorrect result for number above range")
	}
}

func TestUint64(t *testing.T) {
	if clamp.Uint64(5, 1, 10) != 5 {
		t.Error("Incorrect result for number in range")
	}
	if clamp.Uint64(0, 1, 10) != 1 {
		t.Error("Incorrect result for number below range")
	}
	if clamp.Uint64(11, 1, 10) != 10 {
		t.Error("Incorrect result for number above range")
	}
}

func TestPoint(t *testing.T) {
	p := image.Pt(5, 6)
	r := image.Rect(1, 2, 10, 11)
	q := clamp.Point(p, r)
	if q.X != p.X || q.Y != p.Y {
		t.Error("Incorrect result for number in range")
	}
	p = image.Pt(0, 1)
	q = clamp.Point(p, r)
	if q.X != r.Min.X || q.Y != r.Min.Y {
		t.Error("Incorrect result for number below range")
	}
	p = image.Pt(12, 14)
	q = clamp.Point(p, r)
	if q.X != r.Max.X-1 || q.Y != r.Max.Y-1 {
		t.Error("Incorrect result for number above range")
	}
}
