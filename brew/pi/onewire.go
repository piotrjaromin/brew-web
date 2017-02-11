package pi

import (
        "os/exec"
        "fmt"
)

func CreateOneWire(pin int) error {

        if err := callOs("modprobe w1-gpio'"); err != nil {
                return err
        }

        if err := callOs("modprobe w1-therm"); err != nil {
                return err
        }


        //TODO
        //baseDir := "/sys/bus/w1/devices/"
        //deviceFolder := glob.glob(base_dir + '28*')[0]
        //deviceFile := device_folder + "/w1_slave"


        return nil
}

func callOs(command string) error {

        cmd := exec.Command(command)
        err := cmd.Run()
        if err != nil {
                return fmt.Errorf("Could not init 1-wire for command %s. Details: %+v", command, err)
        }

        return nil
}
