package main

import (
	"good_abstracts/actions"
	"good_abstracts/adapters"
	"good_abstracts/models"
	"good_abstracts/uobject"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntegration(t *testing.T) {
	t.Run("Тестирование сценария движения и поворота", func(t *testing.T) {
		// Создаем корабль как в main
		ship := uobject.NewUObject()
		ship.SetProperty("location", models.Point{X: 12, Y: 5})
		ship.SetProperty("angle", models.Angle{Degrees: 155})
		ship.SetProperty("velocity", 8.0)

		// Проверяем начальное состояние
		initialLocation := ship.GetProperty("location").(models.Point)
		initialAngle := ship.GetProperty("angle").(models.Angle)

		assert.Equal(t, models.Point{X: 12, Y: 5}, initialLocation)
		assert.Equal(t, models.Angle{Degrees: 155}, initialAngle)

		// Двигаем корабль
		moveAdapter := adapters.NewMovingObjectAdapter(ship)
		moveAction := actions.NewMove(moveAdapter)
		moveAction.Execute()

		// Проверяем новое положение
		locationAfterMove := ship.GetProperty("location").(models.Point)
		assert.Equal(t, models.Point{X: 5, Y: 8}, locationAfterMove)

		// Поворачиваем корабль
		rotateAdapter := adapters.NewRotatableObjectAdapter(ship)
		rotateAction := actions.NewRotate(rotateAdapter)
		rotateAction.Execute(models.Angle{Degrees: 10})

		// Проверяем новый угол
		angleAfterRotate := ship.GetProperty("angle").(models.Angle)
		assert.Equal(t, models.Angle{Degrees: 165}, angleAfterRotate)

		// Двигаем снова
		moveAction.Execute()

		// Проверяем финальное положение
		finalLocation := ship.GetProperty("location").(models.Point)
		assert.NotEqual(t, locationAfterMove, finalLocation)
	})
}
