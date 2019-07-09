package ctydat

import (
	"ctestsdk/ctydat/internal"
	"ctestsdk/spot"
	"fmt"
	"strings"
	"sync"
)

var ctyDat ctyDatMap = ctyDatMap{internal: make(map[string]spot.CtyDta)}

type ctyDatMap struct {
	sync.RWMutex
	internal map[string]spot.CtyDta //
}

func init() {
	load("")
}

func load(ctyDtaFileName string) {
	// todo loadFromFile in string is not empty and file exist
	if m, err := parseCtyDatRecords(internal.CtyWtModDatFakeFile); err == nil {
		ctyDat.Lock()
		ctyDat.internal = m
		ctyDat.Unlock()
	}

}

func get(key string) (spot.CtyDta, bool) {
	dta := spot.CtyDta{}
	find := false
	if strings.Contains(key, "/") {
		ss := strings.Split(key, "/")
		if len(ss) == 2 && len(ss[0]) > len(ss[1]) {
			//AC2AI/KH2 --> KH2/AC2AI
			key = ss[1] + "/" + ss[0]
		}
	}
	ctyDat.RLock()
	for !find && len(key) > 0 {
		fmt.Println(key)
		dta, find = ctyDat.internal[key]
		key = key[:len(key)-1]
	}
	ctyDat.RUnlock()
	fmt.Println(dta.String())
	return dta, find
}

func AddCtyDat(spot spot.Data) spot.Data {
	deDta, find := get(spot.DE())
	if find {
		spot.DeCtyDta = deDta
	}
	dxDta, find := get(spot.DX())
	if find {
		spot.DeCtyDta = dxDta
	}
	return spot
}
