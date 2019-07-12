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
	FT4
	FT8
	PSK
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
	case "FT4":
		mode = FT4
	case "FT8":
		mode = FT8
	case "PSK":
		mode = PSK
	default:
		if strings.HasPrefix(input, "PSK") {
			mode = PSK
		} else {
			return mode, errors.New("Unsupported mode, input=" + input)
		}
	}
	return mode, nil
}

func (a Mode) String() string {
	switch a {
	case CW:
		return "CW"
	case SSB:
		return "SSB"
	case RTTY:
		return "RTTY"
	case FM:
		return "FM"
	case FT4:
		return "FT4"
	case FT8:
		return "FT8"
	case PSK:
		return "PSK"
	default:
		return "UNKNOWN"
	}
}
