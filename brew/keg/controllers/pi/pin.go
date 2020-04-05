package pi

import (
	"io/ioutil"
	"os"
	"strconv"
)

type pinState bool

var PIN_HIGH = pinState(true)
var PIN_LOW = pinState(false)

var low = []byte("0")
var high = []byte("1")

type pin struct {
	pinFile *os.File
}

func newPin(pinBCM int) (pin, error) {
	pinStr := strconv.Itoa(pinBCM)
	// Export the GPIO kernel object for GPIO to user space via sysfs
	ioutil.WriteFile("/sys/class/gpio/export", []byte(pinStr), 0644)

	//Set to output mode
	ioutil.WriteFile("/sys/class/gpio/gpio"+pinStr+"/direction", []byte("out"), 0644)

	//Open the value file
	pinFile, err := os.Open("/sys/class/gpio/gpio" + pinStr + "/value")
	if err != nil {
		return pin{}, err
	}

	return pin{
		pinFile: pinFile,
	}, nil
}

func (p pin) High() {
	p.pinFile.Write(high)
}

func (p pin) Low() {
	p.pinFile.Write(low)
}

func (p pin) Read() (pinState, error) {
	buffer := make([]byte, 1)
	_, err := p.pinFile.Read(buffer)
	if err != nil {
		return PIN_LOW, err
	}

	if buffer[0] == low[0] {
		return PIN_LOW, nil
	}

	return PIN_HIGH, nil
}
