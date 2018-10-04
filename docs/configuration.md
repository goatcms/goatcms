
# Configuration
Configuration usually involves different application parts (such as infrastructure and security credentials) and different environments (development, production). That's why we recommends that you split the application configuration into three parts.

## Where are my configs files
Application contains all configuration files in **/config** directory.
The goat systems storages config files in [JSON](http://www.json.org/) files.

## Example configuration
We have a separated files for dev, test and prod environments:
* config_dev.json
* config_test.json
* config_prod.json

They all looks like:
```json
{
  "mailer": {
    "smtp": {
      "addr": "smtp.gmail.com:465",
      "auth": {
        "username": "SECRET",
        "password": "SECRET",
        "identity": "SECRET"
      }
    }
  },
  "translate": {
    "langs": "en, pl",
    "default": "en"
  },
  "router": {
    "host": ":3937",
    "security": {
      "mode": "TLS",
      "cert": "./data/certs/fullchain.pem",
      "key": "./data/certs/privkey.pem"
    }
  },
  "database": {
    "engine": "sqlite3",
    "url": "file:data/main.db"
  },
  "app": {
    "baseURL": "URL_FOR_YOUR_APP"
  },
  "oauth": {
    "github": {
      "app": "SECRET",
      "secret": "SECRET"
    }
  }
}
```

### Outgoing mail configuration
**mailer.smtp** key contains your [SMTP (Simple Mail Transfer Protocol)](https://en.wikipedia.org/wiki/Simple_Mail_Transfer_Protocol) config.
* **mailer.smtp.addr** contains SMTP server address. default: *smtp.gmail.com:465*
* **mailer.smtp.auth.username** contains your username.
* **mailer.smtp.auth.password** contains your password.
* **mailer.smtp.auth.identity** contains server identity. Can be empty.

### Translates
**translate** key contains your translates configuration.
* **translate.langs** contains list of supported (allowed) languages.
* **translate.default** is default language. It is used when neither language is not match.

### Routing
**router** key contains your routing configuration.
* **router.host** is a host to bind.
* **router.security.mode** is default language. It is used when neither language is not match. You cam use ["TLS"](https://en.wikipedia.org/wiki/Transport_Layer_Security) or "HTTP" (unsecure).
* **router.security.cert** is path to fullchain certificate file on local disk. fullchain.pem contains cert.pem and chain.pem files.
* **router.security.key** is path to your private key pem file on local disk. Never share it.

### Database
*database* key contains key for local database settings.

* **database.engine** is name of your database engine. You can use "sqlite3" or "postgres"
* **database.url** contain string for database. It is "file:data/main.db" / , for sqlite3
  * ":memory:" for sqlite  engine mean create database in memory. It is use for tests / debugging. It is lost after close app.
  * "file:data/main.db" for sqlite engine mean create database in file. (It can be used for single instance app only).
  * "postgres://username:password@host/dbname?sslmode=disable" for postgres engine mean connect to exist database (with username and password).

### Application
*app* key contains application config.
* **baseURL** is full URL to application. It is used for web callback like oauth callback etc.

### OAuth
*oauth.app_name.** key contains secrets for [OAuth](https://oauth.net/2/) remote applications.

Example for singin in with gothub:
* **oauth.github.app** is application id (created for our own application).
* **oauth.github.app** is application secrets (created for our own application).
[There is tutorial how to register a new GitHub OAuth application](https://developer.github.com/apps/building-oauth-apps/creating-an-oauth-app/)
