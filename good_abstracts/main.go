package main

import (
	"fmt"
	"good_abstracts/actions"
	"good_abstracts/adapters"
	"good_abstracts/models"
	"good_abstracts/uobject"
)

func printShipState(ship *uobject.UObject) {
	location := ship.GetProperty("location").(models.Point)
	angle := ship.GetProperty("angle").(models.Angle)
	velocity := ship.GetProperty("velocity").(float64)
	fmt.Printf("Location: %v, Angle: %v, Velocity: %.2f\n", location, angle, velocity)
}

func main() {
	// Создаем космический корабль
	ship := uobject.NewUObject()
	ship.SetProperty("location", models.Point{X: 12, Y: 5})
	ship.SetProperty("angle", models.Angle{Degrees: 45})
	ship.SetProperty("velocity", 10.0) // Модуль скорости

	fmt.Println("== INITIAL STATE ==")
	printShipState(ship)

	// Двигаем корабль
	moveAdapter := adapters.NewMovingObjectAdapter(ship)
	moveAction := actions.NewMove(moveAdapter)
	moveAction.Execute()
	fmt.Println("\n== AFTER MOVE ==")
	printShipState(ship)

	// Поворачиваем корабль
	rotateAdapter := adapters.NewRotatableObjectAdapter(ship)
	rotateAction := actions.NewRotate(rotateAdapter)
	rotateAction.Execute(models.Angle{Degrees: 15})
	fmt.Println("\n== AFTER ROTATE +15° ==")
	printShipState(ship)

	// Двигаем снова (с новым углом)
	moveAction.Execute()
	fmt.Println("\n== AFTER SECOND MOVE ==")
	printShipState(ship)
}
