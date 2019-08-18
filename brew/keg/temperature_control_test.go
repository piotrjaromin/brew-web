package keg

import (
	"time"

	"testing"

	"github.com/stretchr/testify/mock"
)

type mockKeg struct {
	mock.Mock
	state HeaterState
	temp  float64
}

func (m mockKeg) ToggleHeater(h Heater) {
	m.Called(h)
}

func (m mockKeg) SetHeaterState(h Heater, enabled HeaterState) {
	m.Called(h, enabled)
}

func (m mockKeg) HeaterState(h Heater) HeaterState {
	m.Called(h)
	return m.state
}

func (m mockKeg) Temperature() (float64, error) {
	return m.temp, nil
}

func TestTemperatureControl(t *testing.T) {
	intervalSleep := 3 * time.Millisecond

	initTemp := 20.0
	keg := &mockKeg{
		temp:  initTemp,
		state: OFF,
	}

	keg.On("Temperature").Return(initTemp)
	keg.On("SetHeaterState", FIRST, ON).Return()
	keg.On("SetHeaterState", FIRST, ON).Return()
	keg.On("SetHeaterState", SECOND, OFF).Return()
	keg.On("SetHeaterState", SECOND, OFF).Return()

	tcs := TempControlStruct{
		initTemp,
		make(chan struct{}),
		2,
		false,
		keg,
		1 * time.Millisecond,
	}

	higherTemp := initTemp + 10
	tcs.KeepTemp(higherTemp)
	time.Sleep(intervalSleep)

	keg.AssertCalled(t, "SetHeaterState", SECOND, ON)
	keg.AssertCalled(t, "SetHeaterState", FIRST, ON)

	tcs.temp = higherTemp + 5
	time.Sleep(intervalSleep)

	// keg.AssertCalled(t, "SetHeaterState", FIRST, OFF)
	// keg.AssertCalled(t, "SetHeaterState", SECOND, OFF)

	tcs.Stop()
}
