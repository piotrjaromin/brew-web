package temperature

import (
	"log"
	"time"

	"github.com/piotrjaromin/brew-web/brew/config"
)

type ReadTemperature interface {
	Temperature() (float64, error)
}

type Temperatures interface {
	Get() ([]TemperaturePoint, error)
	Add(s float64) error
}

type TemperaturePoint struct {
	Value     float64   `json:"value"`
	TimeStamp time.Time `json:"timestamp"`
}

func NewTemperatureStore(readTemp ReadTemperature, c config.Temperature) (Temperatures, error) {
	t, err := CreateTempDb()
	if err != nil {
		return nil, err
	}

	ticker := time.NewTicker(time.Duration(c.RefreshIntervalSeconds) * time.Second)
	go func() {
		for {
			<-ticker.C

			temp, err := readTemp.Temperature()
			if err != nil {
				log.Print("[Temp] Could not read temperature", err)
				return
			}
			t.Add(temp)
		}
	}()

	return t, nil
}
