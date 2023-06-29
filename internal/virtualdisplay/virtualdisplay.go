package virtualdisplay

import (
	"image/color"

	"tinygo.org/x/drivers/st7789"

	"machine"
)

var (
	displayNumber = 0
	RGBA_BLACK    = color.RGBA{0, 0, 0, 0}
)

type VirtualDisplay struct {
	ID              int8
	Width           int16
	Height          int16
	BackgroundColor color.RGBA
	Display         st7789.Device
}

func New(width int16, height int16, BackgroundColor color.RGBA) VirtualDisplay {
	// Setup the screen pins
	machine.SPI0.Configure(machine.SPIConfig{
		Frequency: 8000000,
		Mode:      0,
	})

	display := st7789.New(machine.SPI0,
		machine.TFT_RST, // TFT_RESET
		machine.TFT_WRX, // TFT_DC
		machine.TFT_CS,  // TFT_CS
		machine.TFT_BACKLIGHT)

	display.Configure(st7789.Config{
		Rotation: st7789.ROTATION_90,
		Height:   width,
	})

	displayNumber++

	return VirtualDisplay{
		Width:           width,
		Height:          height,
		BackgroundColor: color.RGBA{0, 0, 0, 0},
		Display:         display,
	}
}

func (v *VirtualDisplay) Draw(callback func()) {
	callback()
}

func (v *VirtualDisplay) ResetScreen() {
	v.Display.FillScreen(RGBA_BLACK)
}
