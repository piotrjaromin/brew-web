package recepies

import (
	"time"

	"github.com/piotrjaromin/brew-web/brew/temperature"
)

type Elapsed int

type RecipeStruct struct {
	States map[Elapsed]float64 `json:"states"`
}

type Recipe interface {
	TempForTime(t Elapsed) (float64, bool)
}

type Cook interface {
	Execute(rs Recipe)
	Stop()
}

type CookStruct struct {
	tempControl temperature.TempControl
	quit        chan struct{}
}

func (c CookStruct) Stop() {
	c.quit <- struct{}{}
}

func (c CookStruct) Execute(r Recipe) {

	ticker := time.NewTicker(15 * time.Second)
	start := time.Now()

	go func() {
		for {
			select {
			case <-ticker.C:
				diffTime := time.Now().Sub(start).Minutes()
				temp, isDone := r.TempForTime(Elapsed(diffTime))

				if isDone {
					c.Stop()
					return
				}

				c.tempControl.KeepTemp(temp)
			case <-c.quit:
				ticker.Stop()
				return
			}
		}
	}()
}

func (r RecipeStruct) TempForTime(t Elapsed) (float64, bool) {

	for elapsed, temp := range r.States {
		if t <= elapsed {
			return temp, false
		}
	}

	return 0, true
}

func CreateCook(tempControl temperature.TempControl) Cook {

	return CookStruct{tempControl, make(chan struct{})}
}
