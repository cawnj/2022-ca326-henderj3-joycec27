#!/usr/bin/env python3

import nfc
import ndef
from nfc.clf import RemoteTarget
from time import sleep
from getpass import getpass

from firebase import Firebase
firebase_config = {
  "apiKey": "AIzaSyCHb34Qxrj_lh3XU9ARTnew_uImtDeuVso",
  "authDomain": "fir-auth-451b3.firebaseapp.com",
  "databaseURL": "https://fir-auth-451b3-default-rtdb.europe-west1.firebasedatabase.app",
  "storageBucket": "fir-auth-451b3.appspot.com",
}
firebase = Firebase(firebase_config)
auth = firebase.auth()


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

def get_firebase_uid(email, password):
    try:
        user = auth.sign_in_with_email_and_password(email, password)
        uid = user["localId"]
    except:
        print(f'Error logging in to account "{email}", please try again\n')
        return None
    print(f'Successfully logged in with user "{email}"')
    print(f'UID: "{uid}"')
    return uid


def main():
    uid = None
    print("Login with Firebase:\n")
    while not uid:
        email = input("Email: ")
        password = getpass("Password: ")
        uid = get_firebase_uid(email, password)

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
