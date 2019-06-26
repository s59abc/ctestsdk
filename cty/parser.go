package cty

import (
	"bytes"
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
var regexPfx *regexp.Regexp = regexp.MustCompile(`[0-9A-Za-z/]{1,}`)
var regexOverrideCqZone *regexp.Regexp = regexp.MustCompile(`\([0-9]{1,2}\)`)
var regexOverrideItuZone *regexp.Regexp = regexp.MustCompile(`[([0-9]{1,2}]`)
var regexZone *regexp.Regexp = regexp.MustCompile(`[0-9]{1,2}`)
var regexOverrideLatLon *regexp.Regexp = regexp.MustCompile(`<[0-9-]+\.[0-9]+/[0-9-]+\.[0-9]+>`)
var regexOverrideTimeOffset *regexp.Regexp = regexp.MustCompile(`~[0-9-.]+~`)
var regexIsComment *regexp.Regexp = regexp.MustCompile(`^#`)

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
		return []Dta{}, wrongFormattedRecord("Unexpected number of fields: "+strconv.Itoa(len(fields)), ctyDatRecord)
	}

	//COLUMN	LENGTH	DESCRIPTION
	//
	//1	26	Country Name
	primaryDta.countryName = strings.TrimSpace(fields[0])
	//
	//27	5	CQ Zone
	if cq, err := strconv.Atoi(strings.TrimSpace(fields[1])); err != nil {
		//TODO: test
		return []Dta{}, wrongFormattedRecord("Wrong formatted CQ Zone: "+fields[1], ctyDatRecord)
	} else {
		primaryDta.cqZone = cqzoneEnum(cq)
	}
	//
	//32	5	ITU Zone
	if itu, err := strconv.Atoi(strings.TrimSpace(fields[2])); err != nil {
		//TODO: test
		return []Dta{}, wrongFormattedRecord("Wrong formatted ITU Zone: "+fields[2], ctyDatRecord)
	} else {
		primaryDta.ituZone = ituzoneEnum(itu)
	}
	//
	//37	5	2-letter continent abbreviation
	if c, err := continent(strings.TrimSpace(fields[3])); err != nil {
		//TODO: test
		return []Dta{}, wrongFormattedRecord(err.Error(), ctyDatRecord)
	} else {
		primaryDta.continent = c
	}
	//
	//42	9	Latitude in degrees, + for North
	//51	10	Longitude in degrees, + for West
	lat, err := strconv.ParseFloat(strings.TrimSpace(fields[4]), 64)
	if err != nil {
		//TODO: test
		return []Dta{}, wrongFormattedRecord("Wrong formatted latitude: "+fields[4], ctyDatRecord)
	}
	lon, err := strconv.ParseFloat(strings.TrimSpace(fields[5]), 64)
	if err != nil {
		//TODO: test
		return []Dta{}, wrongFormattedRecord("Wrong formatted longitude: "+fields[5], ctyDatRecord)
	}
	primaryDta.latLon.Lat = lat
	primaryDta.latLon.Lon = lon
	//
	//61	9	Local time offset from GMT
	primaryDta.timeOffset = strings.TrimSpace(fields[6])
	//
	//70	6	Primary DXCC Prefix
	// (A “*” preceding this aliasPrefix indicates that the country is on the DARC WAEDC ctyDatList, and counts in CQ-sponsored contests, but not ARRL-sponsored contests).
	primaryPfx := strings.TrimSpace(fields[7])
	// remove preceding * if is present
	if len(primaryPfx) > 1 && primaryPfx[0] == byte('*') {
		primaryPfx = primaryPfx[1:]
	}
	primaryDta.primaryPrefix = primaryPfx
	primaryDta.aliasPrefix = primaryPfx
	//
	ctyDatList[0] = primaryDta
	//
	//
	// processing aliasRecords
	idx := 1
	for _, v := range aliasRecords {
		pfx := regexPfx.FindString(v)
		if pfx != "" && pfx != primaryPfx { //Alias DXCC prefixes always include the primary one
			aliasDta := primaryDta
			aliasDta.aliasPrefix = pfx
			//
			// check for override cq zone
			overrideCqZone := regexOverrideCqZone.FindString(v)
			if overrideCqZone != "" {
				cqZone := regexZone.FindString(overrideCqZone)
				if cqZone != "" {
					if i, e := strconv.Atoi(cqZone); e != nil {
						//TODO: test
						return []Dta{}, wrongFormattedRecord("wrong formatted override CQ Zone: "+cqZone+" "+e.Error(), ctyDatRecord)
					} else {
						aliasDta.cqZone = cqzoneEnum(i)
					}
				} else {
					//TODO: test
					return []Dta{}, wrongFormattedRecord("wrong formatted override CQ Zone: "+cqZone, ctyDatRecord)
				}
			}
			//
			// check for override itu zone
			overrideItuZone := regexOverrideItuZone.FindString(v)
			if overrideItuZone != "" {
				ituZone := regexZone.FindString(overrideItuZone)
				if ituZone != "" {
					if i, e := strconv.Atoi(ituZone); e != nil {
						//TODO: test
						return []Dta{}, wrongFormattedRecord("wrong formatted override ITU Zone: "+ituZone+" "+e.Error(), ctyDatRecord)
					} else {
						aliasDta.ituZone = ituzoneEnum(i)
					}
				} else {
					//TODO: test
					return []Dta{}, wrongFormattedRecord("wrong formatted override ITU Zone: "+ituZone, ctyDatRecord)
				}
			}
			//
			// check for override lat lon
			overrideLatLon := regexOverrideLatLon.FindString(v)
			if overrideLatLon != "" {
				latLonSS := strings.Split(overrideLatLon, "/")
				latS := latLonSS[0][1:]
				lonS := latLonSS[1][:len(latLonSS[1])-1]
				lat, errLat := strconv.ParseFloat(latS, 64)
				lon, errLon := strconv.ParseFloat(lonS, 64)
				if errLat != nil || errLon != nil {
					return []Dta{}, wrongFormattedRecord("wrong formatted override Latitude/Longitude: "+overrideLatLon, ctyDatRecord)
				} else {
					aliasDta.latLon.Lat = lat
					aliasDta.latLon.Lon = lon
				}
			}
			//
			// Override local time offset from GMT
			overrideTimeOffset := regexOverrideTimeOffset.FindString(v)
			if overrideTimeOffset != "" {
				aliasDta.timeOffset = strings.Trim(overrideTimeOffset, "~")
			}

			if idx >= cap(ctyDatList) {
				ctyDatList = append(ctyDatList, aliasDta)
			} else {
				ctyDatList[idx] = aliasDta
			}
			idx++
		}

	}

	return ctyDatList, nil
}

