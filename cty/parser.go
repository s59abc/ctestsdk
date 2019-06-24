package cty

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func wrongFormattedRecord(description string, ctyDataRecord string) error {
	s := "wrong formatted cty.dat record: " + ctyDataRecord
	if len(description) > 0 {
		s = fmt.Sprintf("%s\n%s", description, s)
	}
	return errors.New(s)
}

///////////////
// regex
var pfxRegex *regexp.Regexp = regexp.MustCompile(`[0-9A-Z/]{1,}`)

// this function parse cty dat record, record format is defined at
// https://www.country-files.com/cty-dat-format/
func parseCtyDatRecord(ctyDatRecord string) (ctyDatList []Dta, err error) {
	if len(ctyDatRecord) < 76 {
		return nil, wrongFormattedRecord("", ctyDatRecord)
	}

	// initializing local variables
	primaryRecord := ctyDatRecord[:76]
	aliasRecords := strings.Split(ctyDatRecord[76:], ",")

	ctyDatList = make([]Dta, len(aliasRecords))
	primaryDta := Dta{}

	//
	// Example of ctyDatRecord
	// Slovenia:                 15:  28:  EU:   46.00:   -14.00:    -1.0:  S5:
	//    S5,=S51LGT/LH,=S52AL/YL,=S52L/LH,=S58U/LH,=S59HIJ/LH
	//
	/////////////////
	//first line
	//
	//Fields are aligned in columns and spaced out for readability only.
	//It is the “:” at the end of each field that acts as a delimiter for that field
	// Eight field delimiters are expected. Let's count them
	fields := strings.Split(primaryRecord, ":")
	if len(fields) != 9 {
		return nil, wrongFormattedRecord("Unexpected number of fields: "+strconv.Itoa(len(fields)), ctyDatRecord)
	}

	//COLUMN	LENGTH	DESCRIPTION
	//
	//1	26	Country Name
	primaryDta.countryName = strings.TrimSpace(fields[0])
	//
	//27	5	CQ Zone
	if cq, err := strconv.Atoi(strings.TrimSpace(fields[1])); err != nil {
		//TODO: test
		return nil, wrongFormattedRecord("Wrong formatted CQ Zone: "+fields[1], ctyDatRecord)
	} else {
		primaryDta.cqZone = cqzoneEnum(cq)
	}
	//
	//32	5	ITU Zone
	if itu, err := strconv.Atoi(strings.TrimSpace(fields[2])); err != nil {
		//TODO: test
		return nil, wrongFormattedRecord("Wrong formatted ITU Zone: "+fields[2], ctyDatRecord)
	} else {
		primaryDta.ituZone = ituzoneEnum(itu)
	}
	//
	//37	5	2-letter continent abbreviation
	if c, err := continent(strings.TrimSpace(fields[3])); err != nil {
		//TODO: test
		return nil, wrongFormattedRecord(err.Error(), ctyDatRecord)
	} else {
		primaryDta.continent = c
	}
	//
	//42	9	Latitude in degrees, + for North
	//51	10	Longitude in degrees, + for West
	lat, err := strconv.ParseFloat(strings.TrimSpace(fields[4]), 64)
	if err != nil {
		//TODO: test
		return nil, wrongFormattedRecord("Wrong formatted latitude: "+fields[4], ctyDatRecord)
	}
	lon, err := strconv.ParseFloat(strings.TrimSpace(fields[5]), 64)
	if err != nil {
		//TODO: test
		return nil, wrongFormattedRecord("Wrong formatted longitude: "+fields[5], ctyDatRecord)
	}
	primaryDta.latLon.Lat = lat
	primaryDta.latLon.Lon = lon
	//
	//61	9	Local time offset from GMT
	primaryDta.timeOffset = strings.TrimSpace(fields[6])
	//
	//70	6	Primary DXCC Prefix
	// (A “*” preceding this prefix indicates that the country is on the DARC WAEDC ctyDatList, and counts in CQ-sponsored contests, but not ARRL-sponsored contests).
	primaryPfx := strings.TrimSpace(fields[7])
	// remove preceding * if is present
	if len(primaryPfx) > 1 && primaryPfx[0] == byte('*') {
		primaryPfx = primaryPfx[1:]
	}
	primaryDta.prefix = primaryPfx
	//
	ctyDatList[0] = primaryDta
	//
	//
	// processing aliasRecords
	idx := 1
	for _, v := range aliasRecords {
		pfx := pfxRegex.FindString(v)
		if pfx != "" && pfx != primaryPfx { //Alias DXCC prefixes always include the primary one
			//TODO: remove
			fmt.Println("---------> PFX: ", pfx)
			aliasDta := primaryDta
			aliasDta.prefix = pfx
			ctyDatList[idx] = aliasDta
			idx++
		}

	}

	return ctyDatList, nil
}
