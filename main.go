package main

import (
	"flag"
	"log"
	"net/http"

	"net"
	"os"

	"github.com/piotrjaromin/brew-web/brew/keg"
	"github.com/piotrjaromin/brew-web/brew/keg/controllers/esp8266"
	"github.com/piotrjaromin/brew-web/brew/keg/controllers/mock"
	"github.com/piotrjaromin/brew-web/brew/keg/controllers/pi"
	"github.com/piotrjaromin/brew-web/brew/recepies"

	"github.com/piotrjaromin/brew-web/brew/temperature"
	"github.com/piotrjaromin/brew-web/brew/web"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/rakyll/statik/fs"

	_ "github.com/piotrjaromin/brew-web/statik"
)

var (
	ShortCommit string
	Version     string
)

func main() {
	kegControl, err := getKegControl()
	if err != nil {
		log.Panic("Error while creating keg. Details: ", err)
	}

	tempStore, err := temperature.NewTemperatureStore(kegControl, 20, 100)
	if err != nil {
		log.Panic("Error while creating tempStore. Details: ", err)
	}

	tempControl := temperature.NewTempControl(kegControl, 20)
	cook := recepies.CreateCook(tempControl)

	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	web.InitHeater(e, keg.FIRST, kegControl)
	web.InitHeater(e, keg.SECOND, kegControl)
	web.InitTemp(e, tempStore)
	web.InitTempControl(e, tempControl)
	web.InitRecipes(e, cook)

	fileServer := http.FileServer(statikFS)
	e.GET("/*", func(c echo.Context) error {
		fileServer.ServeHTTP(c.Response().Writer, c.Request())
		return nil
	})

	log.Printf("Version: %s, commit: %s", Version, ShortCommit)
	log.Println("Listening... :3001")
	e.Logger.Fatal(e.Start(":3001"))
}

func getKegControl() (keg.KegControl, error) {
	controllerTypePtr := flag.String("type", "mock", "Defines keg controller type can be mock, esp, pi. Defaults to mock")
	moduleURL := flag.String("url", "esp8266.local", "Needed for esp type, provides root url of esp8266")
	protocol := flag.String("protocol", "http://", "protocol at which esp8266 works")

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
		return mock.NewKegMock()
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
