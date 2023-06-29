package main

import (
	"context"
	"machine"
	"time"
)

var (
	btnA    machine.Pin
	btnB    machine.Pin
	btnUp   machine.Pin
	btnDown machine.Pin
)

func initButton(pin machine.Pin, mode string) machine.Pin {
	switch mode {
	case "input":
		pin.Configure(machine.PinConfig{Mode: machine.PinInput})
	case "output":
		pin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	}

	return pin
}

func setupButtons() {
	btnA = initButton(machine.BUTTON_A, "input")
	btnB = initButton(machine.BUTTON_B, "input")
	btnUp = initButton(machine.BUTTON_UP, "input")
	btnDown = initButton(machine.BUTTON_DOWN, "input")
}

func listenToBtnPress(ctx context.Context, showTwitter chan bool, enableLEDs chan bool) {
	for {
		select {
		case <-ctx.Done():
			break
		default:
			switch {
			case !btnA.Get():
				showTwitter <- true
			case !btnB.Get():
				showTwitter <- false
			case !btnUp.Get():
				enableLEDs <- false
			case !btnDown.Get():
				enableLEDs <- true
			}

			time.Sleep(100 * time.Millisecond)
		}
	}
}
