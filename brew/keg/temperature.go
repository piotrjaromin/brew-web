package keg

import (
	"log"
	"time"
)

type Temperatures interface {
	Get() []TemperaturePoint
}

type TemperatureCache struct {
	maxCacheSize int
	cache        []TemperaturePoint
}

type TemperaturePoint struct {
	Value     float64   `json:"value"`
	TimeStamp time.Time `json:"timestamp"`
}

func (t TemperatureCache) Get() []TemperaturePoint {

	return t.cache
}

func (t *TemperatureCache) add(s float64) {
	t.cache = append(t.cache, TemperaturePoint{s, time.Now()})
	if len(t.cache) > t.maxCacheSize {
		t.cache = t.cache[1:len(t.cache)]
	}
}

func NewTemperatureCache(keg KegControl, intervalSec time.Duration, cacheSize int) Temperatures {

	t := &TemperatureCache{
		cacheSize,
		make([]TemperaturePoint, 0, cacheSize),
	}

	ticker := time.NewTicker(intervalSec * time.Second)
	go func() {
		for {
			<-ticker.C

			temp, err := keg.Temperature()
			if err != nil {
				log.Print("[T-Cache]Could not read temperature", err)
				return
			}
			t.add(temp)
		}
	}()

	return t
}
