#!/bin/sh -x
set -e

# Start the SDK
/bin/sdk &

# Start the game
/docker-entrypoint.sh