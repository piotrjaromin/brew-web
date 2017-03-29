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
}

func NewTempControl(kegControl KegControl, temp float64) TempControl {

	log.Println("Creating temp control")
	tcs := TempControlStruct{
		temp,
		make(chan struct{}),
		1.0,
	}

	ticker := time.NewTicker(5 * time.Second)

	go func() {
		for {
			select {
			case <-ticker.C:

				currTemp, err := kegControl.Temperature()
				if err != nil {
					log.Printf("Error while reading temeprature. %s", err.Error())
					break
				}

				state := kegControl.HeaterState(FIRST)
				if currTemp+tcs.dispresion > tcs.temp && state {
					log.Println("[tempControl] toggling state of heaters")
					kegControl.ToggleHeater(FIRST)
					kegControl.ToggleHeater(SECOND)
				}

				if currTemp-tcs.dispresion < tcs.temp && !state {
					kegControl.ToggleHeater(FIRST)
					kegControl.ToggleHeater(SECOND)
				}

			case <-tcs.quit:
				ticker.Stop()
				return
			}
		}
	}()

	return TempControl(&tcs)
}

func (tcs *TempControlStruct) KeepTemp(temp float64) {
	log.Println("new temp to keep is ", temp)
	tcs.temp = temp
}

func (tcs TempControlStruct) Stop() {
	tcs.quit <- struct{}{}
}

func (tcs TempControlStruct) GetTemp() float64 {
	return tcs.temp
}
