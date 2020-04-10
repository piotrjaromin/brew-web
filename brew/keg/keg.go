package keg

import (
	"fmt"
	"log"

	"github.com/piotrjaromin/brew-web/brew/config"
	"github.com/piotrjaromin/brew-web/brew/keg/controllers/mock"
	"github.com/piotrjaromin/brew-web/brew/keg/controllers/pi"
)

type KegControl interface {
	SetHeaterPower(float64)
	GetHeaterPower() float64
	Temperature() (float64, error)
}

func CreateKegControl(controllerType string, kegConfig config.Keg) (KegControl, error) {
	switch controllerType {
	case "pi":
		log.Println("initializing pi")
		return initPi(kegConfig)
	case "mock":
		log.Println("Starting mock version")
		return mock.NewKegMock()
	default:
		return nil, fmt.Errorf("Unsupported keg controller type %s", controllerType)
	}
}

func initPi(kegConfig config.Keg) (KegControl, error) {
	devices, devErr := pi.GetDevices()
	if devErr != nil {
		return nil, fmt.Errorf("Could not get list of devices. Details: %s", devErr.Error())
	}

	if len(devices) != 1 {
		return nil, fmt.Errorf("Found wrong amount of 1-wire devices. Got: %d", len(devices))
	}

	heaters, err := pi.GetHeaters(kegConfig)
	if err != nil {
		return nil, fmt.Errorf("Unable to initialize heaters. Got: %s", err.Error())
	}

	log.Println("Starting rpio version")
	return pi.NewKeg(devices[0], heaters)
}
