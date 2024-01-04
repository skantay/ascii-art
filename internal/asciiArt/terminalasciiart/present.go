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

	if int(w)*8 < len(t.output) {
		return errors.New("present() text is too loooong")
	}

	if t.output == "" {
		return nil
	}

	fmt.Print(t.output)

	return nil
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
