package keg

type Heater int

const (
	FIRST  Heater = 0
	SECOND Heater = 1
)

type HeaterState bool

const (
	ON  HeaterState = true
	OFF HeaterState = false
)

type KegControl interface {
	ToggleHeater(h Heater)
	SetHeaterState(h Heater, enabled HeaterState)
	HeaterState(h Heater) HeaterState
	Temperature() (float64, error)
}
