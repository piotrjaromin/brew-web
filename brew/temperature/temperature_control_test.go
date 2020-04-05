package temperature

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockKeg struct {
	mock.Mock
	power float64
	temp  float64
}

func (m mockKeg) SetHeaterPower(f float64) {
	m.power = f
	m.Called(f)
}

func (m mockKeg) GetHeaterPower() float64 {
	m.Called()
	return m.power
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
		power: 0.0,
	}

	keg.On("Temperature").Return(initTemp)

	// NOT call heaters when temp is right
	updateHeaters(keg, initTemp, delta)
	updateHeaters(keg, initTemp, delta)
	keg.AssertNotCalled(t, "SetHeaterState")

	// // Call heaters keg.ON when temp is higher than delta
	keg.On("SetHeaterPower", 1.0).Return()

	tmpHigherThanDelta := initTemp + delta + 1
	updateHeaters(keg, tmpHigherThanDelta, delta)
	updateHeaters(keg, tmpHigherThanDelta, delta)

	// // Call heaters keg.OFF when temp is lower than delta
	keg.On("SetHeaterPower", 0.0).Return()

	tmpLowerThanDelta := initTemp - delta - 1
	updateHeaters(keg, tmpLowerThanDelta, delta)
	updateHeaters(keg, tmpLowerThanDelta, delta)

	keg.AssertExpectations(t)
}

func TestKeepTemp(t *testing.T) {
	initTemp := 20.0
	keg := &mockKeg{
		temp:  initTemp,
		power: 0.0,
	}

	tcs := NewTempControl(keg, initTemp)

	assert.Equal(t, initTemp, tcs.GetTemp(), "init temp should be equal set temp")

	tcs.KeepTemp(50.0)
	assert.Equal(t, 50.0, tcs.GetTemp(), "should keep correct temp")

	tcs.KeepTemp(80.0)
	assert.Equal(t, 80.0, tcs.GetTemp(), "should allows temp change")

	tcs.Stop()
}
