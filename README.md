# Request Echoing Server

## Overview

This repo and Docker image provides a test image that listens on :9080 and :9443 and echoes request headers, client certificates, and content presented to it.

To start, run like this:
```
docker run -p 9080:9080 -p 9443:9443 -ti liggitt/echo
```

# Docker image setup

### Build the Docker image from source

```
make build
```

### Run the Docker image from source


```
make run
```
  
## Example Use

```
curl -k https://localhost:9443/test --cert ./client.crt --key ./client.key
curl -k https://localhost:9080/test -H "Custom-Header: value" -d data
```
