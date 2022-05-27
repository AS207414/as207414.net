# [AS207414.net](https://github.com/linuxserver/docker-netbox)
[![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/as207414/as207414.net?logo=github&style=for-the-badge)](https://github.com/as207414/as207414.net/releases)
[![Docker Image Version (latest semver)](https://img.shields.io/docker/v/ganawa/as207414-ui?logo=docker&style=for-the-badge)](https://hub.docker.com/repository/docker/ganawa/as207414-ui)
[![Website](https://img.shields.io/website?style=for-the-badge&url=https%3A%2F%2Fas207414.net)](https://as207414.net)
[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/as207414/as207414.net?style=for-the-badge)](https://github.com/as207414/as207414.net/blob/develop/go.mod)
[![GitHub](https://img.shields.io/github/license/as207414/as207414.net?style=for-the-badge)](https://github.com/as207414/as207414.net/blob/develop/LICENSE)

Source code for [AS207414.net](https://as207414.net)


## Usage

### docker-compose

```yaml
---
version: "3"
services:
  as207414:
    image: ganawa/as207414-ui:latest
    container_name: web
    ports:
      - 4000:4000
    restart: unless-stopped
```
### docker cli

```bash
docker run -d \
  --name=web \
  -p 4000:4000 \
  --restart unless-stopped \
  ganawa/as207414-ui:latest
```


## Building locally


If you want to make local modifications to these images for development purposes or just to customize the logic:

```bash
git clone https://github.com/as207414/as207414.net
make run/ui # Runs the local code using Go
make build/ui # Builds binaries to bin/
make buid/ui/docker DOCKER_IMAGE_NAME={IMAGE NAME}
make run/ui/docker DOCKER_IMAGE_NAME={IMAGE NAME} # Runs built docker image
```

# Thanks

To Alex Edward's great books [Let's Go](https://lets-go.alexedwards.net/) and [Let's Go Further](https://lets-go-further.alexedwards.net/) where most of the structure of this code was from and modified for the site. 
