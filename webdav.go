package main

import (
	"fmt"
	"golang.org/x/net/webdav"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "hello\n")
}

func main() {
	handler := &webdav.Handler{
		Prefix:     "",
		FileSystem: webdav.Dir("/tmp/"),
		LockSystem: webdav.NewMemLS(),
		Logger:     nil,
	}
	http.ListenAndServe(":8090", handler)
}
