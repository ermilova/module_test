package commands

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"good_abstracts/adapters"
	"good_abstracts/exceptions"
	"good_abstracts/models"
	"good_abstracts/uobject"
)

type MockCommand struct {
	mock.Mock
}

func (m *MockCommand) Execute() error {
	args := m.Called()
	return args.Error(0)
}

type MockMovingObjectAdapter struct {
	mock.Mock
}

func (m *MockMovingObjectAdapter) GetLocation() models.Point {
	args := m.Called()
	return args.Get(0).(models.Point)
}

func (m *MockMovingObjectAdapter) GetVelocity() models.Vector {
	args := m.Called()
	return args.Get(0).(models.Vector)
}

func (m *MockMovingObjectAdapter) SetLocation(newPoint models.Point) {
	m.Called(newPoint)
}

type MockRotatableObjectAdapter struct {
	mock.Mock
}

func (m *MockRotatableObjectAdapter) GetAngle() models.Angle {
	args := m.Called()
	return args.Get(0).(models.Angle)
}

func (m *MockRotatableObjectAdapter) SetAngle(newAngle models.Angle) {
	m.Called(newAngle)
}

// MoveCommand (actions.Move)
type TestMoveAction struct {
	mock.Mock
	moving adapters.MovingObject
}

func NewTestMoveAction(moving adapters.MovingObject) *TestMoveAction {
	return &TestMoveAction{moving: moving}
}

func (t *TestMoveAction) Execute() {
	t.Called()
	location := t.moving.GetLocation()
	velocity := t.moving.GetVelocity()
	newLocation := location.Add(velocity)
	t.moving.SetLocation(newLocation)
}

// RotateCommand (actions.Rotate)
type TestRotateAction struct {
	mock.Mock
	rotatable adapters.RotatableObject
}

func NewTestRotateAction(rotatable adapters.RotatableObject) *TestRotateAction {
	return &TestRotateAction{rotatable: rotatable}
}

func (t *TestRotateAction) Execute(deltaAngle models.Angle) {
	t.Called(deltaAngle)
	currentAngle := t.rotatable.GetAngle()
	newAngle := currentAngle.Add(deltaAngle)
	t.rotatable.SetAngle(newAngle)
}

func TestCommand(t *testing.T) {
	t.Run("Execute исполняется корректно", func(t *testing.T) {
		mockCmd := new(MockCommand)
		mockCmd.On("Execute").Return(nil)

		cmd := NewCommand(mockCmd)
		err := cmd.Execute()

		assert.NoError(t, err)
		mockCmd.AssertExpectations(t)
	})

	t.Run("Execute вернул ошибку от команды", func(t *testing.T) {
		expectedErr := errors.New("test error")
		mockCmd := new(MockCommand)
		mockCmd.On("Execute").Return(expectedErr)

		cmd := NewCommand(mockCmd)
		err := cmd.Execute()

		assert.Equal(t, expectedErr, err)
		mockCmd.AssertExpectations(t)
	})

	t.Run("GetCommandType возвращает тип команды", func(t *testing.T) {
		mockCmd := new(MockCommand)
		cmd := NewCommand(mockCmd)

		result := cmd.GetCommandType()

		assert.Equal(t, mockCmd, result)
	})
}

