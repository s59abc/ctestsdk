// Package cty provides Amateur Radio Country Files functionality.
// CTY.DAT Format is used (https://www.country-files.com/cty-dat-format/)
package cty

import (
	"ctestsdk/geo"
	"fmt"
)

type Dta struct {
	countryName string        //Country Name
	prefix      string        //Primary or Alias DXCC Prefix without optional * indicator
	continent   continentEnum //2-letter continent abbreviation
	cqZone      cqzoneEnum    //CQ Zone
	ituZone     ituzoneEnum   //ITU Zone
	latLon      geo.LatLonDeg //Latitude in degrees, + for North; Longitude in degrees, + for West
	timeOffset  string        //Local time offset from GMT
}

func (a Dta) String() string {
	return fmt.Sprintf("countryName=%s, prefix=%s, continent=%s, cqZone=%d, ituZone=%d, %s, timeOffset=%s", a.countryName, a.prefix, a.continent.String(), a.cqZone, a.ituZone, a.latLon.String(), a.timeOffset)

}
