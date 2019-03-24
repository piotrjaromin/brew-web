package web

import (
	"github.com/piotrjaromin/brew-web/brew/keg"
)

type Temp struct {
	Value float64 `json:"value"`
}

type HeaterState struct {
	State keg.HeaterState `json:"state"`
}
