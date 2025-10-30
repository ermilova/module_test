package exceptions

import "fmt"

type CommandException struct {
	Message string
	Cause   error
}

func (c *CommandException) Error() string {
	if c.Cause != nil {
		return fmt.Sprintf("%s: %v", c.Message, c.Cause)
	}
	return c.Message
}

func NewCommandException(message string, cause ...error) *CommandException {
	if len(cause) > 0 {
		return &CommandException{Message: message, Cause: cause[0]}
	}
	return &CommandException{Message: message}
}
