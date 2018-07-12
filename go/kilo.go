package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	ESC   = 27
	ENTER = 10
	Q     = 113
)

func main() {
	var sig int = 0
	reader := bufio.NewReader(os.Stdin)

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

		fmt.Printf("%d %c\n", rune, rune)
	}
}
