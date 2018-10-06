#!/bin/bash
set -e

cat > config/config_dev.json << EndOfConfig
{
  "mailer": {
    "smtp": {
      "addr": "$SMTP_HOST",
      "auth": {
        "username": "$SMTP_USERNAME",
        "password": "$SMTP_PASSWORD",
        "identity": "$SMTP_IDENTITY"
      }
    }
  },
  "translate": {
    "langs": "en, pl",
    "default": "en"
  },
  "router": {
    "host": ":80",
    "static": {
      "path": "./web/dist/",
      "prefix": "/static/"
    },
    "security": {
      "mode": "HTTP",
      "cert": "/certs/fullchain.pem",
      "key": "/certs/privkey.pem"
    }
  },
  "database": {
    "engine": "sqlite3",
    "url": "file:main.db?cache=shared"
  },
  "app": {
    "baseURL": "$APP_BASE_URL"
  },
  "oauth": {
    "github": {
      "app": "$OAUTH_GITHUB_APP",
      "secret": "$OAUTH_GITHUB_SECRET"
    }
  }
}
EndOfConfig

cp config/config_dev.json config/config_prod.json
cp config/config_dev.json config/config_test.json
