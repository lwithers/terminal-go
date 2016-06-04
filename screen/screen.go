/*
Package screen implements a buffer for managing a screen of content and handles
only sending redraw commands for data that changes. It also allows registration
of callbacks for when the screen size changes.
*/
package screen

import (
	"os"

	"github.com/lwithers/terminal-go/screen/raw"
)

func Init() {
	raw.ScreenSetup(os.Stdout)
	buf.resetBuffer()
}

func Stop() {
	raw.ScreenTeardown(os.Stdout)
}
