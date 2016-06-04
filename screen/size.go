package screen

import (
	"os"
	"os/signal"
	"sync"

	"golang.org/x/sys/unix"

	"github.com/lwithers/terminal-go/screen/raw"
)

var (
	screenSizeLock   sync.RWMutex
	screenW, screenH int
	screenNotifyLock sync.Mutex
	winchNotifiers   []ScreenSizeNotifier
)

// ScreenSizeNotifier objects are notified whenever the screen size changes.
type ScreenSizeNotifier interface {
	ScreenSizeNotify(w, h int)
}

// GetSize queries the current screen size.
func GetSize() (w, h int) {
	screenSizeLock.RLock()
	w, h = screenW, screenH
	screenSizeLock.RUnlock()
	return
}

// RegisterSizeNotifier adds a screen size change notifier to a list which will
// be called if the screen size changes.
func RegisterSizeNotifier(n ScreenSizeNotifier) {
	screenNotifyLock.Lock()
	winchNotifiers = append(winchNotifiers, n)
	screenNotifyLock.Unlock()
}

func init() {
	screenW, screenH = raw.GetWinSize()
	if screenW <= 0 || screenH <= 0 {
		panic("unable to query screen size")
	}

	winchNotifiers = make([]ScreenSizeNotifier, 0)

	c := make(chan os.Signal)
	signal.Notify(c, unix.SIGWINCH)
	go handleWinch(c)
}

func handleWinch(c chan os.Signal) {
	for {
		// will block until signal is received
		<-c

		// read new size, ignore if query fails
		w, h := raw.GetWinSize()
		if w <= 0 || h <= 0 {
			continue
		}

		// update variables we report with
		screenSizeLock.Lock()
		screenW, screenH = w, h
		screenSizeLock.Unlock()

		// notify any registered watchers
		screenNotifyLock.Lock()
		for _, n := range winchNotifiers {
			n.ScreenSizeNotify(w, h)
		}
		screenNotifyLock.Unlock()
	}
}
