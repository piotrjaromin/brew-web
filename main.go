package main

import (
        "net/http"
        "log"
        "github.com/piotrjaromin/brew-web/brew/keg"
        "github.com/piotrjaromin/brew-web/brew/web"
        "os"
        "github.com/piotrjaromin/brew-web/brew/pi"
)

func main() {

        kegControl, err := getKegControl()
        tempCache := keg.NewTemperatureCache(kegControl, 20, 100)

        if err != nil {
                log.Panic("Error while creating keg. Details: ", err)
        }
        
        mux := http.NewServeMux()

        mux.Handle("/", http.FileServer(http.Dir("public")))
        mux.Handle("/heaters/1", http.HandlerFunc(web.CreateHandlerForHeater(keg.FIRST, kegControl)))
        mux.Handle("/heaters/2", http.HandlerFunc(web.CreateHandlerForHeater(keg.SECOND, kegControl)))
        mux.Handle("/temperatures", http.HandlerFunc(web.CreateTempHandler(tempCache)))

        log.Println("Listening...")
        log.Fatal(http.ListenAndServe(":3001", mux))

}

func getKegControl() (keg.KegControl, error) {

        if len(os.Args) > 1 && os.Args[1] == "mock" {
                log.Println("Starting mock version")
                return keg.NewKegMock()
        }

        //PI initialization
        //if err := pi.Init(); err !=nil {
        //        log.Panic("Could not initialize pi one wire module. Details: ", err)
        //}

        devices, devErr := pi.GetDevices()
        if devErr != nil {
                log.Panic("Could not get list of devices. Details: ", devErr)
        }

        if len(devices) != 1 {
                log.Panic("Found wrong amount of 1-wire devices. Got: ", len(devices))
        }

        log.Println("Starting rpio version")
        return keg.NewKeg(devices[0])
}