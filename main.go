package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/piotrjaromin/brew-web/brew/keg"
	"github.com/piotrjaromin/brew-web/brew/recepies"

	"github.com/piotrjaromin/brew-web/brew/config"
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
	controllerTypePtr := flag.String("type", "mock", "Defines keg controller type can be mock, esp, pi. Defaults to mock")
	flag.Parse()

	conf, err := config.GetConfig("./config.yml")
	if err != nil {
		log.Panic("Error while reading config. Details: ", err)
	}

	fmt.Printf("%+v", *conf)

	kegControl, err := keg.CreateKegControl(*controllerTypePtr, conf.Keg)
	if err != nil {
		log.Panic("Error while creating keg. Details: ", err)
	}

	tempStore, err := temperature.NewTemperatureStore(kegControl, conf.Temperature)
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

	web.InitHeater(e, kegControl)
	web.InitTemp(e, tempStore)
	web.InitTempControl(e, tempControl)
	web.InitRecipes(e, cook)

	fileServer := http.FileServer(statikFS)
	e.GET("/*", func(c echo.Context) error {
		fileServer.ServeHTTP(c.Response().Writer, c.Request())
		return nil
	})

	log.Printf("Version: %s, commit: %s", Version, ShortCommit)
	log.Printf("Listening... %s \n", conf.Port)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", conf.Port)))
}
