package command

import "fmt"

// Command is the interface that defines the behavior of the Command pattern
type Command interface {
	Execute()
}

// ConcreteCommand is a type representing and actual Command
type ConcreteCommand struct{}

// Execute is responsible for carrying out the work for this command
func (c ConcreteCommand) Execute() {
	fmt.Println("test command")
}
