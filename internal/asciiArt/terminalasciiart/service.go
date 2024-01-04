package terminalasciiart

import (
	"errors"
	"fmt"
	"strings"
)

type setter interface {
	set() error
}

type validator interface {
	validate() error
}

type producer interface {
	produce()
}

type presenter interface {
	present() error
}

type artService struct {
	setter    setter
	validator validator
	producer  producer
	presenter presenter
}

type Text struct {
	input, output, Banner string
}

func NewArtService(t *Text) *artService {
	s := &textSet{t}
	v := &textValid{t}
	prod := &textProduce{t}
	pres := &textPresent{t}

	return &artService{
		setter:    s,
		validator: v,
		producer:  prod,
		presenter: pres,
	}
}

func (s *artService) Run() error {
	errorMSG := errors.New("terminal ascii art\" error from")

	if err := s.setter.set(); err != nil {
		return fmt.Errorf("%s: %s", errorMSG, err)
	}
	if err := s.validator.validate(); err != nil {
		return fmt.Errorf("%s: %s", errorMSG, err)
	}
	s.producer.produce()

	if err := s.presenter.present(); err != nil {
		return fmt.Errorf("%s: %s", errorMSG, err)
	}

	return nil
}

func trimCwd(cwd string) string {
	lastIndex := strings.LastIndex(cwd, "ascii-art")

	return cwd[:lastIndex+len("ascii-art")]
}
