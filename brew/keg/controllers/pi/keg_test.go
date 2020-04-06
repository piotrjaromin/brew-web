package pi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockHeater struct {
	state HeaterState
}

func (h *mockHeater) SetState(state HeaterState) error {
	h.state = state
	return nil
}

func (h *mockHeater) State() HeaterState {
	return h.state
}

type mockW1Device struct{}

func (w mockW1Device) Name() string {
	return "mockW1Name"
}

func (w mockW1Device) W1Slave() (string, error) {
	return "mockW1Slave", nil
}

func (w mockW1Device) Value(line int, key string) (int, error) {
	return 0, nil
}

func TestGetHeaterPower(t *testing.T) {
	w1 := mockW1Device{}
	h1 := &mockHeater{HEATER_DISABLED}
	h2 := &mockHeater{HEATER_DISABLED}
	h3 := &mockHeater{HEATER_DISABLED}
	h4 := &mockHeater{HEATER_DISABLED}

	heaters := []Heater{h1, h2, h3, h4}
	mockKeg, err := NewKeg(w1, heaters)
	assert.Nil(t, err)

	power := mockKeg.GetHeaterPower()
	assert.Equal(t, 0.0, power)

	h1.state = HEATER_ENABLED
	power = mockKeg.GetHeaterPower()
	assert.Equal(t, 0.25, power)

	h2.state = HEATER_ENABLED
	h3.state = HEATER_ENABLED
	h4.state = HEATER_ENABLED
	power = mockKeg.GetHeaterPower()
	assert.Equal(t, 1.0, power)
}

func TestSetHeaterPower(t *testing.T) {
	w1 := mockW1Device{}
	h1 := &mockHeater{HEATER_DISABLED}
	h2 := &mockHeater{HEATER_DISABLED}

	heaters := []Heater{h1, h2}
	mockKeg, err := NewKeg(w1, heaters)
	assert.Nil(t, err)

	mockKeg.SetHeaterPower(1.0)
	assert.Equal(t, HEATER_ENABLED, h1.state)
	assert.Equal(t, HEATER_ENABLED, h2.state)

	mockKeg.SetHeaterPower(0.5)
	assert.Equal(t, HEATER_ENABLED, h1.state)
	assert.Equal(t, HEATER_DISABLED, h2.state)

	mockKeg.SetHeaterPower(0.0)
	assert.Equal(t, HEATER_DISABLED, h1.state)
	assert.Equal(t, HEATER_DISABLED, h2.state)
}
