package ctydat

import "sync"

var muxSkimmerMap sync.RWMutex
var skimmerMap map[string]string //callsing / locator

func init() {
	skimmerMap = map[string]string{
		"S50ARX": "JN65TW",
		"KM3T":   "FN42ET",
	}
}

func getSkimmerLocator(de string) (locator string, find bool) {
	muxSkimmerMap.RLock()
	locator, find = skimmerMap[de]
	muxSkimmerMap.RUnlock()
	return locator, find
}
