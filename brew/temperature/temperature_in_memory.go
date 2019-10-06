package temperature

import (
	"time"
)

type TemperatureStore struct {
	maxCacheSize int
	cache        []TemperaturePoint
}

func (t TemperatureStore) Get() ([]TemperaturePoint, error) {
	return t.cache, nil
}

func (t *TemperatureStore) Add(s float64) error {
	t.cache = append(t.cache, TemperaturePoint{s, time.Now()})
	if len(t.cache) > t.maxCacheSize {
		t.cache = t.cache[1:len(t.cache)]
	}

	return nil
}

func CreateTempStoreInMemory() Temperatures {
	cacheSize := 100

	return &TemperatureStore{
		cacheSize,
		make([]TemperaturePoint, 0, cacheSize),
	}
}
