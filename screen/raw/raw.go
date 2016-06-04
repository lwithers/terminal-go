/*
Package raw implements raw output codes for the terminal.
*/
package raw

import (
	"fmt"
	"io"
)

const (
	csi = "\033["
)

func ScreenSetup(w io.Writer) {
	fmt.Fprintf(w, "%s?1049h", csi)
}

func ScreenTeardown(w io.Writer) {
	fmt.Fprintf(w, "%s?1049l", csi)
}
