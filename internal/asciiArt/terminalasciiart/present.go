package terminalasciiart

import (
	"errors"
	"fmt"
	"syscall"
	"unsafe"
)

type textPresent struct {
	*Text
}

func (t *textPresent) present() error {
	_, w := getTerminalSize()

	width := getWidth(t.output)

	if int(w) <= width {
		return errors.New("present() text is too loooong")
	}

	if t.output == "" {
		return nil
	}

	fmt.Print(t.output)

	return nil
}

func getWidth(output string) int {
	var bufLen int
	var width int

	for i := 0; i < len(output); i++ {
		if rune(output[i]) == '\n' {
			if bufLen > width {
				width = bufLen
			}
			bufLen = 0
		}
		bufLen++
	}

	return width
}

type winsize struct {
	Row    uint16
	Col    uint16
	Xpixel uint16
	Ypixel uint16
}

func getTerminalSize() (uint, uint) {
	ws := &winsize{}
	retCode, _, _ := syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(syscall.Stdin),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(ws)))

	if int(retCode) == -1 {
		panic("Error getting terminal size")
	}
	return uint(ws.Row), uint(ws.Col)
}
