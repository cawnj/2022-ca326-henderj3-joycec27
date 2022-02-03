#!/usr/bin/env python3

"""
This basic program allows us to read or write to an NTAG215 card with a PN532 NFC module
"""

import board
import busio
from digitalio import DigitalInOut

from adafruit_pn532.i2c import PN532_I2C, BusyError


i2c = busio.I2C(board.SCL, board.SDA)
reset_pin = DigitalInOut(board.D20)
req_pin = DigitalInOut(board.D16)
pn532 = PN532_I2C(i2c, debug=False, reset=reset_pin, req=req_pin)

_, ver, rev, _ = pn532.firmware_version
print("Found PN532 with firmware version: {0}.{1}".format(ver, rev))
pn532.SAM_configuration()


def detect():
    """
    Wait for card detection and return its uid
    """
    print("Waiting for RFID/NFC card to write to...")
    uid = None
    while not uid:
        uid = pn532.read_passive_target(timeout=0.5)
    print(f"Found card with UID: {uid.hex()}\n")
    return uid

def read(block_number):
    """
    Abstraction layer to read a given block_number from the card
    """
    data = pn532.ntag2xx_read_block(block_number)
    return int.from_bytes(data, 'big')

def write(block_number, data):
    """
    Abstraction layer to write a given 4-byte array to a given block_number on the card
    """
    if block_number < 4 or block_number > 129:
        print("Sorry, we cannot write here, please stay within the range of 4-129.")
        return False
    elif block_number == 10:
        print("Sorry, we cannot write here, for some reason 10 doesn't work??")
        return False
    try:
        pn532.ntag2xx_write_block(block_number, data.to_bytes(4, 'big'))
    except OverflowError:
        print("Error: max 4-byte number is 4,294,967,295")
    return read(block_number) == data


def main():
    detect()
    while True:
        choice = input("Input r/w to read or write, or q to quit: ")
        if choice not in "rwq":
            print("Invalid choice, please try again.")
            continue
        if choice == "q":
            print("Goodbye!")
            break

        block_number = int(input("Input desired block number: "))
        if choice == "r":
            data = read(block_number)
            print(data)
        if choice == "w":
            data = int(input("Input desired number to write: "))
            if write(block_number, data):
                print("Success!")
            else:
                print("Failed.")

        print()

if __name__ == "__main__":
    main()
