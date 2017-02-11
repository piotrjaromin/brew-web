package keg

import (
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
	Value     float32 `json:"value"`
	TimeStamp time.Time `json:"timestamp"`
}

func (t TemperatureCache) Get() []TemperaturePoint {

	return t.cache
}

func (t *TemperatureCache) add(s float32) {
	t.cache = append(t.cache, TemperaturePoint{s, time.Now()})
	if len(t.cache) > t.maxCacheSize {
		t.cache = t.cache[1: len(t.cache)]
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
			t.add(float32(keg.Temperature()))
		}
	}()

	return t
}
