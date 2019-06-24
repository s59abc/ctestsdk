package cty

import "errors"

type continentEnum int

const (
	NA continentEnum = iota
	SA
	EU
	AF
	OC
	AS
	AN
)

func continent(continentAbbreviation string) (continentEnum, error) {
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

func (a continentEnum) String() string {
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
