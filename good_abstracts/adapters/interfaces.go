package adapters

import "good_abstracts/models"

type MovingObject interface {
	GetLocation() models.Point
	GetVelocity() models.Vector
	SetLocation(newPoint models.Point)
}

type RotatableObject interface {
	GetAngle() models.Angle
	SetAngle(newAngle models.Angle)
}
