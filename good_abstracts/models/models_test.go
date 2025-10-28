package models

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPoint(t *testing.T) {
	t.Run("Point Add Vector", func(t *testing.T) {
		point := Point{X: 10, Y: 20}
		vector := Vector{X: 5, Y: -3}

		result := point.Add(vector)

		assert.Equal(t, Point{X: 15, Y: 17}, result)
	})
}

func TestVector(t *testing.T) {
	t.Run("Vector Add", func(t *testing.T) {
		v1 := Vector{X: 3, Y: 4}
		v2 := Vector{X: 1, Y: 2}

		result := v1.Add(v2)

		assert.Equal(t, Vector{X: 4, Y: 6}, result)
	})
}

func TestAngle(t *testing.T) {
	t.Run("Конвертируем в радианы", func(t *testing.T) {
		angle := Angle{Degrees: 180}

		radians := angle.Radians()

		assert.InEpsilon(t, math.Pi, radians, 0.0001)
	})

	t.Run("Нормализация положительного угла", func(t *testing.T) {
		angle := Angle{Degrees: 45}

		normalized := angle.Normalized()

		assert.Equal(t, Angle{Degrees: 45}, normalized)
	})

	t.Run("Нормализация", func(t *testing.T) {
		angle := Angle{Degrees: 450}

		normalized := angle.Normalized()

		assert.Equal(t, Angle{Degrees: 90}, normalized)
	})

	t.Run("Нормализация отрицательного угла", func(t *testing.T) {
		angle := Angle{Degrees: -90}

		normalized := angle.Normalized()

		assert.Equal(t, Angle{Degrees: 270}, normalized)
	})

	t.Run("Add Angle", func(t *testing.T) {
		angle1 := Angle{Degrees: 30}
		angle2 := Angle{Degrees: 15}

		result := angle1.Add(angle2)

		assert.Equal(t, Angle{Degrees: 45}, result)
	})

	t.Run("Add int", func(t *testing.T) {
		angle := Angle{Degrees: 30}

		result := angle.Add(20)

		assert.Equal(t, Angle{Degrees: 50}, result)
	})

	t.Run("Subtract Angle", func(t *testing.T) {
		angle1 := Angle{Degrees: 45}
		angle2 := Angle{Degrees: 15}

		result := angle1.Sub(angle2)

		assert.Equal(t, Angle{Degrees: 30}, result)
	})
	t.Run("Subtract int", func(t *testing.T) {
		angle := Angle{Degrees: 30}

		result := angle.Sub(20)

		assert.Equal(t, Angle{Degrees: 10}, result)
	})
	t.Run("Equal Angle", func(t *testing.T) {
		angle1 := Angle{Degrees: 45}
		angle2 := Angle{Degrees: 45}

		assert.True(t, angle1.Equal(angle2))
	})

	t.Run("Equal int", func(t *testing.T) {
		angle := Angle{Degrees: 45}

		assert.True(t, angle.Equal(45))
	})
}
