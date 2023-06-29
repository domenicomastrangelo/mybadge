package main

import (
	"context"
	"image/color"
	"time"

	"machine"

	"tinygo.org/x/drivers/ws2812"
)

var (
	RGBA_BLACK     = color.RGBA{0, 0, 0, 0}
	RGBA_WHITE     = color.RGBA{255, 255, 255, 1}
	RGBA_TURQUOISE = color.RGBA{0, 255, 255, 1}
	RGBA_RED       = color.RGBA{255, 25, 71, 1}
	RGBA_TWITTER   = color.RGBA{29, 161, 242, 1}
	RGBA_LINKEDIN  = color.RGBA{0, 119, 181, 1}

	showTwitter = make(chan bool, 1)
	enableLEDs  = make(chan bool, 1)

	ctx = context.Background()

	virtualDisplay = setupDisplay()
)

func main() {
	setupButtons()
	writeToScreen(virtualDisplay, false)

	go listenToBtnPress(ctx, showTwitter, enableLEDs)
	go setupEyeColors(ctx)

	run()
}

func run() {
	for {
		select {
		case <-ctx.Done():
			break
		default:
			select {
			case show := <-showTwitter:
				writeToScreen(virtualDisplay, show)
			case enable := <-enableLEDs:
				if enable {
					machine.NEOPIXELS.Configure(machine.PinConfig{Mode: machine.PinOutput})
				} else {
					neopixels := machine.NEOPIXELS
					neopixels.Configure(machine.PinConfig{Mode: machine.PinInput})

					leds := ws2812.New(neopixels)
					leds.WriteColors([]color.RGBA{
						RGBA_BLACK,
						RGBA_BLACK,
					})
				}
			}

			time.Sleep(100 * time.Millisecond)
		}
	}
}
