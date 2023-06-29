package main

import (
	"mybadge/internal/virtualdisplay"

	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freemono"
)

func setupDisplay() *virtualdisplay.VirtualDisplay {
	virtualDisplay := initDisplay()
	virtualDisplay.ResetScreen()

	return virtualDisplay
}

func initDisplay() *virtualdisplay.VirtualDisplay {
	virtualDisplay := virtualdisplay.New(320, 240, RGBA_BLACK)

	return &virtualDisplay
}

func writeToScreen(virtualDisplay *virtualdisplay.VirtualDisplay, showTwitter bool) {
	virtualDisplay.Draw(func() {

		virtualDisplay.ResetScreen()

		virtualDisplay.Display.FillRectangle(40, 5, 240, 45, RGBA_WHITE)

		tinyfont.WriteLine(&virtualDisplay.Display, &freemono.Bold24pt7b, 48, 40, "DOMENICO", RGBA_BLACK)

		tinyfont.WriteLine(&virtualDisplay.Display, &freemono.Bold12pt7b, 10, 75, "LINKEDIN:", RGBA_LINKEDIN)
		tinyfont.WriteLine(&virtualDisplay.Display, &freemono.Regular12pt7b, 10, 95, "Domenico Mastrangelo", RGBA_WHITE)

		if showTwitter {
			tinyfont.WriteLine(&virtualDisplay.Display, &freemono.Bold12pt7b, 10, 135, "TWITTER:", RGBA_TWITTER)
			tinyfont.WriteLine(&virtualDisplay.Display, &freemono.Regular12pt7b, 10, 155, "DomeMastrangelo", RGBA_WHITE)
		}

		tinyfont.WriteLine(&virtualDisplay.Display, &freemono.Bold12pt7b, 10, 195, "WORK:", RGBA_RED)
		tinyfont.WriteLine(&virtualDisplay.Display, &freemono.Regular12pt7b, 10, 215, "heyworld GmbH ( FFM )", RGBA_WHITE)
	})
}
