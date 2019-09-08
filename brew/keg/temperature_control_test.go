package keg

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
	m.Called()
	return m.temp, nil
}

func TestTemperatureControl(t *testing.T) {
	delta := 2.0
	initTemp := 20.0
	keg := &mockKeg{
		temp:  initTemp,
		state: OFF,
	}

	keg.On("Temperature").Return(initTemp)

	// NOT call heaters when temp is right
	updateHeaters(keg, initTemp, delta)
	updateHeaters(keg, initTemp, delta)
	keg.AssertNotCalled(t, "SetHeaterState")

	// // Call heaters on when temp is higher than delta
	keg.On("SetHeaterState", FIRST, ON).Return()
	keg.On("SetHeaterState", SECOND, ON).Return()

	tmpHigherThanDelta := initTemp + delta + 1
	updateHeaters(keg, tmpHigherThanDelta, delta)
	updateHeaters(keg, tmpHigherThanDelta, delta)

	// // Call heaters off when temp is lower than delta
	keg.On("SetHeaterState", FIRST, OFF).Return()
	keg.On("SetHeaterState", SECOND, OFF).Return()

	tmpLowerThanDelta := initTemp - delta - 1
	updateHeaters(keg, tmpLowerThanDelta, delta)
	updateHeaters(keg, tmpLowerThanDelta, delta)

	keg.AssertExpectations(t)
}

func TestKeepTemp(t *testing.T) {
	initTemp := 20.0
	keg := &mockKeg{
		temp:  initTemp,
		state: OFF,
	}

	tcs := NewTempControl(keg, initTemp)

	assert.Equal(t, initTemp, tcs.GetTemp(), "init temp should be equal set temp")

	tcs.KeepTemp(50.0)
	assert.Equal(t, 50.0, tcs.GetTemp(), "should keep correct temp")

	tcs.KeepTemp(80.0)
	assert.Equal(t, 80.0, tcs.GetTemp(), "should allows temp change")

	tcs.Stop()
}
