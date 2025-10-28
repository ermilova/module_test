package adapters

import (
	"good_abstracts/models"
	"good_abstracts/uobject"
	"math"
)

type MovingObjectAdapter struct {
	uobj uobject.UObject
}

func NewMovingObjectAdapter(uobj *uobject.UObject) *MovingObjectAdapter {
	return &MovingObjectAdapter{uobj: *uobj}
}

func (m *MovingObjectAdapter) GetLocation() models.Point {
	prop := m.uobj.GetProperty("location")
	if prop == nil {
		panic("Не указан location")
	}
	return prop.(models.Point)
}

func (m *MovingObjectAdapter) GetVelocity() models.Vector {
	angleProp := m.uobj.GetProperty("angle")
	velocityProp := m.uobj.GetProperty("velocity")

	if angleProp == nil || velocityProp == nil {
		panic("Угол или скорость не указаны")
	}

	angle := angleProp.(models.Angle)
	velocity := velocityProp.(float64)

	radian := angle.Radians()
	vx := int(velocity * math.Cos(radian))
	vy := int(velocity * math.Sin(radian))

	return models.Vector{X: vx, Y: vy}
}

func (m *MovingObjectAdapter) SetLocation(newPoint models.Point) {
	m.uobj.SetProperty("location", newPoint)
}

type RotatableObjectAdapter struct {
	uobj uobject.UObject
}

func NewRotatableObjectAdapter(uobj *uobject.UObject) *RotatableObjectAdapter {
	return &RotatableObjectAdapter{uobj: *uobj}
}

func (r *RotatableObjectAdapter) GetAngle() models.Angle {
	prop := r.uobj.GetProperty("angle")
	if prop == nil {
		panic("Уголь не найден")
	}
	return prop.(models.Angle)
}

func (r *RotatableObjectAdapter) SetAngle(newAngle models.Angle) {
	r.uobj.SetProperty("angle", newAngle)
}
