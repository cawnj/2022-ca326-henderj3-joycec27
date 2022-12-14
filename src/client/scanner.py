#!/usr/bin/env python3

import nfc
import ndef
import os
import requests
from nfc.clf import RemoteTarget
from time import sleep, gmtime, strftime

USER_ENDPOINT = "https://sonic.cawnj.dev/user"
ENTRYLOG_ENDPOINT = "https://sonic.cawnj.dev/entrylog"

LOCATION_ID = os.getenv("LOCATION_ID")
if not LOCATION_ID:
    print("Please pass LOCATION_ID enviornment variable")
    exit(1)

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
    user_id = read_tag(tag)
    print("Found user id:", user_id)
    return user_id

def send_entry_log_request(user_id):
    print("\nSending entry log request...")
    payload = {
        "user_id": user_id,
        "location_id": int(LOCATION_ID),
        "timestamp": strftime("%Y-%m-%d %H:%M:%S", gmtime()),
    }
    print("Request:", payload)
    response = requests.post(url=ENTRYLOG_ENDPOINT, json=payload)
    try:
        print("Response:", response.json())
    except:
        endpoint_error(ENTRYLOG_ENDPOINT)

def endpoint_error(endpoint):
    print("Something went wrong sending request to %s", endpoint)

def main():
    with nfc.ContactlessFrontend('tty:S0:pn532') as clf:
        while True:
            tag = find_tag(clf)
            if tag.ndef:
                user_id = get_user_id(tag)
                send_entry_log_request(user_id)
            sleep(2)

if __name__ == "__main__":
    main()
