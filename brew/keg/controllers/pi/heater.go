package pi

import (
	"log"
	"time"

	"github.com/piotrjaromin/brew-web/brew/config"

	"github.com/piotrjaromin/gpio"
)

type HeaterState string

const (
	HEATER_DISABLED HeaterState = "disabled"
	HEATER_ENABLED  HeaterState = "enabled"
)

type Heater interface {
	SetState(state HeaterState) error
	State() HeaterState
}

type rpioHeater struct {
	pin   pin
	state HeaterState
}

type pin interface {
	Read() (value uint, err error)
	High() error
	Low() error
}

func (h *rpioHeater) SetState(state HeaterState) error {
	h.state = state
	if state == HEATER_ENABLED {
		return h.pin.High()
	}

	return h.pin.Low()
}

func (h rpioHeater) State() HeaterState {
	return h.state
}

func GetHeaters(c config.Keg) ([]Heater, error) {
	heaters := []Heater{}
	for pinConfig := range c.Heaters {
		log.Printf("Creating heater for %+v\n", pinConfig)
		pin, err := gpio.NewOutput(uint(pinConfig), false)
		if err != nil {
			log.Printf("Warning. %s", err.Error())
		}

		time.Sleep(5 * time.Second)
		pin.High()

		heater := &rpioHeater{
			pin:   &pin,
			state: HEATER_DISABLED,
		}
		heaters = append(heaters, heater)
	}

	return heaters, nil
}
