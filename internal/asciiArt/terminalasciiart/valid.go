package terminalasciiart

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"io"
	"os"
)

type textValid struct {
	*Text
}

func (t *textValid) validate() error {
	if err := validateBanner(t.Banner); err != nil {
		return err
	}

	if err := validateInput(t.input); err != nil {
		return err
	}

	return nil
}

// This function has unit testing
func validateBanner(banner string) error {
	hashes := make(map[string]string)
	hashes["standard.txt"] = "e194f1033442617ab8a78e1ca63a2061f5cc07a3f05ac226ed32eb9dfd22a6bf"
	hashes["shadow.txt"] = "26b94d0b134b77e9fd23e0360bfd81740f80fb7f6541d1d8c5d85e73ee550f73"
	hashes["thinkertoy.txt"] = "64285e4960d199f4819323c4dc6319ba34f1f0dd9da14d07111345f5d76c3fa3"

	if _, ok := hashes[banner]; !ok {
		return errors.New("banner not found")
	}

	cwd, _ := os.Getwd()
	cwd = trimCwd(cwd)
	file, err := os.Open(cwd + "/banners/" + banner)
	if err != nil {
		return err
	}
	defer file.Close()

	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return err
	}

	computedHash := hex.EncodeToString(hash.Sum(nil))

	if val := hashes[banner]; val != computedHash {
		return errors.New("validate() file is broken")
	}

	return nil
}

// This function has unit testing
func validateInput(input string) error {
	for _, v := range input {
		if 126 < v {
			return errors.New("validate() non ascii char found")
		}
	}

	return nil
}
