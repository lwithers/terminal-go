/*
kbdemo prints details about keypresses.
*/
package main

import (
	"fmt"
	"os"
	"os/signal"

	"golang.org/x/sys/unix"

	"github.com/lwithers/terminal-go/keyboard"
)

func main() {
	keyboard.Init()
	keych := keyboard.StartReader()
	sigch := make(chan os.Signal)
	signal.Notify(sigch, unix.SIGINT, unix.SIGTERM)

	fmt.Println("Press ‘q’ or Control-C to quit.\r")
EventLoop:
	for {
		select {
		case key := <-keych:
			fmt.Printf("Got key %6d\r\n", key)
			switch key {
			case 'q', keyboard.Key_Ctrl_C:
				break EventLoop
			}
		case <-sigch:
			break EventLoop
		}
	}

	keyboard.Stop()
}
