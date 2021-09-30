package main

import (
	"flag"
	"net/http"
	"pw-ws-test/refactor/server"

	"github.com/sirupsen/logrus"
)

var addr = flag.String("addr", "localhost:8000", "http service address")

func serveHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "home.html")
}

func main() {
	flag.Parse()

	hub := server.NewHub()
	go hub.Run()
	http.HandleFunc("/frontend", func(w http.ResponseWriter, r *http.Request) {
		logrus.Info("got new connection")
		server.ServeWs(hub, w, r)
	})

	logrus.Infof("server started ... %s", *addr)
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		panic(err)
	}

}
