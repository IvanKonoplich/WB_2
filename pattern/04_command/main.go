package main

import "WB2/pattern/command/tv"

func main() {
	tv1 := tv.NewTv()
	onCommand := tv.NewOnCommand(tv1)
	offCommand := tv.NewOffCommand(tv1)
	onButton := tv.NewButton(onCommand)
	onButton.Press()
	offButton := tv.NewButton(offCommand)
	offButton.Press()
}
