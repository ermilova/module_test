package actions

import (
	"good_abstracts/models"
	"testing"

	"github.com/stretchr/testify/mock"
)

// MockMovingObject для тестирования Move
type MockMovingObject struct {
	mock.Mock
}

func (m *MockMovingObject) GetLocation() models.Point {
	args := m.Called()
	return args.Get(0).(models.Point)
}

func (m *MockMovingObject) GetVelocity() models.Vector {
	args := m.Called()
	return args.Get(0).(models.Vector)
}

func (m *MockMovingObject) SetLocation(newPoint models.Point) {
	m.Called(newPoint)
}

// MockRotatableObject для тестирования Rotate
type MockRotatableObject struct {
	mock.Mock
}

func (m *MockRotatableObject) GetAngle() models.Angle {
	args := m.Called()
	return args.Get(0).(models.Angle)
}

func (m *MockRotatableObject) SetAngle(newAngle models.Angle) {
	m.Called(newAngle)
}

func TestMove(t *testing.T) {
	t.Run("Execute moves object correctly", func(t *testing.T) {
		mockObj := new(MockMovingObject)

		// Настройка моков
		initialLocation := models.Point{X: 10, Y: 20}
		velocity := models.Vector{X: 5, Y: -3}
		expectedNewLocation := models.Point{X: 15, Y: 17}

		mockObj.On("GetLocation").Return(initialLocation)
		mockObj.On("GetVelocity").Return(velocity)
		mockObj.On("SetLocation", expectedNewLocation).Once()

		// Выполнение
		move := NewMove(mockObj)
		move.Execute()

		// Проверка
		mockObj.AssertExpectations(t)
	})
}

func TestRotate(t *testing.T) {
	t.Run("Execute rotates object correctly", func(t *testing.T) {
		mockObj := new(MockRotatableObject)

		// Настройка моков
		initialAngle := models.Angle{Degrees: 30}
		deltaAngle := models.Angle{Degrees: 15}
		expectedNewAngle := models.Angle{Degrees: 45}

		mockObj.On("GetAngle").Return(initialAngle)
		mockObj.On("SetAngle", expectedNewAngle).Once()

		// Выполнение
		rotate := NewRotate(mockObj)
		rotate.Execute(deltaAngle)

		// Проверка
		mockObj.AssertExpectations(t)
	})

	t.Run("Execute with negative rotation", func(t *testing.T) {
		mockObj := new(MockRotatableObject)

		initialAngle := models.Angle{Degrees: 45}
		deltaAngle := models.Angle{Degrees: -20}
		expectedNewAngle := models.Angle{Degrees: 25}

		mockObj.On("GetAngle").Return(initialAngle)
		mockObj.On("SetAngle", expectedNewAngle).Once()

		rotate := NewRotate(mockObj)
		rotate.Execute(deltaAngle)

		mockObj.AssertExpectations(t)
	})
}
