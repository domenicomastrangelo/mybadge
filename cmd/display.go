package main

import (
	"image/color"
	"machine"

	"tinygo.org/x/drivers/st7789"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freemono"
)

func setupDisplay() st7789.Device {
	display := initDisplay()
	display = configureDisplay(display)

	resetScreen(display)

	return display
}

func initDisplay() st7789.Device {
	// Setup the screen pins
	machine.SPI0.Configure(machine.SPIConfig{
		Frequency: 8000000,
		Mode:      0,
	})

	display := st7789.New(machine.SPI0,
		machine.TFT_RST,       // TFT_RESET
		machine.TFT_WRX,       // TFT_DC
		machine.TFT_CS,        // TFT_CS
		machine.TFT_BACKLIGHT) // TFT_LITE

	return display
}

func configureDisplay(display st7789.Device) st7789.Device {
	display.Configure(st7789.Config{
		Rotation: st7789.ROTATION_90,
		Height:   320,
	})

	return display
}

func writeToScreen(display st7789.Device, showTwitter bool) {
	resetScreen(display)

	display.FillRectangle(40, 5, 240, 45, color.RGBA{R: 255, G: 255, B: 255, A: 255})
	tinyfont.WriteLine(&display, &freemono.Bold24pt7b, 48, 40, "DOMENICO", color.RGBA{R: 0, G: 0, B: 0, A: 1})

	tinyfont.WriteLine(&display, &freemono.Bold12pt7b, 10, 75, "LINKEDIN:", color.RGBA{R: 0, G: 144, B: 177, A: 1})
	tinyfont.WriteLine(&display, &freemono.Regular12pt7b, 10, 95, "Domenico Mastrangelo", color.RGBA{R: 255, G: 255, B: 255, A: 1})

	if showTwitter {
		tinyfont.WriteLine(&display, &freemono.Bold12pt7b, 10, 135, "TWITTER:", color.RGBA{R: 0, G: 172, B: 238, A: 1})
		tinyfont.WriteLine(&display, &freemono.Regular12pt7b, 10, 155, "DomeMastrangelo", color.RGBA{R: 255, G: 255, B: 255, A: 1})
	}

	tinyfont.WriteLine(&display, &freemono.Bold12pt7b, 10, 195, "WORK:", color.RGBA{R: 255, G: 25, B: 71, A: 1})
	tinyfont.WriteLine(&display, &freemono.Regular12pt7b, 10, 215, "heyworld GmbH ( FFM )", color.RGBA{R: 255, G: 255, B: 255, A: 1})
}

func resetScreen(display st7789.Device) {
	display.FillScreen(color.RGBA{0, 0, 0, 255})
}
