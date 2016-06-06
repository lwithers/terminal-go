package raw

import (
	"fmt"
	"io"
)

func CursorHide(w io.Writer) {
	fmt.Fprintf(w, "%s?25l", csi)
}

func CursorShow(w io.Writer) {
	fmt.Fprintf(w, "%s?25h", csi)
}

func CursorMove(w io.Writer, x, y int) {
	fmt.Fprintf(w, "%s%d;%dH", csi, y+1, x+1)
}