// this function removes comments
func removeComments(ctyDatRecords string) string {
	//first we have to check if ctyDatRecords has a comments
	if strings.Contains(ctyDatRecords, "#") {
		// Normalize \r\n (windows) and \r (mac) into \n (unix)
		if strings.Contains(ctyDatRecords, "\r") {
			b := bytes.Replace([]byte(ctyDatRecords), []byte{13, 10}, []byte{10}, -1)
			b = bytes.Replace(b, []byte{13}, []byte{10}, -1)
			ctyDatRecords = string(b)
		}
		// here we have records normalized to unix format
		lines := strings.Split(ctyDatRecords, "\n")
		slice := make([]string, 0, len(lines)) //container for slice without comments, just a payload
		for _, v := range lines {
			if !regexIsComment.MatchString(v) {
				slice = append(slice, v)
			}
		}
		//Joining a slice of strings is faster that using +
		ctyDatRecords = strings.Join(slice, "\n")
	}

	return ctyDatRecords
}

func parseCtyDatRecordsOld(ctyDatRecords string) (err error) {
	ctyDatRecords = removeComments(ctyDatRecords)
	c := 0
	records := strings.Split(ctyDatRecords, ";")
	for _, v := range records {
		if len(v) > 10 {
			v, e := parseCtyDatRecord(v)
			if e != nil {
				return e
			} else {
				c += len(v)
			}
		}
	}
	fmt.Println(c)

	return nil
}

func parseCtyDatRecords(ctyDatRecords string) (size int, err error) {
	ctyDatRecords = removeComments(ctyDatRecords)
	size = 0
	records := strings.Split(ctyDatRecords, ";")
	for _, v := range records {
		if len(v) > 10 {
			v, e := parseCtyDatRecord(v)
			if e != nil {
				return size, e
			} else {
				size += len(v)
			}
		}
	}
	fmt.Println(size)

	return size, nil
}