func TestMacroCommand(t *testing.T) {
	t.Run("Execute все команды, успех", func(t *testing.T) {
		cmd1 := new(MockCommand)
		cmd2 := new(MockCommand)
		cmd3 := new(MockCommand)

		cmd1.On("Execute").Return(nil)
		cmd2.On("Execute").Return(nil)
		cmd3.On("Execute").Return(nil)

		macro := NewMacroCommand([]adapters.CommandInterface{cmd1, cmd2, cmd3})
		err := macro.Execute()

		assert.NoError(t, err)
		cmd1.AssertExpectations(t)
		cmd2.AssertExpectations(t)
		cmd3.AssertExpectations(t)
	})

	t.Run("Останавливаем выполнение на первой ошибке", func(t *testing.T) {
		cmd1 := new(MockCommand)
		cmd2 := new(MockCommand)
		cmd3 := new(MockCommand)

		expectedErr := errors.New("command 2 failed")
		cmd1.On("Execute").Return(nil)
		cmd2.On("Execute").Return(expectedErr)

		macro := NewMacroCommand([]adapters.CommandInterface{cmd1, cmd2, cmd3})
		err := macro.Execute()

		assert.Error(t, err)
		assert.IsType(t, &exceptions.CommandException{}, err)
		assert.Contains(t, err.Error(), "MacroCommand failed")
		cmd1.AssertExpectations(t)
		cmd2.AssertExpectations(t)
		cmd3.AssertNotCalled(t, "Execute")
	})

	t.Run("Пустой список команд в макрокоманде", func(t *testing.T) {
		macro := NewMacroCommand([]adapters.CommandInterface{})
		err := macro.Execute()

		assert.NoError(t, err)
	})
}

func TestCheckFuelCommand(t *testing.T) {
	t.Run("Разрешать движение при избыточном кличестве топлива", func(t *testing.T) {
		obj := uobject.NewUObject()
		obj.SetProperty("fuel", 10)
		obj.SetProperty("fuel_burn_rate", 5)

		cmd := NewCheckFuelCommand(obj)
		err := cmd.Execute()

		assert.NoError(t, err)
	})

	t.Run("Запрещать движение при недостаточном количестве топлива", func(t *testing.T) {
		obj := uobject.NewUObject()
		obj.SetProperty("fuel", 3)
		obj.SetProperty("fuel_burn_rate", 5)

		cmd := NewCheckFuelCommand(obj)
		err := cmd.Execute()

		assert.Error(t, err)
		assert.IsType(t, &exceptions.CommandException{}, err)
		assert.Contains(t, err.Error(), "Недостаточно топлива для движения")
	})

	t.Run("Выбрасывать ошибку при отсутствии свойства fuel", func(t *testing.T) {
		obj := uobject.NewUObject()
		obj.SetProperty("fuel_burn_rate", 5)

		cmd := NewCheckFuelCommand(obj)
		err := cmd.Execute()

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "Не заданы параметры топлива")
	})

	t.Run("Выбрасывать ошибку при отсутствии свойства fuel_burn_rate", func(t *testing.T) {
		obj := uobject.NewUObject()
		obj.SetProperty("fuel", 10)

		cmd := NewCheckFuelCommand(obj)
		err := cmd.Execute()

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "Не заданы параметры топлива")
	})

	t.Run("Разрешать движение, если топлива хатает тютелька в тютельку", func(t *testing.T) {
		obj := uobject.NewUObject()
		obj.SetProperty("fuel", 5)
		obj.SetProperty("fuel_burn_rate", 5)

		cmd := NewCheckFuelCommand(obj)
		err := cmd.Execute()

		assert.NoError(t, err)
	})
}

