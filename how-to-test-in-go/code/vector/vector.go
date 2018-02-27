package vector

import "math"

type Vector struct {
	X float64
	Y float64
}

func New(x float64, y float64) Vector {
	return Vector{X: x, Y: y}
}

func (v Vector) Len() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v Vector) Add(other Vector) Vector {
	return New(v.X+other.X, v.Y+other.Y)
}

func (v Vector) Sub(other Vector) Vector {
	return New(v.X-other.X, v.Y-other.Y)
}
