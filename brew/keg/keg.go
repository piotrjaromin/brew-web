package keg

type Heater int

const (
	FIRST  Heater = 1
	SECOND Heater = 2
)

type HeaterState bool

const (
	ON  HeaterState = true
	OFF HeaterState = false
)

type KegControl interface {
	ToggleHeater(h Heater)
	HeaterState(h Heater) HeaterState
	Temperature() (float64, error)
}
