package pi

import (
	"log"

	"github.com/piotrjaromin/brew-web/brew/config"
	"github.com/piotrjaromin/brew-web/brew/keg"

	rpio "github.com/stianeikeland/go-rpio"
)

type kegStruct struct {
	heaters []rpio.Pin
	temp    W1Device
}

func (k kegStruct) HeaterState(h keg.Heater) keg.HeaterState {
	state := k.heaters[h].Read()
	log.Printf("State for heater %+v is %+v", h, state)
	return state != 0
}

func (k kegStruct) ToggleHeater(h keg.Heater) {
	log.Printf("toggle: %+v", h)
	k.heaters[h].Toggle()
}

func (k kegStruct) Temperature() (float64, error) {
	t, err := k.temp.Value(1, "t")
	return float64(t) / 1000, err
}

func (k kegStruct) SetHeaterState(h keg.Heater, enabled keg.HeaterState) {
	if enabled {
		k.heaters[h].High()
	} else {
		k.heaters[h].Low()
	}
}

func NewKeg(tempDev W1Device, c config.Keg) (keg.KegControl, error) {
	err := rpio.Open()
	if err != nil {
		log.Printf("could not open rpio. Details %+v\n", err)
		return nil, err
	}

	heaters := []rpio.Pin{}
	for pinConfig := range c.Heaters {
		pin := rpio.Pin(pinConfig)
		pin.Output()
		heaters = append(heaters, pin)
	}

	temp := tempDev

	k := kegStruct{
		heaters: heaters,
		temp:    temp,
	}

	return keg.KegControl(k), nil
}
