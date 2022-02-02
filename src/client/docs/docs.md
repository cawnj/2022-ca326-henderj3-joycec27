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

### NTAG215 Card
Datasheet [here](https://www.nxp.com/docs/en/data-sheet/NTAG213_215_216.pdf)
Adafruit library [here](https://github.com/adafruit/Adafruit_CircuitPython_PN532)
- Pages/Blocks 0-3,130-134 are reserved
- Reset Pin = D20; Request Pin = D16;
- Where data is a bytearray(4) or x.to_bytes(4, 'big'):
  - `pn532.ntag2xx_write_block(block_num, data)`
