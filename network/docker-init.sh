#!/bin/bash

SCRIPTDIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
source $SCRIPTDIR/prep.sh

# Stop if it is already running
if pgrep -x "$BINARY" >/dev/null; then
    echo "Terminating $BINARY..."
    killall $BINARY
fi

$SCRIPTDIR/init-chain.sh test-1 chain-test-1 16656 16657 1316 8080 "clock post desk civil pottery foster expand merit dash seminar song memory figure uniform spice circle try happy obvious trash crime hybrid hood cushion" "alley afraid soup fall idea toss can goose become valve initial strong forward bright dish figure check leopard decide warfare hub unusual join cart"
$SCRIPTDIR/init-chain.sh test-2 chain-test-2 26656 26657 1317 8081 "angry twist harsh drastic left brass behave host shove marriage fall update business leg direct reward object ugly security warm tuna model broccoli choice" "record gift you once hip style during joke field prize dust unique length more pencil transfer quit train device arrive energy sort steak upset"
$SCRIPTDIR/init-chain.sh test-3 chain-test-3 36656 36657 3316 8083 "girl absent inform acoustic across borrow pole prison alert cheap love disease fence actor candy february exclude sense iron maximum clinic hill interest bachelor" "surround hip section canvas rocket misery vast mom stereo renew file arctic draw very feel say surge wide crucial parrot jacket wreck focus decade"
