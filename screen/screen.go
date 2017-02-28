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

// Init will create and draw RootWindow. You may wish to set the default
// Background colour first. This function will also clear the screen contents,
// disable scrollback mode, and disable the cursor.
func Init() {
	raw.ScreenSetup(os.Stdout)
	raw.CursorHide(os.Stdout)

	// TODO: register screen size notifier, hook up to RootWindow
	RootWindow = NewWindow(GetSize())
}

// Stop returns the screen contents to how they were before program execution,
// restoring scrollback.
func Stop() {
	RootWindow = nil
	raw.CursorShow(os.Stdout)
	raw.ScreenTeardown(os.Stdout)
}

// clip returns x modified such that min ≤ x ≤ max. More rigourously:
//  · if x < min, returns min
//  · if x > max, return max
//  · else returns x
func clip(x, min, max int) int {
	if x < min {
		return min
	}
	if x > max {
		return max
	}
	return x
}
