package keg

import (
	"github.com/stianeikeland/go-rpio"
	"log"
)

type Heater int

const (
	FIRST  Heater = iota
	SECOND Heater = iota
)

type HeaterState bool
const (
	ON HeaterState = true
	OFF HeaterState = false
)

type KegControl interface {
	ToggleHeater(h Heater)
	HeaterState(h Heater) HeaterState
	Temperature() float32
}

type keg struct {
	heaters []rpio.Pin
	temp    rpio.Pin
}

func (k keg) HeaterState(h Heater) HeaterState {
	return k.heaters[h].Read() == 0
}

func (k keg) ToggleHeater(h Heater) {
	k.heaters[h].Toggle()
}

func (k keg) Temperature() float32 {
	return float32(k.temp.Read())
}

func NewKeg() (KegControl, error) {

	err := rpio.Open()
	if err != nil {
		log.Println("could not open rpio. Details %+v", err)
		return nil, err
	}

	heaters := []rpio.Pin{rpio.Pin(10), rpio.Pin(10)}
	temp := rpio.Pin(10)

	heaters[0].Output()
	heaters[1].Output()
	temp.Input()

	k := keg{
		heaters: heaters,
		temp:    temp,
	}

	return KegControl(k), nil
}
