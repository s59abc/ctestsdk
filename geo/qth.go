package geo

import (
	"errors"
	"fmt"
	"github.com/golang/geo/s2"
	"math"
	"regexp"
	"strings"
)

type QTH struct {
	loc    string    // Maidenhead QTH Locator
	latLon LatLonDeg // latLon represent a point as a pair of latitude and longitude degrees
	latLng s2.LatLng // LatLng represents a point on the unit sphere as a pair of angles.
}

const earthRadiusKm = 6371.0088 // mean Earth radius in km

// returns distance between two Maindenhead locators in km
func DistanceLocator(locatorA string, locatorB string) (float64, error) {
	a, err := NewQthFromLOC(locatorA)
	b, err := NewQthFromLOC(locatorB)
	if err != nil {
		return 0, err
	} else {
		d := a.latLng.Distance(b.latLng)
		return d.Radians() * earthRadiusKm, nil
	}
}

func DistanceQTH(a, b QTH) float64 {
	d := a.latLng.Distance(b.latLng)
	return d.Radians() * earthRadiusKm
}

func illegalArgumentError(arg string) error {
	return errors.New(fmt.Sprintf("Illegal argumet value! qthLocator=%s", arg))
}

func NewQthFromLOC(qthLocator string) (QTH, error) {
	qthLocator = strings.ToUpper(qthLocator)
	qth := QTH{}
	switch len(qthLocator) {
	case 6:
		{
			if matched, _ := regexp.MatchString(`^[A-R]{2}[0-9]{2}[A-X]{2}$`, qthLocator); matched {
				f := fieldDecode(latLonChar{
					latChar: qthLocator[1],
					lonChar: qthLocator[0],
				})
				s := squareDecode(latLonChar{
					latChar: qthLocator[3],
					lonChar: qthLocator[2],
				})
				ss := subsquareDecode(latLonChar{
					latChar: qthLocator[5],
					lonChar: qthLocator[4],
				})
				lat := f.decoded.Lat + s.decoded.Lat + ss.decoded.Lat/60 + 0.02083333 // 1.25' / 60
				lon := f.decoded.Lon + s.decoded.Lon + ss.decoded.Lon/60 + 0.04166667 // 2.5' / 60
				return QTH{
					loc:    qthLocator,
					latLon: LatLonDeg{lat, lon},
					latLng: s2.LatLngFromDegrees(lat, lon),
				}, nil
			} else {
				return qth, illegalArgumentError(qthLocator)
			}
		}
	case 4:
		{
			if matched, _ := regexp.MatchString(`^[A-R]{2}[0-9]{2}$`, qthLocator); matched {
				f := fieldDecode(latLonChar{
					latChar: qthLocator[1],
					lonChar: qthLocator[0],
				})
				s := squareDecode(latLonChar{
					latChar: qthLocator[3],
					lonChar: qthLocator[2],
				})
				lat := f.decoded.Lat + s.decoded.Lat + 0.5
				lon := f.decoded.Lon + s.decoded.Lon + 1
				return QTH{
					loc:    qthLocator,
					latLon: LatLonDeg{lat, lon},
					latLng: s2.LatLngFromDegrees(lat, lon),
				}, nil
			} else {
				return qth, illegalArgumentError(qthLocator)
			}
		}
	case 2:
		{
			if matched, _ := regexp.MatchString(`^[A-R]{2}`, qthLocator); matched {
				f := fieldDecode(latLonChar{
					latChar: qthLocator[1],
					lonChar: qthLocator[0],
				})
				lat := f.decoded.Lat + 5
				lon := f.decoded.Lon + 10

				return QTH{
					loc:    qthLocator,
					latLon: LatLonDeg{lat, lon},
					latLng: s2.LatLngFromDegrees(lat, lon),
				}, nil

			} else {
				return qth, illegalArgumentError(qthLocator)
			}
		}

	default:
		return qth, illegalArgumentError(qthLocator)
	}
}

func NewQthFromLatLon(latitude, longitude float64) (QTH, error) {
	lld := LatLonDeg{
		Lat: latitude,
		Lon: longitude,
	}
	if math.Abs(latitude) > 90 || math.Abs(longitude) > 180 {
		return QTH{}, illegalArgumentError(lld.String())
	}
	f, s, ss := subsquareEncode(lld)
	return QTH{
		loc:    f.encoded.getLonChar() + f.encoded.getLatChar() + s.encoded.getLonChar() + s.encoded.getLatChar() + ss.encoded.getLonChar() + ss.encoded.getLatChar(),
		latLon: lld,
		latLng: s2.LatLngFromDegrees(latitude, longitude),
	}, nil
}
