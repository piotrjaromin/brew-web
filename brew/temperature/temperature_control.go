package temperature

import (
	"log"
	"time"

	"github.com/piotrjaromin/brew-web/brew/keg"
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
	kegControl keg.KegControl
	interval   time.Duration
}

func NewTempControl(kegControl keg.KegControl, temp float64) TempControl {

	log.Println("Creating temp control")
	tcs := TempControlStruct{
		temp,
		make(chan struct{}),
		1,
		false,
		kegControl,
		5 * time.Second,
	}

	return TempControl(&tcs)
}

func (tcs *TempControlStruct) KeepTemp(temp float64) {
	log.Println("new temp to keep is ", temp)
	if !tcs.started {
		ticker := time.NewTicker(tcs.interval)
		tcs.started = true
		go func() {
			go tcs.loopTemp(ticker)
		}()
	}

	tcs.temp = temp
}

func (tcs *TempControlStruct) Stop() {
	tcs.quit <- struct{}{}
	tcs.started = false
}

func (tcs *TempControlStruct) loopTemp(ticker *time.Ticker) {
	for {
		select {
		case <-ticker.C:
			updateHeaters(tcs.kegControl, tcs.temp, tcs.dispresion)
		case <-tcs.quit:
			ticker.Stop()
			return
		}
	}
}

func updateHeaters(kegControl keg.KegControl, desiredTmp, deltaTmp float64) {
	enableHeaters := func(state keg.HeaterState) {
		kegControl.SetHeaterState(keg.FIRST, state)
		kegControl.SetHeaterState(keg.SECOND, state)
	}

	currTemp, err := kegControl.Temperature()
	if err != nil {
		log.Printf("Error while reading temperature. %s", err.Error())
		return
	}

	log.Printf("Current temp is %+v, desired temp is %+v\n", currTemp, desiredTmp)
	if currTemp < desiredTmp-deltaTmp {
		log.Printf("Enabling heaters, temps: %+v < %+v\n", currTemp-deltaTmp, desiredTmp)
		enableHeaters(keg.ON)
	}

	if currTemp > desiredTmp+deltaTmp {
		log.Printf("Disabling heaters, temps: %+v > %+v\n", currTemp+deltaTmp, desiredTmp)
		enableHeaters(keg.OFF)
	}

}

func (tcs TempControlStruct) GetTemp() float64 {
	return tcs.temp
}
