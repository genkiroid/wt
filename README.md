# wt

wt wates any time and return HTTP response.

## Installation

```sh
go get -u github.com/genkiroid/wt
```

## Start server

```sh
wt
```

Specify port number.(default 12345)

```sh
wt -p 8888
```

## Usage

URI is http://localhost:{port}/?wt={duration string}

### Example

```sh
# Wait 3 sec.
http://localhost:12345/?wt=3s

# Wait 1 minute 30 sec.
http://localhost:12345/?wt=1m30s
```

refs. https://golang.org/pkg/time/#ParseDuration
