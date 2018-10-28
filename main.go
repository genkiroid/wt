package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"time"
)

var content = `<html>
<head>
<title>Waited %s</title>
</head>
<body>
Hello world!<br/>
Waited %s.<br/>
</body>
</html>
`

func main() {
	var port int
	var tls bool
	var cert, key string

	flag.IntVar(&port, "p", 12345, "Port number.")
	flag.BoolVar(&tls, "tls", false, "Serve TLS.")
	flag.StringVar(&cert, "cert", "cert.pem", "Certificate file path.")
	flag.StringVar(&key, "key", "key.pem", "Key file path.")

	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		wt := q.Get("wt")
		d, err := time.ParseDuration(wt)
		if err != nil {
			fmt.Fprintln(w, err)
			return
		}

		time.Sleep(d)

		fmt.Fprintf(w, content, d, d)
	})

	if tls {
		certFile, err := filepath.Abs(cert)
		if err != nil {
			panic(err)
		}
		keyFile, err := filepath.Abs(key)
		if err != nil {
			panic(err)
		}
		log.Fatal(http.ListenAndServeTLS(":"+strconv.Itoa(port), certFile, keyFile, nil))
	} else {
		log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), nil))
	}
}
