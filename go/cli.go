package main

/*
Constants for the CLI are pulled primarily from this source. https://godoc.org/github.com/cznic/ccir/libc/unistd
*/

import (
	"fmt"
	"syscall"

	"github.com/pkg/term/termios"
)

const (
	stdInFileNo = 0 //STDIN_FILENO
)

var (
	term = syscall.Termios{}
)

func setDefaultAttributes() {
	err := termios.Tcgetattr(stdInFileNo, &term)

	if err != nil {
		fmt.Println(err)
	}
}
