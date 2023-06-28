package tv

import "WB2/pattern/command/interfaces"

type button struct {
	command interfaces.Command
}

func NewButton(command interfaces.Command) *button {
	return &button{
		command: command,
	}
}

func (b *button) Press() {
	b.command.Execute()
}
