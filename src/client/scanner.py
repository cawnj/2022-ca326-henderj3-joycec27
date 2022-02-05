#!/usr/bin/env python3

import nfc
import ndef
import requests
from nfc.clf import RemoteTarget
from time import sleep, gmtime, strftime


LOCATION_ID = 1
API_ENDPOINT = "https://sonic.cawnj.dev/entrylog"

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

def get_user_id(tag):
    return int(read_tag(tag))

def send_post_request(user_id):
    data = {
        "user_id": user_id,
        "location_id": LOCATION_ID,
        "entry_time": strftime("%Y-%m-%d %H:%M:%S", gmtime()),
        "exit_time": "3000-01-01 00:00:00"
    }
    print(data)
    response = requests.post(url=API_ENDPOINT, json=data)
    try:
        print(response.json())
    except:
        print("Something went wrong sending POST request")

def main():
    with nfc.ContactlessFrontend('tty:S0:pn532') as clf:
        while True:
            tag = find_tag(clf)
            if tag.ndef:
                user_id = get_user_id(tag)
                send_post_request(user_id)
                break

if __name__ == "__main__":
    main()
