package main

import (
	"flag"
	"log"
	"net/http"

	"net"
	"os"

	"github.com/piotrjaromin/brew-web/brew/esp8266"
	"github.com/piotrjaromin/brew-web/brew/keg"
	"github.com/piotrjaromin/brew-web/brew/pi"
	"github.com/piotrjaromin/brew-web/brew/web"

	"github.com/rakyll/statik/fs"

	_ "github.com/piotrjaromin/brew-web/statik"
)

func main() {

	kegControl, err := getKegControl()
	tempCache := keg.NewTemperatureStore(kegControl, 20, 100)
	tempControl := keg.NewTempControl(kegControl, 20)
	cook := keg.CreateCook(tempControl)

	if err != nil {
		log.Panic("Error while creating keg. Details: ", err)
	}

	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()

	route(mux, "/", http.FileServer(statikFS))
	route(mux, "/heaters/1", web.CreateHandlerForHeater(keg.FIRST, kegControl))
	route(mux, "/heaters/2", web.CreateHandlerForHeater(keg.SECOND, kegControl))
	route(mux, "/temperatures", web.CreateTempHandler(tempCache))
	route(mux, "/temperatures/control", web.CreateTempControlHandler(tempControl))
	route(mux, "/recipes", web.CreateRecipesHandler(cook))

	log.Println("Listening... :3001")
	log.Fatal(http.ListenAndServe(":3001", mux))

}

func route(mux *http.ServeMux, path string, handler http.Handler) {
	mux.Handle(path, corsHandler(handler))
}

func corsHandler(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding")

		if r.Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, r)
	}
}

func getKegControl() (keg.KegControl, error) {

	controllerTypePtr := flag.String("type", "mock", "Defines keg controller type can be mock, esp, pi. Defaults to mock")
	moduleURL := flag.String("url", "esp8266.local", "Needed for esp type, provides root url of esp8266")
	protocol := flag.String("protocol", "http://", "protcol at which esp8266 works")

	flag.Parse()

	switch *controllerTypePtr {
	case "esp":
		log.Println("initializing esp")
		return initEsp8266(*moduleURL, *protocol)
	case "pi":
		log.Println("initializing pi")
		return initPi()
	case "mock":
		log.Println("Starting mock version")
		return keg.NewKegMock()
	default:
		flag.PrintDefaults()
		os.Exit(0)
		return nil, nil
	}
}

func initEsp8266(host, protocol string) (keg.KegControl, error) {

	ips, err := net.LookupIP(host)
	if err != nil {
		return nil, err
	}
	return esp8266.NewKeg(protocol + ips[0].String())
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
