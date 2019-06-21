package freq

import (
	"errors"
	"strconv"
	"strings"
)

type Band int

const (
	BandUNKNOWN Band = iota
	Band160M
	Band80M
	Band40M
	Band20M
	Band15M
	Band10M
)

func (a Band) String() string {
	switch a {
	case Band160M:
		return "160M"
	case Band80M:
		return "80M"
	case Band40M:
		return "40M"
	case Band20M:
		return "20M"
	case Band15M:
		return "15M"
	case Band10M:
		return "10M"

	default:
		return "BandUNKNOWN"
	}
}

func GetBand(kHz string) (Band, error) {
	kHz = strings.TrimSpace(kHz)
	band := BandUNKNOWN
	freq, err := strconv.ParseFloat(kHz, 32)
	if err != nil {
		return band, err
	}
	switch {
	case freq >= 1800 && freq < 2000:
		band = Band160M
	case freq >= 3500 && freq < 4000:
		band = Band80M
	case freq >= 7000 && freq < 7300:
		band = Band40M
	case freq >= 14000 && freq < 14350:
		band = Band20M
	case freq >= 21000 && freq < 21450:
		band = Band15M
	case freq >= 28000 && freq < 29700:
		band = Band10M
	default:
		return band, errors.New("Unsupported band, freq=" + kHz + " kHz")
	}
	return band, nil
}
