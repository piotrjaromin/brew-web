package pi

import (
	"log"
	"math"

	"github.com/piotrjaromin/brew-web/brew/config"

	rpio "github.com/stianeikeland/go-rpio"
)

type kegStruct struct {
	heaters      []rpio.Pin
	heaterAmount int
	temp         W1Device
}

func (k *kegStruct) GetHeaterPower() float64 {
	sum := 0.0

	for _, heater := range k.heaters {
		if heater.Read() == rpio.High {
			sum = +1
		}
	}

	return sum / float64(k.heaterAmount)
}

func (k *kegStruct) Temperature() (float64, error) {
	t, err := k.temp.Value(1, "t")
	return float64(t) / 1000, err
}

func (k *kegStruct) SetHeaterPower(power float64) {
	val := int(math.Round(power * float64(k.heaterAmount)))

	for heaterIndex := 0; heaterIndex <= val; heaterIndex++ {
		k.heaters[heaterIndex].High()
	}
}

func NewKeg(tempDev W1Device, c config.Keg) (*kegStruct, error) {
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

	return &kegStruct{
		heaters:      heaters,
		temp:         temp,
		heaterAmount: len(heaters),
	}, nil
}
