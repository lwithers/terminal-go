package raw

import (
	"fmt"
	"io"
)

func CursorHide(w io.Writer) {
	fmt.Fprintf(w, "%s?25l")
}

func CursorShow(w io.Writer) {
	fmt.Fprintf(w, "%s?25h")
}

func CursorMove(w io.Writer, x, y int) {
	fmt.Fprintf(w, "%s%d;%dH", csi, y+1, x+1)
}
