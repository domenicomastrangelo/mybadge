package main

import (
	"context"
	"image/color"
	"machine"
	"time"

	"tinygo.org/x/drivers/ws2812"
)

func setupEyeColors(ctx context.Context) {
	neopixels := machine.NEOPIXELS
	neopixels.Configure(machine.PinConfig{Mode: machine.PinInput})

	leds := ws2812.New(neopixels)

	for {
		select {
		case <-ctx.Done():
			break
		default:
			for i := 0; i < 10; i++ {
				setNormalEyeColor(leds)
			}

			for i := 0; i < 4; i++ {
				setAlternateEyeColor(leds)
			}

			time.Sleep(500 * time.Millisecond)
		}
	}
}

func setAlternateEyeColor(leds ws2812.Device) {
	leds.WriteColors([]color.RGBA{
		RGBA_TURQUOISE,
		RGBA_BLACK,
	})

	time.Sleep(50 * time.Millisecond)

	leds.WriteColors([]color.RGBA{
		RGBA_BLACK,
		RGBA_TURQUOISE,
	})

	time.Sleep(50 * time.Millisecond)
}

func setNormalEyeColor(leds ws2812.Device) {
	leds.WriteColors([]color.RGBA{
		RGBA_TURQUOISE,
		RGBA_TURQUOISE,
	})

	time.Sleep(300 * time.Millisecond)

	leds.WriteColors([]color.RGBA{
		RGBA_BLACK,
		RGBA_BLACK,
	})

	time.Sleep(300 * time.Millisecond)
}
