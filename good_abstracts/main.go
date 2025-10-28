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
	ship := uobject.NewUObject()
	ship.SetProperty("location", models.Point{X: 12, Y: 5})
	ship.SetProperty("angle", models.Angle{Degrees: 155})
	ship.SetProperty("velocity", 8.0) // Модуль скорости

	fmt.Println("== Начало игры ==")
	printShipState(ship)

	// Движение
	moveAdapter := adapters.NewMovingObjectAdapter(ship)
	moveAction := actions.NewMove(moveAdapter)
	moveAction.Execute()
	fmt.Println("\n== После движения ==")
	printShipState(ship)

	// Поворот
	rotateAdapter := adapters.NewRotatableObjectAdapter(ship)
	rotateAction := actions.NewRotate(rotateAdapter)
	rotateAction.Execute(models.Angle{Degrees: 10})
	fmt.Println("\n== После поворота на 10° ==")
	printShipState(ship)

	// Движение после поворота
	moveAction.Execute()
	fmt.Println("\n== После движения ==")
	printShipState(ship)
}
