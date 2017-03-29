package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/piotrjaromin/brew-web/brew/esp8266"
	"github.com/piotrjaromin/brew-web/brew/keg"
	"github.com/piotrjaromin/brew-web/brew/pi"
	"github.com/piotrjaromin/brew-web/brew/web"
)

func main() {

	kegControl, err := getKegControl()
	tempCache := keg.NewTemperatureCache(kegControl, 20, 100)
	tempControl := keg.NewTempControl(kegControl, 20)
	cook := keg.CreateCook(tempControl)

	if err != nil {
		log.Panic("Error while creating keg. Details: ", err)
	}

	mux := http.NewServeMux()

	mux.Handle("/", http.FileServer(http.Dir("public")))
	mux.Handle("/heaters/1", http.HandlerFunc(web.CreateHandlerForHeater(keg.FIRST, kegControl)))
	mux.Handle("/heaters/2", http.HandlerFunc(web.CreateHandlerForHeater(keg.SECOND, kegControl)))
	mux.Handle("/temperatures", http.HandlerFunc(web.CreateTempHandler(tempCache)))
	mux.Handle("/temperatures/control", http.HandlerFunc(web.CreateTempControlHandler(tempControl)))
	mux.Handle("/recepies", http.HandleFunc(web.CreateRecepiesHandler(cook)))

	log.Println("Listening... :3001")
	log.Fatal(http.ListenAndServe(":3001", mux))

}

func getKegControl() (keg.KegControl, error) {

	controllerTypePtr := flag.String("type", "mock", "Defines keg controller type can be mock, esp, pi. Defaults to mock")
	moduleURL := flag.String("url", "http://esp8266.local", "Needed for esp type, provides root url of esp8266")

	flag.Parse()

	switch *controllerTypePtr {
	case "esp":
		log.Println("initializing esp")
		return initEsp8266(*moduleURL)
	case "pi":
		log.Println("initializing pi")
		return initPi()
	default:
		log.Println("Starting mock version")
		return keg.NewKegMock()
	}
}

func initEsp8266(host string) (keg.KegControl, error) {
	return esp8266.NewKeg(host)
}

func initPi() (keg.KegControl, error) {
	devices, devErr := pi.GetDevices()
	if devErr != nil {
		log.Panic("Could not get list of devices. Details: ", devErr)
	}

	if len(devices) != 1 {
		log.Panic("Found wrong amount of 1-wire devices. Got: ", len(devices))
	}

	log.Println("Starting rpio version")
	return pi.NewKeg(devices[0])
}
