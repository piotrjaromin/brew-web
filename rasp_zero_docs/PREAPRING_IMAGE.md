# Preparing image

1. Download raspbian lite - https://downloads.raspberrypi.org/raspbian_lite_latest
2. Install balena etcher - https://www.balena.io/etcher/
3. Burn image to sd card with balenaEtcher
4. on sdcard create file `wpa_supplicant.conf` with contents:

    ctrl_interface=/var/run/wpa_supplicant
    update_config=1
    country=ISO_COUNTRY_CODE

    network={
        scan_ssid=1
        ssid="WIFI_SSID"
        psk="WIFI_PASSWORD"
    }

5. create empty `ssh` file on sd card root