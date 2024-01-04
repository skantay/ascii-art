package asciiart

import (
	"fmt"

	"github.com/skantay/ascii-art/internal/asciiArt/terminalasciiart"
)

func Run() error {
	text := &terminalasciiart.Text{
		Banner: "standard.txt",
	}

	service := terminalasciiart.NewArtService(text)

	if err := service.Run(); err != nil {
		return fmt.Errorf("service \"%w", err)
	}

	return nil
}
