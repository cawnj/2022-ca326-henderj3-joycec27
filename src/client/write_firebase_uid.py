#!/usr/bin/env python3

import nfc
import ndef
from nfc.clf import RemoteTarget
from time import sleep

import firebase_admin
from firebase_admin import credentials, auth
cred = credentials.Certificate("serviceAccountKey.json")
firebase_admin.initialize_app(cred)


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
    if tag.format():
        return True
    else:
        return False

def write_tag(tag, uid):
    record = ndef.TextRecord(uid)
    try:
        tag.ndef.records = [record]
    except:
        return False
    return True

def get_firebase_uid(email):
    try:
        user = auth.get_user_by_email(email)
        uid = user.uid
    except:
        print(f'Error finding account "{email}", please try again\n')
        return None
    print(f'Successfully found user "{email}"')
    print(f'UID: "{uid}"')
    return uid


def main():
    uid = None
    while not uid:
        email = input("Enter user's firebase email: ")
        uid = get_firebase_uid(email)

    with nfc.ContactlessFrontend('tty:S0:pn532') as clf:
        read = True
        while read:
            tag = find_tag(clf)
            if tag and tag.ndef: read = False

        if not format_tag(tag):
            print("Tag format failure, exiting...")
        elif not write_tag(tag, uid):
            print("Tag write failure, exiting...")
        else:
            print("UID written to tag successfully! Exiting...")

if __name__ == "__main__":
    main()
