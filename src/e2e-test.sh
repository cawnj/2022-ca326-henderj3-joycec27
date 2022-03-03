#!/bin/bash

BASE_URL="https://sonic.cawnj.dev"
DATE=$(date '+%Y-%m-%d')

MAIN_USER="sPxteAo65YhizCxCirkzjDkfE3w1"
MAIN_EXPOTOKEN="ExponentPushToken[HN5H4pLr5vHrlzTMZStAQD]"

TEST_USER="test"
TEST_EXPOTOKEN="test"

function send_post_req () {
    local ENDPOINT=$1
    local DATA=$2
    curl --location --request POST "$BASE_URL$ENDPOINT" \
    --header 'Content-Type: application/json' \
    --data-raw "$DATA"
}

function register () {
    local USER=$1
    local EXPOTOKEN=$2
    send_post_req "/register" "{\"user_id\":\"$USER\",\"expo_token\":\"$EXPOTOKEN\"}"
}

register $MAIN_USER $MAIN_EXPOTOKEN
register $TEST_USER $TEST_EXPOTOKEN

function entrylog () {
    local USER=$1
    local TIME=$2
    send_post_req "/entrylog" "{\"user_id\":\"$USER\",\"location_id\":2,\"timestamp\":\"$DATE $TIME\"}"
}

entrylog $MAIN_USER "16:00:00"
entrylog $TEST_USER "16:10:00"
entrylog $MAIN_USER "16:30:00"
entrylog $TEST_USER "16:40:00"

function trace () {
    local USER=$1
    send_post_req "/trace" "{\"user_id\":\"$USER\"}"
}

trace $TEST_USER
