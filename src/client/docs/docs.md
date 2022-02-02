# Sonic Client

## Documentation/Help

### PN532 NFC Hat Setup
Following guide [here](https://littlebirdelectronics.com.au/guides/181/nfc-module-with-raspberry-pi)
- Enable I2C via raspi-config
- Install dependencies: `libusb-dev libpcsclite-dev i2c-tools`
- Build and install libnfc:
  - Get latest release from `https://github.com/nfc-tools/libnfc/releases`
  - These commands:
    - `tar xf libnfc*.tar.bz2 && cd libnfc*`
    - `./configure --prefix=/usr --sysconfdir=/etc`
    - `make`
    - `sudo make install`
- Copy `libnfc.conf` to /etc/nfc/
- Set jumpers/switches on device, guide [here](https://www.waveshare.com/wiki/PN532_NFC_HAT)
