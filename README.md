# wt

wt wates any time and return HTTP/HTTPS response.

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

## Start TLS server

```sh
wt -tls
```

If error ocurred about cert file path and key file path, specify file path.
(Cert and key is contained this repo.)

```sh
wt -tls -cert /path/to/cert.pem -key /path/to/key.pem
```

## Usage

URI is `http://localhost:{port}/?wt={duration string}`

### Example

```sh
# Wait 3 sec.
http://localhost:12345/?wt=3s

# Wait 1 minute 30 sec.
http://localhost:12345/?wt=1m30s
```

refs. https://golang.org/pkg/time/#ParseDuration
