package main

import (
	"github.com/cgiacomi/go-designpatterns/command"
)

// Invoker represents the type that is responsible of executing the concrete command
type Invoker struct {
	commands []command.Command
}

// RegisterCommand handles registration
func (i *Invoker) RegisterCommand(c command.Command) {
	i.commands = append(i.commands, c)
}

// ExecuteCommands is responsible for instantiation and execution
func (i *Invoker) ExecuteCommands() {
	for _, element := range i.commands {
		element.Execute()
	}
}

func main() {
	var c command.Command = command.ConcreteCommand{}

	i := Invoker{}
	i.RegisterCommand(c)
	i.ExecuteCommands()
}
