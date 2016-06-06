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

// Init will clear the screen contents, disable scrollback mode, and disable
// the cursor.
func Init() {
	raw.ScreenSetup(os.Stdout)
	buf.resetBuffer()
}

// Stop returns the screen contents to how they were before program execution,
// restoring scrollback.
func Stop() {
	raw.ScreenTeardown(os.Stdout)
}