func TestBurnFuelCommand(t *testing.T) {
	t.Run("Успешно сожги топливо", func(t *testing.T) {
		obj := uobject.NewUObject()
		obj.SetProperty("fuel", 10)
		obj.SetProperty("fuel_burn_rate", 3)

		cmd := NewBurnFuelCommand(obj)
		err := cmd.Execute()

		assert.NoError(t, err)
		assert.Equal(t, 7, obj.GetProperty("fuel").(int))
	})

	t.Run("Сожгли все топливо", func(t *testing.T) {
		obj := uobject.NewUObject()
		obj.SetProperty("fuel", 2)
		obj.SetProperty("fuel_burn_rate", 5)

		cmd := NewBurnFuelCommand(obj)
		err := cmd.Execute()

		assert.NoError(t, err)
		assert.Equal(t, 0, obj.GetProperty("fuel").(int))
	})

	t.Run("Куда-то пропало свойство fuel", func(t *testing.T) {
		obj := uobject.NewUObject()
		obj.SetProperty("fuel_burn_rate", 5)

		cmd := NewBurnFuelCommand(obj)
		err := cmd.Execute()

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "Не заданы параметры топлива")
	})

	t.Run("Куда-то пропало свойство fuel_burn_rate", func(t *testing.T) {
		obj := uobject.NewUObject()
		obj.SetProperty("fuel", 10)

		cmd := NewBurnFuelCommand(obj)
		err := cmd.Execute()

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "Не заданы параметры топлива")
	})

	t.Run("Сожгли топливо, которого едва хватало", func(t *testing.T) {
		obj := uobject.NewUObject()
		obj.SetProperty("fuel", 8)
		obj.SetProperty("fuel_burn_rate", 8)

		cmd := NewBurnFuelCommand(obj)
		err := cmd.Execute()

		assert.NoError(t, err)
		assert.Equal(t, 0, obj.GetProperty("fuel").(int))
	})
}

func TestMoveCommand(t *testing.T) {
	t.Run("Выполняем движение с адаптером", func(t *testing.T) {
		obj := uobject.NewUObject()
		obj.SetProperty("location", models.Point{X: 0, Y: 0})
		obj.SetProperty("angle", models.Angle{Degrees: 0})
		obj.SetProperty("velocity", 5.0)

		movingAdapter := adapters.NewMovingObjectAdapter(obj)

		cmd := NewMoveCommand(movingAdapter)
		err := cmd.Execute()

		assert.NoError(t, err)

		finalLocation := obj.GetProperty("location").(models.Point)
		assert.NotEqual(t, models.Point{X: 0, Y: 0}, finalLocation)
	})

	t.Run("Создали команду с адаптером", func(t *testing.T) {
		obj := uobject.NewUObject()
		movingAdapter := adapters.NewMovingObjectAdapter(obj)

		cmd := NewMoveCommand(movingAdapter)

		assert.NotNil(t, cmd)
		assert.NotNil(t, cmd.action)
	})
}

func TestRotateCommand(t *testing.T) {
	t.Run("Вращаемся с адаптером", func(t *testing.T) {
		obj := uobject.NewUObject()
		obj.SetProperty("angle", models.Angle{Degrees: 0})

		rotatableAdapter := adapters.NewRotatableObjectAdapter(obj)
		deltaAngle := models.Angle{Degrees: 45}

		cmd := NewRotateCommand(rotatableAdapter, deltaAngle)
		err := cmd.Execute()

		assert.NoError(t, err)

		finalAngle := obj.GetProperty("angle").(models.Angle)
		assert.Equal(t, models.Angle{Degrees: 45}, finalAngle)
	})

	t.Run("Создали команду с адаптером и еще углом", func(t *testing.T) {
		obj := uobject.NewUObject()
		rotatableAdapter := adapters.NewRotatableObjectAdapter(obj)
		deltaAngle := models.Angle{Degrees: 30}

		cmd := NewRotateCommand(rotatableAdapter, deltaAngle)

		assert.NotNil(t, cmd)
		assert.NotNil(t, cmd.action)
		assert.Equal(t, deltaAngle, cmd.deltaAngle)
	})
}

