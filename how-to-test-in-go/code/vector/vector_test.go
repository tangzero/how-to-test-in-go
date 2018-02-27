package vector_test

import (
	"testing"

	"github.com/tangzero/vector"
)

func TestNewVector(t *testing.T) {
	x, y := 12.0, 34.0
	v := vector.New(x, y)

	if v.X != x {
		t.Errorf("v.X should be %f", x)
	}
	if v.Y != y {
		t.Errorf("v.Y should be %f", y)
	}
}

func TestVectorLen(t *testing.T) {
	length := 5.0
	x, y := 3.0, 4.0
	v := vector.New(x, y)

	if v.Len() != length {
		t.Errorf("v.Len() should be %f", length)
	}
}

func TestVectorAdd(t *testing.T) {
	a := vector.New(10, 20)
	b := vector.New(40, 60)
	c := a.Add(b)

	if c.X != 50 {
		t.Errorf("v.X should be 50")
	}
	if c.Y != 80 {
		t.Errorf("v.Y should be 80")
	}
}

func TestVectorSub(t *testing.T) {
	a := vector.New(10, 20)
	b := vector.New(40, 60)
	c := a.Sub(b)

	if c.X != -30 {
		t.Errorf("v.X should be -30")
	}
	if c.Y != -40 {
		t.Errorf("v.Y should be -40")
	}
}
