package terminalasciiart

import (
	"errors"
	"os"
)

type textSet struct {
	*Text
}

func (t *textSet) set() error {
	args := os.Args[1:]

	if len(args) != 1 {
		return errors.New("set() text must be wrapped in \"\"")
	}

	t.input = args[0]

	return nil
}
