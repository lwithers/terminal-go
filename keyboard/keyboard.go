/*
Package keyboard provides a channel for notifying the application of keypresses.
It operates the keyboard in raw mode, in which characters are not interpreted
and not echoed to the screen. It provides constants for various keypress
combinations.
*/
package keyboard

import "github.com/lwithers/terminal-go/keyboard/kbraw"

// Init will switch the keyboard into raw mode, so that we can interpret
// keypresses properly. This will switch off keyboard echo to the terminal.
func Init() {
	kbraw.RawSetup()
}

// Stop restores the program's original keyboard mode.
func Stop() {
	kbraw.RawTeardown()
}
