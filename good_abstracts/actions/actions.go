package actions

import (
	"good_abstracts/adapters"
	"good_abstracts/models"
)

type Move struct {
	movingObject adapters.MovingObject
}

func NewMove(movingObject adapters.MovingObject) *Move {
	return &Move{movingObject: movingObject}
}

func (m *Move) Execute() {
	location := m.movingObject.GetLocation()
	velocity := m.movingObject.GetVelocity()

	newLoc := location.Add(velocity)
	m.movingObject.SetLocation(newLoc)
}

type Rotate struct {
	rotatableObject adapters.RotatableObject
}

func NewRotate(rotatableObject adapters.RotatableObject) *Rotate {
	return &Rotate{rotatableObject: rotatableObject}
}

func (r *Rotate) Execute(deltaAngle models.Angle) {
	angle := r.rotatableObject.GetAngle()
	newAngle := angle.Add(deltaAngle)
	r.rotatableObject.SetAngle(newAngle)
}
