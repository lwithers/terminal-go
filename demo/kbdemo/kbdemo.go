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

	fmt.Println("Press ‘q’ to quit.\r")
EventLoop:
	for {
		select {
		case key := <-keych:
			fmt.Printf("Got key %6d\r\n", key)
			if key == 'q' {
				break EventLoop
			}
		case <-sigch:
			break EventLoop
		}
	}

	keyboard.Stop()
}
