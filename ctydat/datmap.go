package ctydat

import (
	"ctestsdk/spot"
	"sync"
)

var ctyDat ctyDatMap = ctyDatMap{internal: make(map[string]spot.CtyDta)}

type ctyDatMap struct {
	sync.RWMutex
	internal map[string]spot.CtyDta //
}

func load(ctyDtaFileName string) {
	// todo loadFromFile in string is not empty and file exist

}

func (a *ctyDatMap) get(key string) (spot.CtyDta, bool) {
	a.RLock()
	dta, has := a.internal[key]
	a.RUnlock()
	return dta, has
}

func AddCtyDat(spot spot.Data) spot.Data {

	return spot
}
