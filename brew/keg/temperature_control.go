package keg

import (
	"log"
	"time"
)

type TempControl interface {
	KeepTemp(temp float64)
	GetTemp() float64
	Stop()
	Start()
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
		1.0,
		false,
		kegControl,
	}

	return TempControl(&tcs)
}

func (tcs *TempControlStruct) KeepTemp(temp float64) {
	log.Println("new temp to keep is ", temp)
	if !tcs.started {
		tcs.Start()
	}

	tcs.temp = temp
}

func (tcs TempControlStruct) Stop() {
	tcs.quit <- struct{}{}
	tcs.started = false
}

func (tcs TempControlStruct) Start() {
	ticker := time.NewTicker(5 * time.Second)
	tcs.started = true
	go func() {
		go tcs.loopTemp(ticker)
	}()
}

func (tcs TempControlStruct) loopTemp(ticker *time.Ticker) {
	for {
		select {
		case <-ticker.C:

			currTemp, err := tcs.kegControl.Temperature()
			if err != nil {
				log.Printf("Error while reading temeprature. %s", err.Error())
				break
			}

			state := tcs.kegControl.HeaterState(FIRST)
			if currTemp+tcs.dispresion > tcs.temp && state {
				log.Println("[tempControl] toggling state of heaters")
				tcs.kegControl.ToggleHeater(FIRST)
				tcs.kegControl.ToggleHeater(SECOND)
			}

			if currTemp-tcs.dispresion < tcs.temp && !state {
				tcs.kegControl.ToggleHeater(FIRST)
				tcs.kegControl.ToggleHeater(SECOND)
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
