package freq

import (
	"errors"
	"strconv"
	"strings"
)

type Band int

const (
	UNKNOWN Band = iota
	M160
	M80
	M40
	M20
	M15
	M10
)

func (a Band) String() string {
	switch a {
	case M160:
		return "160M"
	case M80:
		return "80M"
	case M40:
		return "40M"
	case M20:
		return "20M"
	case M15:
		return "15M"
	case M10:
		return "10M"

	default:
		return "UNKNOWN"
	}
}

func GetBand(kHz string) (Band, error) {
	kHz = strings.TrimSpace(kHz)
	band := UNKNOWN
	freq, err := strconv.ParseFloat(kHz, 32)
	if err != nil {
		return band, err
	}
	switch {
	case freq >= 1800 && freq < 2000:
		band = M160
	case freq >= 3500 && freq < 4000:
		band = M80
	case freq >= 7000 && freq < 7300:
		band = M40
	case freq >= 14000 && freq < 14350:
		band = M20
	case freq >= 21000 && freq < 21450:
		band = M15
	case freq >= 28000 && freq < 29700:
		band = M10
	default:
		return band, errors.New("Unsupported band, freq=" + kHz + " kHz")
	}
	return band, nil
}
