#!/usr/bin/env python3

import nfc
import ndef
import requests
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

def read_tag(tag):
    records = tag.ndef.records
    if len(records) == 0:
        print("No records found\n")
        return None
    return records[0].text

def main():
    with nfc.ContactlessFrontend('tty:S0:pn532') as clf:
        while True:
            tag = find_tag(clf)
            if tag.ndef:
                data = read_tag(tag)
                url = "https://sonic.cawnj.dev/entrylog"
                response = requests.post(url, json={"data": data})
                print(response.json())
                break

if __name__ == "__main__":
    main()
