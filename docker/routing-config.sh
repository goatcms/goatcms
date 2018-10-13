#!/bin/bash
set -e

### Prepare routing configs
cat > /app/config/routing.json << EndOfMessage
{"static":[{
  "prefix": "/static/",
  "path": "./web/dist/"
}]}
EndOfMessage
