package keg

import (
	"log"
	"math/rand"
)

type KegMock struct {
	state HeaterState
	temp  float64
}

func (k KegMock) HeaterState(h Heater) HeaterState {

	log.Println("[mock]heater state", h)
	return k.state
}

func (k *KegMock) ToggleHeater(h Heater) {

	if h == FIRST {
		k.state = !k.state
		log.Println("[mock]toggle heater", h, " state is ", k.state)
	}
}

func (k *KegMock) Temperature() (float64, error) {
	log.Println("[mock]Temperature")

	if k.state == ON {
		k.temp++
	} else {
		k.temp--
	}
	return k.temp, nil
}

func NewKegMock() (KegControl, error) {

	return &KegMock{HeaterState(false), rand.Float64() * 100}, nil
}
