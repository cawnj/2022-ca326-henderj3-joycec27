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
- **NOTE**: This setup is not currently in use, but is some useful info to keep here still

### NTAG215 Card
Datasheet [here](https://www.nxp.com/docs/en/data-sheet/NTAG213_215_216.pdf)
Adafruit library [here](https://github.com/adafruit/Adafruit_CircuitPython_PN532)
- Pages/Blocks 0-3,130-134 are reserved
- Reset Pin = D20; Request Pin = D16;
- Where data is a bytearray(4) or x.to_bytes(4, 'big'):
  - `pn532.ntag2xx_write_block(block_num, data)`

### nfcpy
- `pip install nfcpy`
- PN532 set to UART (this is a serial input at /dev/ttyS0)
- Testing connection:
  ```
  clf = nfc.ContactlessFrontend()
  clf.open('tty:S0:pn532')
  clf.close()
  ```

### ndeflib
- `pip install ndeflib`
- Can use this tool [here](https://github.com/nfcpy/nfcpy/blob/master/examples/tagtool.py) to dump, load, write, format
- `python3 tagtool.py --device tty:S0:pn532 <command>`
- Useful usage stuff:
  ```
  clf = nfc.ContactlessFrontend('tty:S0:pn532')
  target = clf.sense(RemoteTarget('106A'))
  tag = nfc.tag.activate(clf, target)
  tag.format()
  records = tag.ndef.records
  ndef.TextRecord("Hello world!")
  tag.ndef.records += [record]
  ```

### Expected POST request body
```
{
  "user_id": scanned_from_card,
  "location_id": constant_in_scanner_script,
  "timestamp": current_time
}
```
