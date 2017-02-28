/*
boxchardemo draws box characters on the screen, with a nice grid.
*/
package main

import (
	"time"

	"github.com/lwithers/terminal-go/keyboard"
	"github.com/lwithers/terminal-go/screen"
)

var (
	clockFg = screen.Colour(38, 139, 210)
)

func main() {
	screen.Background = screen.White
	screen.Init()
	keyboard.Init()

	win := screen.NewWindow(60, 10)
	drawBorder(win, screen.White, screen.Black)
	win.DrawString("Use arrow keys to move",
		screen.White, screen.Black, 2, 1, 0)
	wx, wy := 10, 10
	win.Attach(screen.RootWindow, wx, wy, 0)

	subWin := screen.NewWindow(35, 3)
	drawBorder(subWin, clockFg, screen.White)
	subWin.Attach(win, 5, 5, 0)

	tick := time.Tick(time.Second)
	keych := keyboard.StartReader()
MainLoop:
	for {
		screen.Flush()

		var moved bool
		select {
		case key := <-keych:
			switch key {
			case keyboard.Key_Ctrl_C, 'q', 'Q':
				break MainLoop
			case keyboard.Key_Left:
				wx--
				moved = true
			case keyboard.Key_Right:
				wx++
				moved = true
			case keyboard.Key_Up:
				wy--
				moved = true
			case keyboard.Key_Down:
				wy++
				moved = true
			}

		case t := <-tick:
			ts := t.Format(time.RFC3339)
			subWin.DrawString(ts, clockFg, screen.White, 1, 1, 0)
		}

		if moved {
			win.Attach(screen.RootWindow, wx, wy, 0)
		}
	}

	keyboard.Stop()
	screen.Stop()
}

var winBorders = [3][3]rune{
	{'┌', '─', '┐'},
	{'│', ' ', '│'},
	{'└', '─', '┘'},
}

func drawBorder(win *screen.Window, fg, bg uint32) {
	w, h := win.Size()
	r := make([]rune, w)

	for i := range r {
		r[i] = '─'
	}
	r[0], r[w-1] = '┌', '┐'
	win.DrawRunes(r, fg, bg, 0, 0)
	r[0], r[w-1] = '└', '┘'
	win.DrawRunes(r, fg, bg, 0, h-1)

	for i := range r {
		r[i] = ' '
	}
	r[0], r[w-1] = '│', '│'
	for y := 1; y < h-1; y++ {
		win.DrawRunes(r, fg, bg, 0, y)
	}
}
