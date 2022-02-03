#!/usr/bin/env python3

import nfc
import ndef
from nfc.clf import RemoteTarget
from time import sleep


def find_tag(clf):
    print("\nWaiting for target...")
    while True:
        target = clf.sense(RemoteTarget('106A'))

        if target:
            break
        sleep(0.5)

    tag = nfc.tag.activate(clf, target)
    print(f"Found tag: {tag}\n")
    return tag

def format_tag(tag):
    if input("Would you like to format this card? y/n: ") != "y": return None

    if tag.format():
        print("Format success")
        return True
    else:
        print("Format failed")
        return False

def read_tag(tag):
    records = tag.ndef.records
    if len(records) == 0:
        print("No records found\n")
        return False

    print("Records:")
    for i, record in enumerate(records):
        print(f"\t{i}: {record.text}")
    print()
    return True

def write_tag(tag):
    if input("Would you like to write to this card? y/n: ") != "y": return None

    message = input("Please input desired message to write:\n")
    record = ndef.TextRecord(message)

    try:
        tag.ndef.records += [record]
    except:
        print("Write failed\n")
        return False
    print("Write success\n")
    return True


def main():
    with nfc.ContactlessFrontend('tty:S0:pn532') as clf:
        while True:
            tag = find_tag(clf)
            if tag.ndef:
                read_tag(tag)
                write_tag(tag)
            format_tag(tag)
            if input("Would you like to scan another card? y/n: ") == "n": break
    print("\nGoodbye!")

if __name__ == "__main__":
    main()
