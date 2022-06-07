package command_pattern

// RemoteControlWithUndo 控制器
type RemoteControlWithUndo struct {
	onCommands  []Command
	offCommands []Command
	undoCommand Command
}

func (c *RemoteControlWithUndo) setCommand(slot int, onCommand, offCommand Command) {
	c.onCommands[slot] = onCommand
	c.offCommands[slot] = offCommand
}

func (c *RemoteControlWithUndo) onButtonWasPushed(slot int) {
	c.onCommands[slot].execute()
	c.undoCommand = c.onCommands[slot]
}
func (c *RemoteControlWithUndo) offButtonWasPushed(slot int) {
	c.offCommands[slot].execute()
	c.undoCommand = c.offCommands[slot]
}

func (c *RemoteControlWithUndo) undoButtonWasPushed() {
	c.undoCommand.execute()
}

func NewRemoteControlWithUndo() *RemoteControlWithUndo {
	control := &RemoteControlWithUndo{
		onCommands:  make([]Command, 7),
		offCommands: make([]Command, 7),
	}
	for i := 0; i < 7; i++ {
		control.onCommands[i] = &NoCommand{}
		control.offCommands[i] = &NoCommand{}
	}
	control.undoCommand = &NoCommand{}
	return control
}