func TestModifyVelocityOnRotateCommand(t *testing.T) {
	t.Run("Изменили скорость", func(t *testing.T) {
		obj := uobject.NewUObject()
		obj.SetProperty("velocity", 10.0)
		obj.SetProperty("angle", models.Angle{Degrees: 0}) // 0 degrees = right direction

		cmd := NewModifyVelocityOnRotateCommand(obj)
		err := cmd.Execute()

		assert.NoError(t, err)

		velocityVector := obj.GetProperty("velocity_vector").(models.Vector)
		assert.Equal(t, models.Vector{X: 10, Y: 0}, velocityVector)
	})

	t.Run("Изменили угол на 90", func(t *testing.T) {
		obj := uobject.NewUObject()
		obj.SetProperty("velocity", 5.0)
		obj.SetProperty("angle", models.Angle{Degrees: 90}) // 90 degrees = up direction

		cmd := NewModifyVelocityOnRotateCommand(obj)
		err := cmd.Execute()

		assert.NoError(t, err)

		velocityVector := obj.GetProperty("velocity_vector").(models.Vector)
		assert.Equal(t, 0, velocityVector.X)
		assert.Equal(t, 5, velocityVector.Y)
	})

	t.Run("Изменили на ноль", func(t *testing.T) {
		obj := uobject.NewUObject()
		obj.SetProperty("velocity", 0.0)
		obj.SetProperty("angle", models.Angle{Degrees: 45})

		cmd := NewModifyVelocityOnRotateCommand(obj)
		err := cmd.Execute()

		assert.NoError(t, err)
		assert.Nil(t, obj.GetProperty("velocity_vector"))
	})

	t.Run("пропало свойство velocity", func(t *testing.T) {
		obj := uobject.NewUObject()
		obj.SetProperty("angle", models.Angle{Degrees: 45})

		cmd := NewModifyVelocityOnRotateCommand(obj)
		err := cmd.Execute()

		assert.NoError(t, err)
		assert.Nil(t, obj.GetProperty("velocity_vector"))
	})

	t.Run("пропало свойство вектор angle", func(t *testing.T) {
		obj := uobject.NewUObject()
		obj.SetProperty("velocity", 10.0)

		cmd := NewModifyVelocityOnRotateCommand(obj)
		err := cmd.Execute()

		assert.NoError(t, err)
		assert.Nil(t, obj.GetProperty("velocity_vector"))
	})
}

func TestMoveWithFuelMacroCommand(t *testing.T) {
	t.Run("Create macro command with correct sequence", func(t *testing.T) {
		obj := uobject.NewUObject()
		movingAdapter := adapters.NewMovingObjectAdapter(obj)

		cmd := NewMoveWithFuelMacroCommand(obj, movingAdapter)

		assert.NotNil(t, cmd)
		assert.NotNil(t, cmd.MacroCommand)
	})

	t.Run("Макрокоманда успешно выполнилась", func(t *testing.T) {
		obj := uobject.NewUObject()
		obj.SetProperty("fuel", 10)
		obj.SetProperty("fuel_burn_rate", 3)
		obj.SetProperty("location", models.Point{X: 0, Y: 0})
		obj.SetProperty("angle", models.Angle{Degrees: 0})
		obj.SetProperty("velocity", 5.0)

		movingAdapter := adapters.NewMovingObjectAdapter(obj)

		cmd := NewMoveWithFuelMacroCommand(obj, movingAdapter)
		err := cmd.Execute()

		assert.NoError(t, err)
		assert.Equal(t, 7, obj.GetProperty("fuel").(int)) // 10 - 3 = 7

		location := obj.GetProperty("location").(models.Point)
		assert.NotEqual(t, models.Point{X: 0, Y: 0}, location)
	})

	t.Run("Топлива не хватило - макрокоманда", func(t *testing.T) {
		obj := uobject.NewUObject()
		obj.SetProperty("fuel", 2)
		obj.SetProperty("fuel_burn_rate", 5) // Not enough fuel
		obj.SetProperty("location", models.Point{X: 0, Y: 0})
		obj.SetProperty("angle", models.Angle{Degrees: 0})
		obj.SetProperty("velocity", 5.0)

		movingAdapter := adapters.NewMovingObjectAdapter(obj)

		cmd := NewMoveWithFuelMacroCommand(obj, movingAdapter)
		err := cmd.Execute()

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "Недостаточно топлива")
		assert.Equal(t, 2, obj.GetProperty("fuel").(int))
		location := obj.GetProperty("location").(models.Point)
		assert.Equal(t, models.Point{X: 0, Y: 0}, location)
	})
}

