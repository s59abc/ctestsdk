package spot

import (
	"ctestsdk/adif"
	"ctestsdk/geo"
	"errors"
	"fmt"
	"regexp"
	"strings"
)

type CtyDta struct {
	CountryName   string //Country Name
	PrimaryPrefix string
	AliasPrefix   string             //Primary or Alias DXCC Prefix without optional * indicator
	Continent     adif.ContinentEnum //2-letter Continent abbreviation
	CqZone        adif.CqzoneEnum    //CQ Zone
	ItuZone       adif.ItuzoneEnum   //ITU Zone
	LatLon        geo.LatLonDeg      //Latitude in degrees, + for North; Longitude in degrees, + for West
	TimeOffset    string             //Local time offset from GMT
}

func CtyDtaEqual(a, b CtyDta) bool {
	eq := a.CountryName == b.CountryName
	if eq {
		eq = a.PrimaryPrefix == b.PrimaryPrefix
	}
	if eq {
		eq = a.AliasPrefix == b.AliasPrefix
	}
	if eq {
		eq = a.Continent == b.Continent
	}
	if eq {
		eq = a.CqZone == b.CqZone
	}
	if eq {
		eq = a.ItuZone == b.ItuZone
	}
	if eq {
		eq = a.LatLon.Equal(b.LatLon)
	}
	if eq {
		eq = a.TimeOffset == b.TimeOffset
	}
	return eq
}

func (a CtyDta) String() string {
	return fmt.Sprintf("CountryName=%s, PrimaryPrefix=%s, AliasPrefix=%s, Continent=%s, CqZone=%d, ItuZone=%d, %s, TimeOffset=%s", a.CountryName, a.PrimaryPrefix, a.AliasPrefix, a.Continent.String(), a.CqZone, a.ItuZone, a.LatLon.String(), a.TimeOffset)

}

type Data struct {
	Dx       string
	De       string
	Freq     string
	Raw      string
	Comments string
	// TODO;
	Source   string //Source (origin) of the spot
	IsRbn    bool
	BAND     adif.Band
	MODE     adif.Mode
	DxQTH    geo.QTH
	DxCtyDta CtyDta
	DeQTH    geo.QTH
	DeCtyDta CtyDta
}

func (a Data) String() string {
	return fmt.Sprintf("\n%s\n"+
		"Dx:%s\n"+
		"De:%s\n"+
		"Freq:%s Band:%s Mode:%s\n"+
		"DeCtyDta:%s\n"+
		"DxCtyDta:%s\n"+
		"DeQTH:%s\n"+
		"DxQTH:%s\n"+
		"IsRBN:%t, Source:%s\n"+
		"Comments:%s\n", a.Raw, a.Dx, a.De, a.Freq, a.BAND.String(), a.MODE.String(), a.DeCtyDta.String(), a.DxCtyDta.String(), a.DeQTH, a.DxQTH, a.IsRbn, a.Source, a.Comments)
}

func DataEqual(a, b Data) bool {
	eq := a.Dx == b.Dx
	if eq {
		eq = a.De == b.De
	}
	if eq {
		eq = a.Freq == b.Freq
	}
	if eq {
		eq = a.Raw == b.Raw
	}
	if eq {
		eq = a.Comments == b.Comments
	}
	if eq {
		eq = a.Source == b.Source
	}
	if eq {
		eq = a.IsRbn == b.IsRbn
	}
	if eq {
		eq = a.BAND == b.BAND
	}
	if eq {
		eq = a.MODE == b.MODE
	}
	if eq {
		eq = geo.QthEqual(a.DxQTH, b.DxQTH)
	}
	if eq {
		eq = CtyDtaEqual(a.DxCtyDta, b.DxCtyDta)
	}
	if eq {
		eq = geo.QthEqual(a.DeQTH, b.DeQTH)
	}
	if eq {
		eq = CtyDtaEqual(a.DeCtyDta, b.DeCtyDta)
	}
	return eq
}

var ignore error = errors.New("it is not a spot, ignore it")

///////////
//regex
var isSpotRegex *regexp.Regexp = regexp.MustCompile(`^DX DE \w`)
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

	//DX De S50ARX-#:   7035.3  LA9QJA       CW 16 dB 16 WPM CQ             1553Z
	if matched := isSpotRegex.MatchString(strings.ToUpper(rawData)); matched {

		data := Data{Raw: rawData} //Raw: DX De S50ARX-#:   7035.3  LA9QJA       CW 16 dB 16 WPM CQ             1553Z
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
		data.De = deCallSignRegex.FindString(ss[0])
		if data.De == "" {
			return Data{}, errors.New("Not a regular spot; DE is wrong formatted! " + rawData)
		}

		//ss[1]=    7035.3  LA9QJA       CW 16 dB 16 WPM CQ             1553Z
		sss := splitSpotBySpaceRegex.Split(ss[1], 4)
		if len(sss) != 4 {
			return Data{}, errors.New("Not a regular spot; Unexpected split by space! " + rawData)
		}

		//Freq, band
		data.Freq = sss[1]
		if b, e := adif.GetBand(data.Freq); e == nil {
			data.BAND = b
		}

		data.Dx = sss[2]
		data.Comments = sss[3]

		//mode, if available
		ssss := splitSpotBySpaceRegex.Split(sss[3], 2)
		if m, e := adif.GetMode(ssss[0]); e == nil {
			data.MODE = m
		}

		//data.Raw = rawData
		data.Source = source
		return data, nil

	} else {
		return Data{}, ignore
	}
}

//func NewSpotPoc(rawData string, source string) (Data, error) {
//	//DX De S50ARX-#:   7035.3  LA9QJA       CW 16 dB 16 WPM CQ             1553Z
//	if matched := isSpotRegex.MatchString(rawData); matched {
//
//		data := Data{Raw: rawData} //Raw: DX De S50ARX-#:   7035.3  LA9QJA       CW 16 dB 16 WPM CQ             1553Z
//		s := rawData[6:]           //s:         S50ARX-#:   7035.3  LA9QJA       CW 16 dB 16 WPM CQ             1553Z
//		fmt.Println("s:", s)
//		ss := splitSpotByColonRegex.Split(s, 2)
//		if len(ss) != 2 {
//			return Data{}, errors.New("Not a regular spot; Unexpected split by colon! " + rawData)
//		}
//		fmt.Println("ss[0]=", ss[0])
//		fmt.Println("ss[1]=", ss[1])
//		//ss[0]= S50ARX-#
//		data.IsRbn = ss[0][len(ss[0])-1] == '#'
//		//spot sender
//		data.De = deCallSignRegex.FindString(ss[0])
//		if data.De == "" {
//			return Data{}, errors.New("Not a regular spot; DE is wrong formatted! " + rawData)
//		}
//
//		//ss[1]=    7035.3  LA9QJA       CW 16 dB 16 WPM CQ             1553Z
//		sss := splitSpotBySpaceRegex.Split(ss[1], 4)
//		if len(sss) != 4 {
//			return Data{}, errors.New("Not a regular spot; Unexpected split by space! " + rawData)
//		}
//		data.Freq = sss[1]
//		data.Dx = sss[2]
//		data.Comments = sss[3]
//
//		fmt.Println(len(sss))
//		for _, j := range sss {
//			fmt.Println(j)
//		}
//
//		data.Source = source
//		return data, nil
//
//	} else {
//		return Data{}, ignore
//	}
//}
