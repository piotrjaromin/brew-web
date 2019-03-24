package keg

import (
	"log"
	"math/rand"
)

type KegMock struct {
	heater1 HeaterState
	heater2 HeaterState
	temp    float64
}

func (k KegMock) HeaterState(h Heater) HeaterState {

	if h == FIRST {
		log.Println("[mock] heater state", h, "is ", k.heater1)
		return k.heater1
	}

	log.Println("[mock] heater state", h, "is ", k.heater2)
	return k.heater2
}

func (k *KegMock) ToggleHeater(h Heater) {

	if h == FIRST {
		k.heater1 = !k.heater1
		log.Println("[mock] toggle heater", h, ", new state is ", k.heater1)
	}

	if h == SECOND {
		k.heater2 = !k.heater2
		log.Println("[mock] toggle heater", h, ", new state is ", k.heater2)
	}
}

func (k *KegMock) SetHeaterState(h Heater, state HeaterState) {

	if h == FIRST {
		k.heater1 = state
		log.Println("[mock] SetHeaterState heater", h, ", new state is ", state)
	}

	if h == SECOND {
		k.heater2 = state
		log.Println("[mock] SetHeaterState heater", h, ", new state is ", state)
	}
}

func (k *KegMock) Temperature() (float64, error) {
	log.Println("[mock] Temperature")

	if k.heater2 == ON {
		k.temp += 0.6
	}

	if k.heater1 == ON {
		k.temp += 0.6
	}

	if k.heater1 == OFF && k.heater1 == OFF {
		k.temp--
	}

	return k.temp, nil
}

func NewKegMock() (KegControl, error) {

	return &KegMock{HeaterState(false), HeaterState(false), rand.Float64() * 100}, nil
}
