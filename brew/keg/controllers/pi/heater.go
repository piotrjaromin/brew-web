package pi

import (
	"log"

	"github.com/piotrjaromin/brew-web/brew/config"

	"github.com/brian-armstrong/gpio"
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
	pin pin
}

type pin interface {
	Read() (value uint, err error)
	High() error
	Low() error
}

func (h rpioHeater) SetState(state HeaterState) {
	if state == HEATER_ENABLED {
		h.pin.High()
		return
	}

	h.pin.Low()
}

func (h rpioHeater) State() HeaterState {
	state, err := h.pin.Read()
	if err != nil {
		panic(err) // terrible idea!
	}

	if state == 1 {
		return HEATER_ENABLED
	}

	return HEATER_DISABLED
}

func GetHeaters(c config.Keg) ([]Heater, error) {
	heaters := []Heater{}
	for pinConfig := range c.Heaters {
		log.Printf("Creating heater for %+v\n", pinConfig)
		pin := gpio.NewOutput(uint(pinConfig), false)

		heater := rpioHeater{
			pin: pin,
		}
		heaters = append(heaters, heater)
	}

	return heaters, nil
}
