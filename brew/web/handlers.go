package web

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/piotrjaromin/brew-web/brew/keg"
)

func CreateHandlerForHeater(heater keg.Heater, kegControl keg.KegControl) http.HandlerFunc {

	writeState := func(rw http.ResponseWriter) {

		state, err := json.Marshal(struct {
			State keg.HeaterState `json:"state"`
		}{
			kegControl.HeaterState(heater),
		})

		if err != nil {
			HandlerError(rw, err)
			return
		}

		rw.Write(state)
	}

	return func(rw http.ResponseWriter, req *http.Request) {

		switch req.Method {
		case http.MethodGet:
			writeState(rw)
		case http.MethodPost:
			kegControl.ToggleHeater(heater)
			writeState(rw)
		}
	}

}

func CreateTempHandler(t keg.Temperatures) http.HandlerFunc {

	return func(rw http.ResponseWriter, req *http.Request) {
		tempsJson, err := json.Marshal(t.Get())

		if err != nil {
			HandlerError(rw, err)
			return
		}

		rw.Write(tempsJson)
	}

}

func CreateTempControlHandler(tempControl keg.TempControl) http.HandlerFunc {

	return func(rw http.ResponseWriter, req *http.Request) {

		switch req.Method {
		case http.MethodGet:
			tempJSON, err := json.Marshal(Temp{tempControl.GetTemp()})

			if err != nil {
				HandlerError(rw, err)
				return
			}

			rw.Write(tempJSON)

		case http.MethodPost:
			decoder := json.NewDecoder(req.Body)
			temp := Temp{}
			if err := decoder.Decode(&temp); err != nil {
				HandlerError(rw, err)
				return
			}
			tempControl.KeepTemp(temp.Value)

		case http.MethodDelete:
			tempControl.Stop()
		}
	}
}

func CreateRecipesHandler(cook keg.Cook) http.HandlerFunc {

	var recipe keg.RecipeStruct
	return func(rw http.ResponseWriter, req *http.Request) {

		switch req.Method {
		case http.MethodGet:
			recipeJSON, err := json.Marshal(recipe)
			if err != nil {
				HandlerError(rw, err)
				return
			}

			rw.Write(recipeJSON)

		case http.MethodPost:
			decoder := json.NewDecoder(req.Body)

			if err := decoder.Decode(&recipe); err != nil {
				HandlerError(rw, err)
				return
			}
			cook.Execute(recipe)

		case http.MethodDelete:
			cook.Stop()
		}
	}
}

func HandlerError(rw http.ResponseWriter, err error) {
	log.Fatal("Error while handling request. ", err)
	rw.Write([]byte("error: " + err.Error()))
	rw.WriteHeader(http.StatusInternalServerError)
}
