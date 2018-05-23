#!/bin/sh -x
set -e

# Generate the RCON password
mkdir -p $CONFIG
if [ ! -f $CONFIG/rconpw ]; then
  echo $(pwgen 15 1) > $CONFIG/rconpw
fi

# Start the SDK
/bin/agones-factorio-sdk --password=$(cat $CONFIG/rconpw) &

# Start the game
/docker-entrypoint.sh