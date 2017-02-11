package web

import (
	"github.com/piotrjaromin/brew-web/brew/keg"
	"net/http"
	"encoding/json"
	"log"
)


func CreateHandlerForHeater(heater keg.Heater, kegControl keg.KegControl) http.HandlerFunc {

	writeState := func(rw http.ResponseWriter) {

		state, err := json.Marshal(struct {
			State keg.HeaterState `json:"state"`
		}{
			kegControl.HeaterState(heater),
		})

		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write([]byte("error: " + err.Error()))
		} else {
			rw.Write(state)
		}
	}

	return func(rw http.ResponseWriter, req *http.Request) {

		switch req.Method {
		case "GET":
			writeState(rw)
		case "POST":
			kegControl.ToggleHeater(heater)
			writeState(rw)
		}
	}

}

func CreateTempHandler(t keg.Temperatures) http.HandlerFunc {

	return func(rw http.ResponseWriter, req *http.Request) {

		tempsJson, err := json.Marshal(t.Get())

		if err != nil {
			log.Fatal("error while writing temps array. ", err)
			rw.Write([]byte("error: " + err.Error()))
			rw.WriteHeader(http.StatusInternalServerError)
		} else {
			rw.Write(tempsJson)
		}

	}
}
