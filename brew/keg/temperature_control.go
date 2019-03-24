package keg

import (
	"log"
	"time"
)

type TempControl interface {
	KeepTemp(temp float64)
	GetTemp() float64
	Stop()
}

type TempControlStruct struct {
	temp       float64
	quit       chan struct{}
	dispresion float64
	started    bool
	kegControl KegControl
}

func NewTempControl(kegControl KegControl, temp float64) TempControl {

	log.Println("Creating temp control")
	tcs := TempControlStruct{
		temp,
		make(chan struct{}),
		2,
		false,
		kegControl,
	}

	return TempControl(&tcs)
}

func (tcs *TempControlStruct) KeepTemp(temp float64) {
	log.Println("new temp to keep is ", temp)
	if !tcs.started {
		ticker := time.NewTicker(5 * time.Second)
		tcs.started = true
		go func() {
			go tcs.loopTemp(ticker)
		}()
	}

	tcs.temp = temp
}

func (tcs TempControlStruct) Stop() {
	tcs.quit <- struct{}{}
	tcs.started = false
}

func (tcs TempControlStruct) loopTemp(ticker *time.Ticker) {

	enableHeaters := func(state HeaterState) {
		log.Println("[tempControl] toggling state of heaters")

		tcs.kegControl.SetHeaterState(FIRST, state)
		tcs.kegControl.SetHeaterState(SECOND, state)
	}

	for {
		select {
		case <-ticker.C:
			currTemp, err := tcs.kegControl.Temperature()
			if err != nil {
				log.Printf("Error while reading temperature. %s", err.Error())
				break
			}

			log.Printf("Current temp is %+v, heater state is %+v\n", currTemp)
			if currTemp+tcs.dispresion > tcs.temp {
				log.Printf("Disabling heaters, temps: %+v > %+v\n", currTemp+tcs.dispresion, tcs.temp)
				enableHeaters(OFF)
			}

			if currTemp-tcs.dispresion < tcs.temp {
				log.Printf("Enabling heaters, temps: %+v < %+v\n", currTemp-tcs.dispresion, tcs.temp)
				enableHeaters(ON)
			}

		case <-tcs.quit:
			ticker.Stop()
			return
		}
	}
}

func (tcs TempControlStruct) GetTemp() float64 {
	return tcs.temp
}
