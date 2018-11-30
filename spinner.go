package multispinner

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/leaanthony/synx"
)

// Status code constants.
const (
	runningStatus int = iota
	successStatus
	errorStatus
)

// Spinner defines a single s
type Spinner struct {
	message *synx.String
	status  *synx.Int
	group   *SpinnerGroup
}

// UpdateMessage updates the spinner message
func (s *Spinner) UpdateMessage(message string) {
	s.message.SetValue(message)
}

// Success marks spinner as success and update message
func (s *Spinner) Success(message string) {
	s.UpdateMessage(message)
	s.stop(successStatus)
}

// Error marks spinner as error and update message
func (s *Spinner) Error(message string) {
	s.UpdateMessage(message)
	s.stop(errorStatus)
}

func (s *Spinner) stop(status int) {
	s.status.SetValue(status)
	s.group.redraw()
	s.group.Done()
}

func (s *Spinner) sprint() string {
	line := fmt.Sprintf("%s %s", s.getSymbol(), s.message.GetValue())
	switch s.status.GetValue() {
	case successStatus:
		return color.HiGreenString(line)
	case errorStatus:
		return color.HiRedString(line)
	default:
		return line
	}
}

func (s *Spinner) getSymbol() string {
	switch s.status.GetValue() {
	case successStatus:
		return s.group.successSymbol
	case errorStatus:
		return s.group.errorSymbol
	default:
		return s.group.currentFrame()
	}
}
