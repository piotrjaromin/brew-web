package keg

import (
        "github.com/stianeikeland/go-rpio"
        "log"
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
        Temperature() float32
}

type keg struct {
        heaters []rpio.Pin
        temp    rpio.Pin
}

func (k keg) HeaterState(h Heater) HeaterState {
        return k.heaters[h].Read() == 0
}

func (k keg) ToggleHeater(h Heater) {
        k.heaters[h].Toggle()
}

func (k keg) Temperature() float32 {
        //return float32(k.temp.Read())
        return 5
}

func NewKeg() (KegControl, error) {

        err := rpio.Open()
        if err != nil {
                log.Println("could not open rpio. Details %+v", err)
                return nil, err
        }

        heaters := []rpio.Pin{rpio.Pin(HEATER1_PIN), rpio.Pin(HEATER2_PIN)}
        //temp := rpio.Pin(TEMPERATURE_PIN)
        temp := rpio.Pin(HEATER1_PIN)

        heaters[0].Output()
        heaters[1].Output()
        //temp.Input()

        k := keg{
                heaters: heaters,
                temp:    temp,
        }

        return KegControl(k), nil
}
