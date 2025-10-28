package adapters

import (
	"good_abstracts/models"
	"good_abstracts/uobject"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMovingObjectAdapter(t *testing.T) {
	t.Run("GetLocation возвращает корректное значение", func(t *testing.T) {
		obj := uobject.NewUObject()
		expected := models.Point{X: 10, Y: 20}
		obj.SetProperty("location", expected)

		adapter := NewMovingObjectAdapter(obj)
		result := adapter.GetLocation()

		assert.Equal(t, expected, result)
	})

	t.Run("GetLocation panics", func(t *testing.T) {
		obj := uobject.NewUObject()
		adapter := NewMovingObjectAdapter(obj)

		assert.Panics(t, func() {
			adapter.GetLocation()
		})
	})

	t.Run("GetVelocity вычисляет корректный вектор", func(t *testing.T) {
		obj := uobject.NewUObject()
		obj.SetProperty("angle", models.Angle{Degrees: 0}) // 0 degrees = right direction
		obj.SetProperty("velocity", 10.0)

		adapter := NewMovingObjectAdapter(obj)
		result := adapter.GetVelocity()

		// At 0 degrees, cos=1, sin=0, so vector should be (10, 0)
		assert.Equal(t, models.Vector{X: 10, Y: 0}, result)
	})

	t.Run("GetVelocity и 90 градусов", func(t *testing.T) {
		obj := uobject.NewUObject()
		obj.SetProperty("angle", models.Angle{Degrees: 90}) // 90 degrees = up direction
		obj.SetProperty("velocity", 5.0)

		adapter := NewMovingObjectAdapter(obj)
		result := adapter.GetVelocity()

		// At 90 degrees, cos=0, sin=1, so vector should be (0, 5)
		assert.Equal(t, models.Vector{X: 0, Y: 5}, result)
	})

	t.Run("SetLocation обновляет значение", func(t *testing.T) {
		obj := uobject.NewUObject()
		adapter := NewMovingObjectAdapter(obj)
		newPoint := models.Point{X: 100, Y: 200}

		adapter.SetLocation(newPoint)

		result := obj.GetProperty("location").(models.Point)
		assert.Equal(t, newPoint, result)
	})
}

func TestRotatableObjectAdapter(t *testing.T) {
	t.Run("GetAngle возвраает корректный угол", func(t *testing.T) {
		obj := uobject.NewUObject()
		expected := models.Angle{Degrees: 45}
		obj.SetProperty("angle", expected)

		adapter := NewRotatableObjectAdapter(obj)
		result := adapter.GetAngle()

		assert.Equal(t, expected, result)
	})

	t.Run("GetAngle panics", func(t *testing.T) {
		obj := uobject.NewUObject()
		adapter := NewRotatableObjectAdapter(obj)

		assert.Panics(t, func() {
			adapter.GetAngle()
		})
	})

	t.Run("SetAngle обновляет значение", func(t *testing.T) {
		obj := uobject.NewUObject()
		adapter := NewRotatableObjectAdapter(obj)
		newAngle := models.Angle{Degrees: 90}

		adapter.SetAngle(newAngle)

		result := obj.GetProperty("angle").(models.Angle)
		assert.Equal(t, newAngle, result)
	})
}