func TestRotateWithVelocityMacroCommand(t *testing.T) {
	t.Run("макрокоманда", func(t *testing.T) {
		obj := uobject.NewUObject()
		rotatableAdapter := adapters.NewRotatableObjectAdapter(obj)
		deltaAngle := models.Angle{Degrees: 45}

		cmd := NewRotateWithVelocityMacroCommand(obj, rotatableAdapter, deltaAngle)

		assert.NotNil(t, cmd)
		assert.NotNil(t, cmd.MacroCommand)
	})

	t.Run("Успешное выполнение", func(t *testing.T) {
		obj := uobject.NewUObject()
		obj.SetProperty("angle", models.Angle{Degrees: 0})
		obj.SetProperty("velocity", 10.0)

		rotatableAdapter := adapters.NewRotatableObjectAdapter(obj)
		deltaAngle := models.Angle{Degrees: 45}

		cmd := NewRotateWithVelocityMacroCommand(obj, rotatableAdapter, deltaAngle)
		err := cmd.Execute()

		assert.NoError(t, err)

		finalAngle := obj.GetProperty("angle").(models.Angle)
		assert.Equal(t, models.Angle{Degrees: 45}, finalAngle)

		velocityVector := obj.GetProperty("velocity_vector").(models.Vector)
		assert.NotEqual(t, models.Vector{X: 0, Y: 0}, velocityVector)
	})

	t.Run("Считаем скорочть после поворота", func(t *testing.T) {
		obj := uobject.NewUObject()
		obj.SetProperty("angle", models.Angle{Degrees: 0})
		obj.SetProperty("velocity", 10.0)

		rotatableAdapter := adapters.NewRotatableObjectAdapter(obj)
		deltaAngle := models.Angle{Degrees: 90}

		cmd := NewRotateWithVelocityMacroCommand(obj, rotatableAdapter, deltaAngle)
		err := cmd.Execute()

		assert.NoError(t, err)

		velocityVector := obj.GetProperty("velocity_vector").(models.Vector)
		assert.InDelta(t, 0, velocityVector.X, 0.1)
		assert.InDelta(t, 10, velocityVector.Y, 0.1)
	})
}

func TestIntegration(t *testing.T) {
	t.Run("Интеграционный тест для проверки совместной работы команд", func(t *testing.T) {
		obj := uobject.NewUObject()
		obj.SetProperty("fuel", 20)
		obj.SetProperty("fuel_burn_rate", 2)
		obj.SetProperty("location", models.Point{X: 0, Y: 0})
		obj.SetProperty("angle", models.Angle{Degrees: 0})
		obj.SetProperty("velocity", 5.0)

		movingAdapter := adapters.NewMovingObjectAdapter(obj)
		rotatableAdapter := adapters.NewRotatableObjectAdapter(obj)

		// Выполняем несколько движений и поворотов
		moveCmd := NewMoveWithFuelMacroCommand(obj, movingAdapter)
		rotateCmd := NewRotateWithVelocityMacroCommand(obj, rotatableAdapter, models.Angle{Degrees: 90})

		// Двигаемся
		err := moveCmd.Execute()
		assert.NoError(t, err)

		// Поворачиваем
		err = rotateCmd.Execute()
		assert.NoError(t, err)

		// Двигаемся снова (уже с новым направлением)
		err = moveCmd.Execute()
		assert.NoError(t, err)

		// Проверяем итоговое состояние
		assert.Equal(t, 16, obj.GetProperty("fuel").(int)) // 20 - 2*2 = 16
		finalLocation := obj.GetProperty("location").(models.Point)
		finalAngle := obj.GetProperty("angle").(models.Angle)

		assert.NotEqual(t, models.Point{X: 0, Y: 0}, finalLocation)
		assert.Equal(t, models.Angle{Degrees: 90}, finalAngle)
	})
}
