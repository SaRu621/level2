package main

import (
	"command/pkg"
	"fmt"
)

type Tv struct {
	isRunning bool
}

func (t *Tv) On() {
	t.isRunning = true
	fmt.Println("TV switch on")
}

func (t *Tv) Off() {
	t.isRunning = false
	fmt.Println("TV switch off")
}

func main() {

	tv := &Tv{}

	onCommand := &pkg.OnCommand{Device: tv}

	offCommand := &pkg.OffCommand{Device: tv}

	onButton := &pkg.Button{Command: onCommand}

	onButton.Press()

	fmt.Println(tv)

	offButton := &pkg.Button{Command: offCommand}

	offButton.Press()

	fmt.Println(tv)
}
