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
