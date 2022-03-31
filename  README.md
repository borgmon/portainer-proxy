# Docker Status Proxy
The fastest way to expose your container state with a public API

## Setup
```bash
git clone https://github.com/borgmon/state-proxy
docker-compose -f "./state-proxy/docker-compose.yml" up -d
```

Then you will have this service on `:8090/state/{YOUR_CONTAINER_NAME}`
```json
{
    "state": "running"
}
```

## Config
You can pass following ENV to config the service

ENV | Default | Note
---|---|---
WHITELIST | - | Pass a csv here. If empty: all containers can be reached. Ex: web1,web2
PORT | 8090 | Port number

## Volume
This service need to access to docker socket. Please mount via
```
/var/run/docker.sock:/var/run/docker.sock
```
