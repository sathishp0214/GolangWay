package main

import "fmt"

type Command interface {
	Execute()
}

// LightReceiver - an example receiver
type LightReceiver struct{}

func (l *LightReceiver) TurnOn() {
	fmt.Println("Light is ON")
}
func (l *LightReceiver) TurnOff() {
	fmt.Println("Light is OFF")
}

// Concrete Command
type LightOnCommand struct {
	light *LightReceiver
}

func (c *LightOnCommand) Execute() {
	c.light.TurnOn()
}

type LightOffCommand struct {
	light *LightReceiver
}

func (c *LightOffCommand) Execute() {
	c.light.TurnOff()
}

// Invoker
type RemoteControl struct {
	command Command
}

func (r *RemoteControl) SetCommand(command Command) {
	r.command = command
}
func (r *RemoteControl) PressButton() {
	r.command.Execute()
}

// Client code
func main() {
	light := &LightReceiver{}
	onCommand := &LightOnCommand{light}
	offCommand := &LightOffCommand{light}
	remote := &RemoteControl{}
	remote.SetCommand(onCommand)
	remote.PressButton()
	remote.SetCommand(offCommand)
	remote.PressButton()
}
