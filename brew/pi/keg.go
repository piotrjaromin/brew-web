package pi

import (
	"log"

	"github.com/piotrjaromin/brew-web/brew/keg"
	rpio "github.com/stianeikeland/go-rpio"
)

const HEATER1_PIN = 14
const HEATER2_PIN = 15

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
	log.Printf("toggle heater: %+v", h)
	k.heaters[h].Toggle()
}

func (k kegStruct) Temperature() (float64, error) {
	t, err := k.temp.Value(1, "t")
	return float64(t) / 1000, err
}

func NewKeg(tempDev W1Device) (keg.KegControl, error) {

	err := rpio.Open()
	if err != nil {
		log.Println("could not open rpio. Details %+v", err)
		return nil, err
	}

	heaters := []rpio.Pin{rpio.Pin(HEATER1_PIN), rpio.Pin(HEATER2_PIN)}
	temp := tempDev

	heaters[0].Output()
	heaters[1].Output()

	k := kegStruct{
		heaters: heaters,
		temp:    temp,
	}

	return keg.KegControl(k), nil
}
