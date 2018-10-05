#!/bin/bash
set -e

cat > .goat/secrets.json << EndOfMessage
{
  "database": {
    "engine": "sqlite",
    "host": "",
    "name": "",
    "password": "",
    "username": ""
  },
  "smtp": {
    "address": "",
    "identity": "",
    "password": "",
    "user": ""
  },
  "oauth": {
    "github": {
      "app": "",
      "secret": ""
    }
  }
}
EndOfMessage
