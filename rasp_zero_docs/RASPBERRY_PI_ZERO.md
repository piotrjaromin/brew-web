
# RASPBERRY PI ZERO

## Prepare sd card

to enable ssh drop `ssh` file to root of sd card.
to enable wifi drop `wpa_supplicant.conf` on root of sd card.

```wpa_supplicant.conf
ctrl_interface=DIR=/var/run/wpa_supplicant GROUP=netdev
update_config=1
country=PL # Country code

network={
    ssid="SS_ID_OF_WIFI"
    psk="WIFI_PASSWORD"
}
```




## W1 DS18B20 temperature sensor

[source](https://pinout.xyz/pinout/1_wire)

To enable the one-wire interface you need to add the following line to /boot/config.txt, before rebooting your Pi:

```
dtoverlay=w1-gpio
```

```
dtoverlay w1-gpio gpiopin=4
lsmod
```

or

```
dtoverlay=w1-gpio,gpiopin=x
```

if you would like to use a custom pin (default is BCM4, as illustrated in pinout herein).

Alternatively you can enable the one-wire interface on demand using raspi-config, or the following:

```
sudo modprobe w1-gpio
```

Newer kernels (4.9.28 and later) allow you to use dynamic overlay loading instead, including creating multiple 1-Wire busses to be used at the same time:

```
sudo dtoverlay w1-gpio gpiopin=4 pullup=0  # header pin 7
sudo dtoverlay w1-gpio gpiopin=17 pullup=0 # header pin 11
sudo dtoverlay w1-gpio gpiopin=27 pullup=0 # header pin 13
```

once any of the steps above have been performed, and discovery is complete you can list the devices that your Raspberry Pi has discovered via all 1-Wire busses (by default BCM4), like so:

```
ls /sys/bus/w1/devices/
```

n.b. Using w1-gpio on the Raspberry Pi typically needs a 4.7 kΩ pull-up resistor connected between the GPIO pin and a 3.3v supply (e.g. header pin 1 or 17). Other means of connecting 1-Wire devices to the Raspberry Pi are also possible, such as using i2c to 1-Wire bridge chips.


## Running ansible setup

```bash
ansible-playbook -k -i hosts ansible.yml
```