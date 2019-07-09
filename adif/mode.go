package adif

import (
	"errors"
	"strings"
)

type Mode int

const (
	UNKNOWN Mode = iota
	CW
	SSB
	RTTY
	FM
)

func GetMode(input string) (Mode, error) {
	mode := UNKNOWN
	input = strings.TrimSpace(input)
	input = strings.ToUpper(input)
	switch input {
	case "CW":
		mode = CW
	case "SSB":
		mode = SSB
	case "RTTY":
		mode = RTTY
	case "FM":
		mode = FM
	default:
		return mode, errors.New("Unsupported mode, input=" + input)
	}
	return mode, nil
}

func (a Mode) String() string {
	switch a {
	case CW:
		return "CQ"
	case SSB:
		return "SSB"
	case RTTY:
		return "RTTY"
	case FM:
		return "FM"
	default:
		return "UNKNOWN"
	}
}
