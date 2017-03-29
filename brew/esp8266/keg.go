package esp8266

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/piotrjaromin/brew-web/brew/keg"
)

type kegStruct struct {
	heatersPath []string
	host        string
	tempPath    string
}

func (k kegStruct) HeaterState(h keg.Heater) keg.HeaterState {
	stateStr, err := readResp(k.host + k.heatersPath[h])
	state, err := strconv.Atoi(stateStr)
	if err != nil {
		log.Panic("Could not read heater state", err.Error())
	}

	return state != 0
}

func (k kegStruct) ToggleHeater(h keg.Heater) {
	log.Printf("toggle heater: %+v", h)
	resp, err := http.Post(k.host+k.heatersPath[h], "text", strings.NewReader(""))
	if err != nil {
		log.Panic("Could not switch heater state", err.Error())
	}

	if resp.StatusCode != 200 {
		log.Panic("Could not switch state, wrong response code")
	}
}

func (k kegStruct) Temperature() (float64, error) {
	temp, err := readResp(k.host + k.tempPath)
	if err != nil {
		return 0, err
	}

	return strconv.ParseFloat(temp, 64)
}

func readResp(path string) (string, error) {
	resp, err := http.Get(path)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return "", errors.New("wrong status code returned")
	}

	body, errRead := ioutil.ReadAll(resp.Body)
	if errRead != nil {
		return "", errRead
	}

	return string(body), nil
}

func NewKeg(moduleUrl string) (keg.KegControl, error) {
	log.Println("Esp module will connect to " + moduleUrl)

	k := kegStruct{
		make([]string, 2, 2),
		moduleUrl,
		"/temp",
	}

	k.heatersPath[0] = "/heater1"
	k.heatersPath[1] = "/heater2"

	return keg.KegControl(k), nil
}
