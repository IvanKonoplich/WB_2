package tv

import "WB2/pattern/04_command/interfaces"

type offCommand struct {
	device interfaces.Device
}

func NewOffCommand(device interfaces.Device) *offCommand {
	return &offCommand{
		device: device,
	}
}

func (c *offCommand) Execute() {
	c.device.Off()
}
