package models

import "math"

type Point struct {
	X, Y int
}

func (p Point) Add(v Vector) Point {
	return Point{
		X: p.X + v.X,
		Y: p.Y + v.Y,
	}
}

type Vector struct {
	X, Y int
}

func (v Vector) Add(other Vector) Vector {
	return Vector{
		X: v.X + other.X,
		Y: v.Y + other.Y,
	}
}

type Angle struct {
	Degrees int
}

func (a Angle) Radians() float64 {
	return math.Pi * float64(a.Degrees) / 180.0
}

func (a Angle) Normalized() Angle {
	normalized := a.Degrees % 360
	if normalized < 0 {
		normalized += 360
	}
	return Angle{Degrees: normalized}
}

func (a Angle) Add(other interface{}) Angle {
	switch v := other.(type) {
	case Angle:
		return Angle{Degrees: a.Degrees + v.Degrees}
	case int:
		return Angle{Degrees: a.Degrees + v}
	default:
		panic("unsupported type for addition with Angle")
	}
}

func (a Angle) Sub(other interface{}) Angle {
	switch v := other.(type) {
	case Angle:
		return Angle{Degrees: a.Degrees - v.Degrees}
	case int:
		return Angle{Degrees: a.Degrees - v}
	default:
		panic("unsupported type for subtraction with Angle")
	}
}

func (a Angle) Equal(other interface{}) bool {
	switch v := other.(type) {
	case Angle:
		return a.Degrees == v.Degrees
	case int:
		return a.Degrees == v
	default:
		return false
	}
}
