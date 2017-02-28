package screen

import "sort"

// Window represents a buffer that may be attached and detached from other
// windows. When drawing the screen, we first start at the RootWindow and
// then iterate down over its children.
type Window struct {
	// width and height of this window
	w, h int

	// backing buffer
	runes  []rune
	fg, bg []uint32

	// parent, if attached
	p *Window

	// x,y offset from parent origin
	px, py int

	// z offset within parent (lower value = higher on screen)
	pz uint

	// children, sorted by Z order
	c []*Window

	// damaged area (relative to our origin)
	dx1, dy1, dx2, dy2 int
}

// RootWindow represents the background of the application. It ignores calls
// such as Attach and Resize.
var RootWindow *Window

// NewWindow returns a new window object with the given size. The window is not
// attached.
func NewWindow(w, h int) *Window {
	n := w * h
	win := &Window{
		w:     w,
		h:     h,
		runes: make([]rune, n),
		fg:    make([]uint32, n),
		bg:    make([]uint32, n),
		dx1:   0,
		dy1:   0,
		dx2:   w,
		dy2:   h,
	}
	for i := 0; i < n; i++ {
		win.runes[i] = ' '
	}
	for i := 0; i < n; i++ {
		win.bg[i] = Background
	}
	return win
}

// Attach will attach the given window to another. It implicitly detaches the
// given window first if it is already attached. This function may be used to
// move the position of a window and to alter its z-order by re-attaching to the
// same parent. This call is ignored on the root window.
func (w *Window) Attach(parent *Window, px, py int, pz uint) {
	if w == RootWindow {
		return
	}

	// detach from current parent if necessary
	if w.p != nil {
		w.Detach()
	}

	// attach to new parent
	w.p = parent
	w.px = px
	w.py = py
	w.pz = pz
	i := sort.Search(len(parent.c), func(i int) bool {
		return parent.c[i].pz <= pz
	})
	parent.c = append(parent.c, nil)
	copy(parent.c[i+1:], parent.c[i:])
	parent.c[i] = w

	// mark damaged area in parent
	parent.damaged(px, py, px+w.w, py+w.h)
}

// Detach will detach the given window. This call is ignored on the root window.
func (w *Window) Detach() {
	if w == RootWindow || w.p == nil {
		return
	}

	p := w.p
	w.p = nil

	// remove us from our parent's children
	for i := range p.c {
		if p.c[i] == w {
			copy(p.c[i:], p.c[i+1:])
			p.c[len(p.c)-1] = nil
			p.c = p.c[:len(p.c)-1]
			break
		}
	}

	// mark damaged area in parent
	p.damaged(w.px, w.py, w.px+w.w, w.py+w.h)
}

// Clear clears the contents of a window to the default screen background
// colour. All contents are lost.
func (w *Window) Clear() {
	for pos := range w.runes {
		w.runes[pos] = ' '
	}
	for pos := range w.fg {
		w.fg[pos] = White // doesn't really matter
	}
	for pos := range w.bg {
		w.bg[pos] = Background
	}
	w.damaged(0, 0, w.w, w.h)
}

// Size reports the current size of the window.
func (w *Window) Size() (width, height int) {
	return w.w, w.h
}

// damaged reports damage to a window. This will recurse all the way to the root
// window if visible. Damage coordinates are relative to w's origin, and are
// clipped automatically to the correct size.
func (w *Window) damaged(dx1, dy1, dx2, dy2 int) {
	dx1 = clip(dx1, 0, w.w-1)
	dy1 = clip(dy1, 0, w.h-1)
	dx2 = clip(dx2, 0, w.w-1)
	dy2 = clip(dy2, 0, w.h-1)

	// Update our damage markers. If we update any marker, take a note that
	// we will report this damage to our parent too. If not, then we need
	// take no further action.
	var recurse bool
	if dx1 < w.dx1 {
		w.dx1 = dx1
		recurse = true
	}
	if dy1 < w.dy1 {
		w.dy1 = dy1
		recurse = true
	}
	if dx2 > w.dx2 {
		w.dx2 = dx2
		recurse = true
	}
	if dy2 > w.dy2 {
		w.dy2 = dy2
		recurse = true
	}

	switch {
	case !recurse:
		// damage was out of bounds, or already encompassed
	case w.p == nil:
		// root window, or unattached; no need to recurse
	default:
		// report the damage to our parent, translating coordinates so
		// that they are relative to the parent's origin.
		w.p.damaged(w.px+w.dx1, w.py+w.dy1, w.px+w.dx2, w.py+w.dy2)
	}
}

// clearDamage recursively clears damage markers on the given window and its
// children. It is intended to be called after flushing any updates to the
// screen.
func (w *Window) clearDamage() {
	w.dx1, w.dy1, w.dx2, w.dy2 = w.w, w.h, 0, 0
	for _, c := range w.c {
		c.clearDamage()
	}
}

// resolveChar returns the character we would like to draw at the given offset,
// relative to our own origin.
func (w *Window) resolveChar(x, y int) (r rune, fg, bg uint32) {
	for _, c := range w.c {
		cx, cy := x-c.px, y-c.py
		if cx >= 0 && cx < c.w && cy >= 0 && cy < c.h {
			return c.resolveChar(cx, cy)
		}
	}
	pos := y*w.w + x
	return w.runes[pos], w.fg[pos], w.bg[pos]
}
