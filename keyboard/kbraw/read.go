package kbraw

import "os"

func reader(c chan byte) {
	b := make([]byte, 1)

	for {
		_, err := os.Stdin.Read(b)
		if err != nil {
			panic(err)
		}
		c <- b[0]
	}
}

func StartReader() chan byte {
	c := make(chan byte, 100)
	go reader(c)
	return c
}
