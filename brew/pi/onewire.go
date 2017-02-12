package pi

import (
        "os/exec"
        "fmt"
        "io/ioutil"
        "strings"
        "strconv"
)

const W1_DEVICES_PATH = "/sys/bus/w1/devices/"
const W1_SLAVE = "w1_slave"
const W28_PREFIX = "28"

type W1Device interface {
        Name() string
        W1Slave() (string, error)
        Value(line int, key string) (int, error)
}

type w1Struct struct {
        name string
        path string
}

func (w1 w1Struct) Name() string {
        return w1.name
}

func (w1 w1Struct) W1Slave() (string, error) {
        data, err := ioutil.ReadFile(w1.path + "/" + W1_SLAVE)
        if err != nil {
                return "", fmt.Errorf("[W1] Could not read W1Slave data file. Details %+v.", err)
        }

        return string(data), nil
}

func (w1 w1Struct) Value(lineNo int, key string) (int, error) {

        v1slave, err := w1.W1Slave()
        if err != nil {
                return 0, err
        }

        lines := strings.Split(v1slave, "\n")

        if len(lines) < lineNo {
                return 0, fmt.Errorf("[W1] w1slave does not contains only %d lines(provided %d).", len(lines), lineNo)
        }

        line := lines[lineNo]
        keyEntry := key + "="
        valIndex := strings.Index(line, keyEntry)
        if valIndex == -1 {
                return 0, fmt.Errorf("[W1] w1slave line %d does not contain key %s.", lineNo, key)
        }
        valIndex += len(keyEntry)
        return strconv.Atoi(strings.Split(line[valIndex:], " ")[0])
}

func Init() error {

        //this one probably require being called with sudo
        if err := callOs("sudo modprobe w1-gpio'"); err != nil {
                return err
        }

        if err := callOs("sudo modprobe w1-therm"); err != nil {
                return err
        }

        return nil
}

func GetDevices() ([]W1Device, error) {
        devices := make([]W1Device, 0, 1)

        dirs, err := ioutil.ReadDir(W1_DEVICES_PATH)
        if err != nil {
                return devices, fmt.Errorf("[W1] could not read directory %s. Details %+v", W1_DEVICES_PATH, err)
        }

        for _, w1dir := range dirs {

                if (strings.HasPrefix(w1dir.Name(), W28_PREFIX)) {
                        devices = append(devices, w1Struct{
                                name: w1dir.Name(),
                                path: W1_DEVICES_PATH + "/" + w1dir.Name(),
                        })
                }
        }

        return devices, nil
}

func callOs(command string) error {

        cmd := exec.Command(command)
        err := cmd.Run()
        if err != nil {
                return fmt.Errorf("[W1] Could not init 1-wire for command %s. Details: %+v", command, err)
        }

        return nil
}
