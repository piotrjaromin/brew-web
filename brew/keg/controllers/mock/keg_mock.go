package mock

import (
	"log"
)

type KegMock struct {
	power float64
	temp  float64
}

func (k KegMock) GetHeaterPower() float64 {
	return k.power
}

func (k *KegMock) SetHeaterPower(power float64) {
	k.power = power
}

func (k *KegMock) Temperature() (float64, error) {
	log.Println("[mock] Temperature")
	if k.power <= 0 {
		k.temp -= 0.5
	} else {
		k.temp += k.power
	}

	return k.temp, nil
}

func NewKegMock() (*KegMock, error) {
	return &KegMock{
		power: 0,
		temp:  0,
	}, nil
}
