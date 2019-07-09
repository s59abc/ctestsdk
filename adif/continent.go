package adif

import "errors"

type ContinentEnum int

const (
	UN ContinentEnum = iota //UN Unknown
	NA
	SA
	EU
	AF
	OC
	AS
	AN
)

func Continent(continentAbbreviation string) (ContinentEnum, error) {
	switch continentAbbreviation {
	case "EU":
		return EU, nil
	case "NA":
		return NA, nil
	case "SA":
		return SA, nil
	case "AF":
		return AF, nil
	case "OC":
		return OC, nil
	case "AS":
		return AS, nil
	case "AN":
		return AN, nil
	default:
		return 0, errors.New("Wrong Continent Abbreviation: " + continentAbbreviation)
	}

}

func (a ContinentEnum) String() string {
	switch a {
	case EU:
		return "EU"
	case NA:
		return "NA"
	case SA:
		return "SA"
	case AF:
		return "AF"
	case OC:
		return "OC"
	case AS:
		return "AS"
	case AN:
		return "AN"
	default:
		return "??"
	}

}
