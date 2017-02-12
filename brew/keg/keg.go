package keg

import (
        "github.com/stianeikeland/go-rpio"
        "log"
        "github.com/piotrjaromin/brew-web/brew/pi"
)

type Heater int

const HEATER1_PIN = 14
const HEATER2_PIN = 15
const TEMPERATURE_PIN = 10

const (
        FIRST Heater = iota
        SECOND Heater = iota
)

type HeaterState bool

const (
        ON HeaterState = true
        OFF HeaterState = false
)

type KegControl interface {
        ToggleHeater(h Heater)
        HeaterState(h Heater) HeaterState
        Temperature() (float32, error)
}

type keg struct {
        heaters []rpio.Pin
        temp    pi.W1Device
}

func (k keg) HeaterState(h Heater) HeaterState {
        state := k.heaters[h].Read()
        log.Printf("State for heater %+v is %+v", h, state)
        return state == 0
}

func (k keg) ToggleHeater(h Heater) {
        log.Printf("toggle heater: %+v",h)
        k.heaters[h].Toggle()
}

func (k keg) Temperature() (float32, error) {
        t, err := k.temp.Value(1, "t")
        return  float32(t)/ 1000, err
}

func NewKeg(tempDev pi.W1Device) (KegControl, error) {

        err := rpio.Open()
        if err != nil {
                log.Println("could not open rpio. Details %+v", err)
                return nil, err
        }

        heaters := []rpio.Pin{rpio.Pin(HEATER1_PIN), rpio.Pin(HEATER2_PIN)}
        temp := tempDev

        heaters[0].Output()
        heaters[1].Output()

        k := keg{
                heaters: heaters,
                temp:    temp,
        }

        return KegControl(k), nil
}
