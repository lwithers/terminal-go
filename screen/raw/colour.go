package raw

import (
	"fmt"
	"io"
)

func ColourFG(w io.Writer, r, g, b uint8) {
	fmt.Fprintf(w, "%s38;2;%d;%d;%dm", csi, r, g, b)
}

func ColourBG(w io.Writer, r, g, b uint8) {
	fmt.Fprintf(w, "%s48;2;%d;%d;%dm", csi, r, g, b)
}
