package main

import (
	"golang.org/x/net/webdav"
	"net/http"
)

func NewWebDavServer(root string) http.Handler {
	return &WebDavHandler{
		fs: http.FileServer(http.Dir(root)),
		dav: &webdav.Handler{
			Prefix:     "",
			FileSystem: webdav.Dir(root),
			LockSystem: webdav.NewMemLS(),
			Logger:     nil,
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

func main() {
	http.ListenAndServe(":8090", NewWebDavServer("/tmp/"))
}
