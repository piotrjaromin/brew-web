package main

import (
        "net/http"
        "log"
        "github.com/piotrjaromin/brew-web/brew/keg"
        "github.com/piotrjaromin/brew-web/brew/web"
        "os"
)

func main() {

        kegControl, err := getKegControl()
        tempCache := keg.NewTemperatureCache(kegControl, 20, 100)

        if err != nil {
                log.Fatal("Error while creating keg. Details: ", err)
                return
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

        log.Println("Startin rpio version")
        return keg.NewKeg()
}