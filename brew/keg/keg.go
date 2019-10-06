package keg

import (
	"flag"
	"log"
	"os"

	"github.com/piotrjaromin/brew-web/brew/config"
	"github.com/piotrjaromin/brew-web/brew/keg/controllers/mock"
	"github.com/piotrjaromin/brew-web/brew/keg/controllers/pi"
)

type Heater int

const (
	FIRST  Heater = 1
	SECOND Heater = 2
)

type HeaterState bool

const (
	ON  HeaterState = true
	OFF HeaterState = false
)

type KegControl interface {
	ToggleHeater(h Heater)
	SetHeaterState(h Heater, enabled HeaterState)
	HeaterState(h Heater) HeaterState
	Temperature() (float64, error)
}

func getKegControl(controllerType string, kegConfig config.Keg) (KegControl, error) {
	switch controllerType {
	case "pi":
		log.Println("initializing pi")
		return initPi(kegConfig)
	case "mock":
		log.Println("Starting mock version")
		return mock.NewKegMock()
	default:
		flag.PrintDefaults()
		os.Exit(0)
		return nil, nil
	}
}

func initPi(kegConfig config.Keg) (KegControl, error) {
	devices, devErr := pi.GetDevices()
	if devErr != nil {
		log.Panic("Could not get list of devices. Details: ", devErr)
	}

	if len(devices) != 1 {
		log.Panic("Found wrong amount of 1-wire devices. Got: ", len(devices))
	}

	log.Println("Starting rpio version")
	return pi.NewKeg(devices[0], kegConfig)
}
