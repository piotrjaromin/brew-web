package temperature

import (
	"log"
	"time"

	"github.com/piotrjaromin/brew-web/brew/keg"
)

type Temperatures interface {
	Get() ([]TemperaturePoint, error)
	Add(s float64) error
}

type TemperaturePoint struct {
	Value     float64   `json:"value"`
	TimeStamp time.Time `json:"timestamp"`
}

func NewTemperatureStore(keg keg.KegControl, intervalSec time.Duration, cacheSize int) (Temperatures, error) {
	t, err := CreateTempDb()
	if err != nil {
		return nil, err
	}

	ticker := time.NewTicker(intervalSec * time.Second)
	go func() {
		for {
			<-ticker.C

			temp, err := keg.Temperature()
			if err != nil {
				log.Print("[Temp] Could not read temperature", err)
				return
			}
			t.Add(temp)
		}
	}()

	return t, nil
}
