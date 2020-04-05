package pi

import (
	"log"

	"github.com/piotrjaromin/brew-web/brew/config"
	rpio "github.com/stianeikeland/go-rpio/v4"
)

type HeaterState string

const (
	HEATER_DISABLED HeaterState = "disabled"
	HEATER_ENABLED  HeaterState = "enabled"
)

type Heater interface {
	SetState(state HeaterState)
	State() HeaterState
}

type rpioHeater struct {
	pin rpio.Pin
}

func (h rpioHeater) SetState(state HeaterState) {
	if state == HEATER_ENABLED {
		h.pin.High()
		return
	}

	h.pin.Low()
}

func (h rpioHeater) State() HeaterState {
	if h.pin.Read() == rpio.High {
		return HEATER_ENABLED
	}

	return HEATER_DISABLED
}

func GetHeaters(c config.Keg) ([]Heater, error) {
	err := rpio.Open()
	if err != nil {
		log.Printf("could not open rpio. Details %+v\n", err)
		return []Heater{}, err
	}

	heaters := []Heater{}
	for pinConfig := range c.Heaters {
		log.Printf("Creating heater for %+v\n", pinConfig)
		pin := rpio.Pin(pinConfig)
		pin.Output()

		heater := rpioHeater{
			pin: pin,
		}
		heaters = append(heaters, heater)
	}

	return heaters, nil
}
