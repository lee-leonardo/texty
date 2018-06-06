package main

// #include <unistd.h>
import "C"

import (
	"fmt"
	"syscall"

	"github.com/pkg/term/termios"
)

/*
Constants for the CLI are pulled primarily from this source. https://godoc.org/github.com/cznic/ccir/libc/unistd
*/

const (
	stdIn uintptr = 0
)

var (
	term = syscall.Termios{}
)

func getAttributes() error {
	err := termios.Tcgetattr(stdIn, &term)

	if err != nil {
		fmt.Println(err)
		return err
	}
}

func setDefaultAttributes() error {
	err := getAttributes()

	if err != nil {
		return err
	}

	// syscall.ECHO

	err = termios.Tcsetattr(stdIn, syscall.TCSAFLUSH, &term)

	if err != nil {
		fmt.Println(err)
		return err
	}
}
