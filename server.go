package main

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	handler := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		buffer := &bytes.Buffer{}
		fmt.Fprintf(buffer, "req.Host: %q\n", req.Host)
		fmt.Fprintf(buffer, "req.URL: %q\n", req.URL.String())
		fmt.Fprintf(buffer, "req.Header: %#v\n", req.Header)
		if req.TLS != nil {
			for i, cert := range req.TLS.PeerCertificates {
				fmt.Fprintf(buffer, "req.TLS.PeerCertificates[%d]: %#v\n", i, cert.Subject)
			}
		}

		content, _ := ioutil.ReadAll(req.Body)
		if len(content) > 0 {
			fmt.Fprintf(buffer, "req.Body: %s\n", string(content))
		}

		w.Write(buffer.Bytes())
		fmt.Print(buffer.String())
	})

	go func() {
		server := &http.Server{
			Addr:      "0.0.0.0:9443",
			Handler:   handler,
			TLSConfig: &tls.Config{ClientAuth: tls.RequestClientCert},
		}
		fmt.Println("Listening on 0.0.0.0:9443...")
		fmt.Println(server.ListenAndServeTLS("localhost.crt", "localhost.key"))
	}()

	go func() {
		server := &http.Server{
			Addr:    "0.0.0.0:9080",
			Handler: handler,
		}
		fmt.Println("Listening on 0.0.0.0:9080...")
		fmt.Println(server.ListenAndServe())
	}()

	select {}
}
