#!/bin/bash
set -e

## prepare config
/bin/bash /app/docker/config.sh

### run app
./webslots run --loglvl=$LOGLVL
