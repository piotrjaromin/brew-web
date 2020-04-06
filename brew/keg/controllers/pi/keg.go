package pi

import (
	"log"
	"math"
)

type kegStruct struct {
	heaters []Heater
	temp    W1Device
}

func (k *kegStruct) GetHeaterPower() float64 {
	sum := 0.0

	for index, heater := range k.heaters {
		log.Printf("Heater state %d is %s \n", index, heater.State())
		if heater.State() == HEATER_ENABLED {
			sum++
		}
	}

	lenHeaters := len(k.heaters)

	log.Printf("sum is %f is %d \n", sum, lenHeaters)
	return sum / float64(lenHeaters)
}

func (k *kegStruct) Temperature() (float64, error) {
	t, err := k.temp.Value(1, "t")
	return float64(t) / 1000, err
}

func (k *kegStruct) SetHeaterPower(power float64) {
	lenHeaters := len(k.heaters)
	val := int(math.Round(power * float64(lenHeaters)))
	log.Printf("[SetHeaterPower] val: %d", val)

	for heaterIndex := 0; heaterIndex < lenHeaters; heaterIndex++ {
		state := HEATER_DISABLED
		if val <= heaterIndex {
			state = HEATER_DISABLED
		} else {
			state = HEATER_ENABLED
		}

		log.Printf("Setting heater %d to %s\n", heaterIndex, state)
		err := k.heaters[heaterIndex].SetState(state)
		if err != nil {
			log.Printf("[Error] unable to set pin State: %+v\n", err)
		}
	}
}

func NewKeg(tempDev W1Device, heaters []Heater) (*kegStruct, error) {

	temp := tempDev
	return &kegStruct{
		heaters: heaters,
		temp:    temp,
	}, nil
}
