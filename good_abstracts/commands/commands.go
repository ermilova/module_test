package commands

import (
	"fmt"
	"math"

	"good_abstracts/actions"
	"good_abstracts/adapters"
	"good_abstracts/exceptions"
	"good_abstracts/models"
	"good_abstracts/uobject"
)

type Command struct {
	command adapters.CommandInterface
}

func NewCommand(command adapters.CommandInterface) *Command {
	return &Command{command: command}
}

func (c *Command) GetCommandType() adapters.CommandInterface {
	return c.command
}

func (c *Command) Execute() error {
	err := c.command.Execute()
	return err
}

// ----------------- //

type MacroCommand struct {
	commands []adapters.CommandInterface
}

func NewMacroCommand(commands []adapters.CommandInterface) *MacroCommand {
	return &MacroCommand{commands: commands}
}

func (m *MacroCommand) Execute() error {
	for _, cmd := range m.commands {
		if err := cmd.Execute(); err != nil {
			return exceptions.NewCommandException(
				fmt.Sprintf("MacroCommand failed on %T", cmd),
				err,
			)
		}
	}
	return nil
}

// ----------------- //

type CheckFuelCommand struct {
	uobj uobject.UObject
}

func NewCheckFuelCommand(uobj *uobject.UObject) *CheckFuelCommand {
	return &CheckFuelCommand{uobj: *uobj}
}

func (c *CheckFuelCommand) Execute() error {
	fuelProp := c.uobj.GetProperty("fuel")
	burnRateProp := c.uobj.GetProperty("fuel_burn_rate")

	if fuelProp == nil || burnRateProp == nil {
		return exceptions.NewCommandException("Не заданы параметры топлива: fuel или fuel_burn_rate")
	}

	fuel := fuelProp.(int)
	burnRate := burnRateProp.(int)

	if fuel < burnRate {
		return exceptions.NewCommandException("Недостаточно топлива для движения")
	}

	return nil
}

// ----------------- //

type BurnFuelCommand struct {
	uobj uobject.UObject
}

func NewBurnFuelCommand(uobj *uobject.UObject) *BurnFuelCommand {
	return &BurnFuelCommand{uobj: *uobj}
}

func (b *BurnFuelCommand) Execute() error {
	fuelProp := b.uobj.GetProperty("fuel")
	burnRateProp := b.uobj.GetProperty("fuel_burn_rate")

	if fuelProp == nil || burnRateProp == nil {
		return exceptions.NewCommandException("Не заданы параметры топлива: fuel или fuel_burn_rate")
	}

	fuel := fuelProp.(int)
	burnRate := burnRateProp.(int)
	newValue := fuel - burnRate
	if newValue < 0 {
		newValue = 0
	}

	b.uobj.SetProperty("fuel", newValue)
	return nil
}

// ----------------- //

type MoveCommand struct {
	action *actions.Move
}

func NewMoveCommand(moving *adapters.MovingObjectAdapter) *MoveCommand {
	return &MoveCommand{
		action: actions.NewMove(moving),
	}
}

func (m *MoveCommand) Execute() error {
	m.action.Execute()
	return nil
}

// ----------------- //

type RotateCommand struct {
	action     *actions.Rotate
	deltaAngle models.Angle
}

func NewRotateCommand(rotatable *adapters.RotatableObjectAdapter, deltaAngle models.Angle) *RotateCommand {
	return &RotateCommand{
		action:     actions.NewRotate(rotatable),
		deltaAngle: deltaAngle,
	}
}

func (r *RotateCommand) Execute() error {
	r.action.Execute(r.deltaAngle)
	return nil
}

// ----------------- //

type ModifyVelocityOnRotateCommand struct {
	uobj uobject.UObject
}

func NewModifyVelocityOnRotateCommand(uobj *uobject.UObject) *ModifyVelocityOnRotateCommand {
	return &ModifyVelocityOnRotateCommand{uobj: *uobj}
}

func (m *ModifyVelocityOnRotateCommand) Execute() error {
	speedProp := m.uobj.GetProperty("velocity")
	angleProp := m.uobj.GetProperty("angle")

	if speedProp == nil || angleProp == nil {
		return nil
	}

	speed := speedProp.(float64)
	angle := angleProp.(models.Angle)

	if speed == 0 {
		return nil
	}

	rad := angle.Radians()
	vx := int(speed * math.Cos(rad))
	vy := int(speed * math.Sin(rad))

	m.uobj.SetProperty("velocity_vector", models.Vector{X: vx, Y: vy})
	return nil
}

// ----------------- //

type MoveWithFuelMacroCommand struct {
	*MacroCommand
}

func NewMoveWithFuelMacroCommand(uobj *uobject.UObject, moving *adapters.MovingObjectAdapter) *MoveWithFuelMacroCommand {
	commands := []adapters.CommandInterface{
		NewCheckFuelCommand(uobj),
		NewMoveCommand(moving),
		NewBurnFuelCommand(uobj),
	}
	return &MoveWithFuelMacroCommand{
		MacroCommand: NewMacroCommand(commands),
	}
}

// ----------------- //

type RotateWithVelocityMacroCommand struct {
	*MacroCommand
}

func NewRotateWithVelocityMacroCommand(uobj *uobject.UObject, rotatable *adapters.RotatableObjectAdapter, deltaAngle models.Angle) *RotateWithVelocityMacroCommand {
	commands := []adapters.CommandInterface{
		NewRotateCommand(rotatable, deltaAngle),
		NewModifyVelocityOnRotateCommand(uobj),
	}
	return &RotateWithVelocityMacroCommand{
		MacroCommand: NewMacroCommand(commands),
	}
}
