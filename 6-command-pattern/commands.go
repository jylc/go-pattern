package command_pattern

import "fmt"

// Command 命令分装成类
type Command interface {
	execute()
	undo()
}
type operate interface {
	on()
	off()
}

type Light struct {
}

func (l Light) on() {
	fmt.Println("Light on")
}
func (l Light) off() {
	fmt.Println("Light off")
}

type NoCommand struct {
}

func (c *NoCommand) execute() {

}

func (c *NoCommand) undo() {

}

type LightOffCommand struct {
	light Light
}

func (c *LightOffCommand) execute() {
	c.light.off()
}

func (c *LightOffCommand) undo() {
	c.light.on()
}

func NewLightOffCommand(light Light) *LightOffCommand {
	return &LightOffCommand{light: light}
}

type LightOnCommand struct {
	light Light
}

func (c *LightOnCommand) execute() {
	c.light.on()
}

func (c *LightOnCommand) undo() {
	c.light.off()
}
func NewLightOnCommand(light Light) *LightOnCommand {
	return &LightOnCommand{light: light}
}
