package ctydat

import (
	"ctestsdk/ctydat/internal"
	"ctestsdk/spot"
	"strings"
	"sync"
	"time"
)

var muxCtyDatMap sync.RWMutex
var ctyDatMap map[string]spot.CtyDta

//////////////// cache, we need timeStamp to periodically clear old items
//TODO: metla
type cachedItem struct {
	timeStamp int64
	item      spot.CtyDta
}

var muxCtyDatMapCache sync.RWMutex
var ctyDatMapCache map[string]cachedItem // key is call sign

func init() {
	ctyDatMapCache = map[string]cachedItem{}
	load("")
}

func load(ctyDtaFileName string) {
	// todo loadFromFile in string is not empty and file exist
	if m, err := parseCtyDatRecords(internal.CtyWtModDatFakeFile); err == nil {
		muxCtyDatMap.Lock()
		ctyDatMap = m
		muxCtyDatMap.Unlock()
	}

}

func cacheTryFrom(callSign string) (spot.CtyDta, bool) {
	muxCtyDatMapCache.RLock()
	defer muxCtyDatMapCache.RUnlock()
	if v, has := ctyDatMapCache[callSign]; has {
		return v.item, has
	} else {
		return spot.CtyDta{}, false
	}
}

func cachePutInto(callSign string, ctyDta spot.CtyDta) {
	muxCtyDatMapCache.Lock()
	ctyDatMapCache[callSign] = cachedItem{time.Now().Unix(), ctyDta}
	muxCtyDatMapCache.Unlock()
}

func cacheLen() int {
	muxCtyDatMapCache.RLock()
	defer muxCtyDatMapCache.RUnlock()
	return len(ctyDatMapCache)

}

func get(callSign string) (spot.CtyDta, bool) {
	callSign = strings.TrimSpace(callSign)
	callSign = strings.ToUpper(callSign)

	if dta, has := cacheTryFrom(callSign); has {
		return dta, has
	}

	tempCallSign := callSign
	dta := spot.CtyDta{}
	find := false

	// handling such cases like AC2AI/KH2 --> KH2/AC2AI
	if strings.Contains(tempCallSign, "/") {
		ss := strings.Split(tempCallSign, "/")
		if len(ss) == 2 && len(ss[0]) > len(ss[1]) {
			tempCallSign = ss[1] + "/" + ss[0]
		}
	}
	muxCtyDatMap.RLock()
	for !find && len(tempCallSign) > 0 {
		dta, find = ctyDatMap[tempCallSign]
		tempCallSign = tempCallSign[:len(tempCallSign)-1]
	}
	muxCtyDatMap.RUnlock()

	if find {
		cachePutInto(callSign, dta)
	}
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
