package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	ESC   = 0
	ENTER = 0
	Q     = 0
)

func main() {
	var sig int
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Hello, World!")

	for sig == 0 {
		rune, _, err := reader.ReadRune()

		if err != nil {
			fmt.Print(err)
			sig = -1
		}

		if rune == ENTER {
			continue
		}

		if rune == ESC || rune == Q {
			sig = 1
			continue
		}

		fmt.Printf("%d %c", rune, rune)
	}
}
