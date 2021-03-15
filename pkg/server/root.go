package server

import (
	"net/http"

	"github.com/markbates/pkger"
	"github.com/pyroscope-io/pyroscope/pkg/build"
)

//fs := http.FileServer(dir)
//mux.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
//	if r.URL.Path == "/" {
//		ctrl.statsInc("index")
//		ctrl.renderIndexPage(dir, rw, r)
//	} else if r.URL.Path == "/comparison" {
//		ctrl.statsInc("index")
//		ctrl.renderIndexPage(dir, rw, r)
//	} else {
//		fs.ServeHTTP(rw, r)
//	}
//})

func (ctrl *Controller) rootHandler(rw http.ResponseWriter, r *http.Request) {
	var dir http.FileSystem
	if build.UseEmbeddedAssets {
		// for this to work you need to run `pkger` first. See Makefile for more information
		dir = pkger.Dir("/webapp/public")
	} else {
		dir = http.Dir("./webapp/public")
	}
	fs := http.FileServer(dir)

	if r.URL.Path == "/" {
		ctrl.statsInc("index")
		ctrl.renderIndexPage(dir, rw, r)
	} else if r.URL.Path == "/comparison" {
		ctrl.statsInc("index")
		ctrl.renderIndexPage(dir, rw, r)
	} else {
		fs.ServeHTTP(rw, r)
	}
}
