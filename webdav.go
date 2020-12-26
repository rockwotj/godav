package main

import (
	"flag"
	"fmt"
	"golang.org/x/net/webdav"
	"log"
	"net/http"
)

func NewWebDavServer(root string) http.Handler {
	return &WebDavHandler{
		fs: http.FileServer(http.Dir(root)),
		dav: &webdav.Handler{
			Prefix:     "",
			FileSystem: webdav.Dir(root),
			LockSystem: webdav.NewMemLS(),
			Logger: func(r *http.Request, err error) {
				log.Println("Error serving", r.URL, "with error", err)
			},
		},
	}
}

type WebDavHandler struct {
	fs  http.Handler
	dav http.Handler
}

func (h *WebDavHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// NOTE: The builtin webdav server doesn't list directories, so we do that
	// by using a plain HTTP Server
	if r.Method == "GET" || r.Method == "HEAD" {
		h.fs.ServeHTTP(w, r)
	} else {
		h.dav.ServeHTTP(w, r)
	}
}

var port = flag.Int("port", 8090, "The port to run the server on")
var root = flag.String("root", "/tmp/", "The root directory for the server")

func main() {
	log.Println("Starting webdav server at", *root, "on port", *port)
	http.ListenAndServe(fmt.Sprintf(":%d", *port), NewWebDavServer(*root))
}
