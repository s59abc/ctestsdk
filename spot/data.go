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
	dx       string
	de       string
	freq     string
	raw      string
	comments string
	// TODO;
	Source   string //Source (origin) of the spot
	IsRbn    bool
	BAND     freq.Band
	MODE     Mode
	DxQTH    geo.QTH
	DxCtyDta cty.Dta
	DeQTH    geo.QTH
	DeCtyDta cty.Dta
}

var ignore error = errors.New("it is not a spot, ignore it")

///////////
//regex
var isSpotRegex *regexp.Regexp = regexp.MustCompile(`^DX de \w`)
var splitSpotByColonRegex *regexp.Regexp = regexp.MustCompile(":")
var splitSpotBySpaceRegex *regexp.Regexp = regexp.MustCompile(`[ ]+`)
var deCallSignRegex *regexp.Regexp = regexp.MustCompile(`[0-9A-Za-z/]{3,}`)

//
///////

// This function create Data object representing a Spot. It is not necessary that creation is
// successful. Any string can be passed to that function and if it is a spot, Data returned has
// representing valid Spot and error is nil.
// Nothing is wrong if error is not nil. Normally that is just sign to caller function
// that something else was passed to this function .e.g. WWV.
// If it is a RBN or Dx Cluster AND returned error is not nil that is sing that it is spot but it
// can not be decoded into Data cos there is bug in this function or spot is really wrong formatted.
func NewSpot(rawData string, source string) (Data, error) {
	//DX de S50ARX-#:   7035.3  LA9QJA       CW 16 dB 16 WPM CQ             1553Z
	if matched := isSpotRegex.MatchString(rawData); matched {

		data := Data{raw: rawData} //raw: DX de S50ARX-#:   7035.3  LA9QJA       CW 16 dB 16 WPM CQ             1553Z
		s := rawData[6:]           //s:         S50ARX-#:   7035.3  LA9QJA       CW 16 dB 16 WPM CQ             1553Z
		//		fmt.Println("s:", s)
		ss := splitSpotByColonRegex.Split(s, 2)
		if len(ss) != 2 {
			return Data{}, errors.New("Not a regular spot; Unexpected split by colon! " + rawData)
		}
		//fmt.Println("ss[0]=", ss[0])
		//fmt.Println("ss[1]=",ss[1])
		//ss[0]= S50ARX-#
		data.IsRbn = ss[0][len(ss[0])-1] == '#'
		//spot sender
		data.de = deCallSignRegex.FindString(ss[0])
		if data.de == "" {
			return Data{}, errors.New("Not a regular spot; DE is wrong formatted! " + rawData)
		}

		//ss[1]=    7035.3  LA9QJA       CW 16 dB 16 WPM CQ             1553Z
		sss := splitSpotBySpaceRegex.Split(ss[1], 4)
		if len(sss) != 4 {
			return Data{}, errors.New("Not a regular spot; Unexpected split by space! " + rawData)
		}

		//freq, band
		data.freq = sss[1]
		if b, e := freq.GetBand(data.freq); e == nil {
			data.BAND = b
		}
		//		BAND, _ = freq.GetBand(data.freq)
		data.dx = sss[2]
		data.comments = sss[3]

		//fmt.Println(len(sss))
		//for _,j := range sss {
		//	fmt.Println(j)
		//}

		data.raw = rawData
		data.Source = source
		return data, nil

	} else {
		return Data{}, ignore
	}
}

func NewSpotPoc(rawData string, source string) (Data, error) {
	//DX de S50ARX-#:   7035.3  LA9QJA       CW 16 dB 16 WPM CQ             1553Z
	if matched := isSpotRegex.MatchString(rawData); matched {

		data := Data{raw: rawData} //raw: DX de S50ARX-#:   7035.3  LA9QJA       CW 16 dB 16 WPM CQ             1553Z
		s := rawData[6:]           //s:         S50ARX-#:   7035.3  LA9QJA       CW 16 dB 16 WPM CQ             1553Z
		fmt.Println("s:", s)
		ss := splitSpotByColonRegex.Split(s, 2)
		if len(ss) != 2 {
			return Data{}, errors.New("Not a regular spot; Unexpected split by colon! " + rawData)
		}
		fmt.Println("ss[0]=", ss[0])
		fmt.Println("ss[1]=", ss[1])
		//ss[0]= S50ARX-#
		data.IsRbn = ss[0][len(ss[0])-1] == '#'
		//spot sender
		data.de = deCallSignRegex.FindString(ss[0])
		if data.de == "" {
			return Data{}, errors.New("Not a regular spot; DE is wrong formatted! " + rawData)
		}

		//ss[1]=    7035.3  LA9QJA       CW 16 dB 16 WPM CQ             1553Z
		sss := splitSpotBySpaceRegex.Split(ss[1], 4)
		if len(sss) != 4 {
			return Data{}, errors.New("Not a regular spot; Unexpected split by space! " + rawData)
		}
		data.freq = sss[1]
		data.dx = sss[2]
		data.comments = sss[3]

		fmt.Println(len(sss))
		for _, j := range sss {
			fmt.Println(j)
		}

		data.Source = source
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
