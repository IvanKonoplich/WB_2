package tv

import "WB2/patterns/command/interfaces"

type onCommand struct {
	device interfaces.Device
}

func NewOnCommand(device interfaces.Device) *onCommand {
	return &onCommand{
		device: device,
	}
}

func (c *onCommand) Execute() {
	c.device.On()
}
