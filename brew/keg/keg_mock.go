package keg

import (
	"log"
	"math/rand"
)

type KegMock struct {
	state HeaterState
}

func (k KegMock) HeaterState(h Heater) HeaterState {

	log.Println("[mock]heater state", h)
	return k.state
}

func (k*KegMock) ToggleHeater(h Heater) {

	k.state = !k.state

	log.Println("[mock]toggle heater", h, " state is ", k.state)
}

func (k KegMock) Temperature() float32 {
	log.Println("[mock]Temperature")
	return rand.Float32() * 100
}

func NewKegMock() (KegControl, error) {

	return KegControl(&KegMock{}), nil
}
