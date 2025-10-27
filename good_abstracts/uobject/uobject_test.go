package uobject

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUObject(t *testing.T) {
	t.Run("NewUObject creates empty object", func(t *testing.T) {
		obj := NewUObject()
		assert.NotNil(t, obj)
		assert.Nil(t, obj.GetProperty("nonexistent"))
	})

	t.Run("Set and Get property", func(t *testing.T) {
		obj := NewUObject()

		obj.SetProperty("test", "value")
		result := obj.GetProperty("test")

		assert.Equal(t, "value", result)
	})

	t.Run("Overwrite property", func(t *testing.T) {
		obj := NewUObject()

		obj.SetProperty("key", "old")
		obj.SetProperty("key", "new")

		assert.Equal(t, "new", obj.GetProperty("key"))
	})

	t.Run("Multiple properties", func(t *testing.T) {
		obj := NewUObject()

		obj.SetProperty("string", "text")
		obj.SetProperty("number", 42)
		obj.SetProperty("bool", true)

		assert.Equal(t, "text", obj.GetProperty("string"))
		assert.Equal(t, 42, obj.GetProperty("number"))
		assert.Equal(t, true, obj.GetProperty("bool"))
	})
}
