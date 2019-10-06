package mock

import (
	"log"
	"math/rand"

	"github.com/piotrjaromin/brew-web/brew/keg"
)

type KegMock struct {
	heater1 keg.HeaterState
	heater2 keg.HeaterState
	temp    float64
}

func (k KegMock) HeaterState(h keg.Heater) keg.HeaterState {

	if h == keg.FIRST {
		log.Println("[mock] heater state", h, "is ", k.heater1)
		return k.heater1
	}

	log.Println("[mock] heater state", h, "is ", k.heater2)
	return k.heater2
}

func (k *KegMock) ToggleHeater(h keg.Heater) {

	if h == keg.FIRST {
		k.heater1 = !k.heater1
		log.Println("[mock] toggle heater", h, ", new state is ", k.heater1)
	}

	if h == keg.SECOND {
		k.heater2 = !k.heater2
		log.Println("[mock] toggle heater", h, ", new state is ", k.heater2)
	}
}

func (k *KegMock) SetHeaterState(h keg.Heater, state keg.HeaterState) {

	if h == keg.FIRST {
		k.heater1 = state
		log.Println("[mock] SetHeaterState heater", h, ", new state is ", state)
	}

	if h == keg.SECOND {
		k.heater2 = state
		log.Println("[mock] SetHeaterState heater", h, ", new state is ", state)
	}
}

func (k *KegMock) Temperature() (float64, error) {
	log.Println("[mock] Temperature")

	if k.heater2 == keg.ON {
		k.temp += 0.6
	}

	if k.heater1 == keg.ON {
		k.temp += 0.6
	}

	if k.heater1 == keg.OFF && k.heater2 == keg.OFF {
		k.temp--
	}

	return k.temp, nil
}

func NewKegMock() (keg.KegControl, error) {

	return &KegMock{keg.HeaterState(false), keg.HeaterState(false), rand.Float64() * 100}, nil
}
