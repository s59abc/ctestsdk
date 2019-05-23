package spot

import (
	"ctestsdk/cty"
	"ctestsdk/freq"
	"ctestsdk/geo"
	"errors"
	"fmt"
	"regexp"
)

type Data struct {
	dx   string
	de   string
	freq string
	raw  string
	// TODO;
	BAND     freq.Band
	MODE     Mode
	DxQTH    geo.QTH
	DxCtyDta cty.Dta
	DeQTH    geo.QTH
	DeCtyDta cty.Dta
}

var ignore error = errors.New("") //it is not a spot, ignore it

func NewSpot(rawData string) (Data, error) {
	if matched, _ := regexp.MatchString(`^DX de \w{1}`, rawData); matched {
		data := Data{raw: rawData}
		s := rawData[6:]
		fmt.Println(s)
		return data, nil

	} else {
		return Data{}, ignore
	}
}

func (a Data) DX() string {
	return a.dx
}

func (a Data) DE() string {
	return a.de
}

func (a Data) FREQ() string {
	return a.freq
}

func (a Data) Raw() string {
	return a.raw
}
