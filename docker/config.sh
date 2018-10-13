#!/bin/bash
set -e

mkdir -p /app/config
mkdir -p /app/data

if [ $MODE = "HTTPS" ]; then
  sh -x /app/docker/https-config.sh
else
  sh -x /app/docker/http-config.sh
fi

### Build database
if [ ! -f main.db ]; then
  ./webslots db:build
fi

### Prepare routing files
/bin/bash -x /app/docker/routing-config.sh

### Add users
for i in `env | grep -E "^USER_.*_USERNAME="`
do
    baseKey="USER_$(echo $i| cut -d'_' -f 2)_"
    eval username=\${${baseKey}USERNAME}
    eval email=\${${baseKey}EMAIL}
    eval firstname=\${${baseKey}FIRSTNAME}
    eval lastname=\${${baseKey}LASTNAME}
    ./goatcms user:add --interactive=false --username=$username --email=$email --firstname=$firstname --lastname=$lastname
    echo "User $username ($email)... added "
done

### Update users roles
for i in `env | grep -E "^USER_.*_ROLES="`
do
    baseKey="USER_$(echo $i| cut -d'_' -f 2)_"
    eval username=\${${baseKey}USERNAME}
    eval email=\${${baseKey}EMAIL}
    eval roles=\${${baseKey}ROLES}
    ./goatcms user:roles:update --by=$email --roles=$roles
    echo "User $username ($email)... roles updated"
done

### Update users passwords
for i in `env | grep -E "^USER_.*_PASSWORD="`
do
    baseKey="USER_$(echo $i| cut -d'_' -f 2)_"
    eval username=\${${baseKey}USERNAME}
    eval email=\${${baseKey}EMAIL}
    eval password=\${${baseKey}PASSWORD}
    ./goatcms user:password:update --by=$username --password=$password
    echo "User $username ($email)... password updated"
done

### Connect user to remote accounts
for i in `env | grep -E "^USER_.*_CONNECT_.*="`
do
    userBaseKey="USER_$(echo $i| cut -d'_' -f 2)_"
    connectBaseKey="USER_$(echo $i| cut -d'_' -f 2)_CONNECT_$(echo $i| cut -d'_' -f 4)_"
    service="$(echo $i| cut -d'_' -f 4 | cut -d'=' -f 1 | tr '[:upper:]' '[:lower:]')"
    eval username=\${${userBaseKey}USERNAME}
    eval email=\${${userBaseKey}EMAIL}
    eval remote=\${${i}}
    ./goatcms user:connected:add --local=$email --remote=$remote --service=$service
    echo "Connect $username ($email) user to $service remote service... success"
done
