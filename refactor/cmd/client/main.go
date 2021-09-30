package main

import (
	"flag"
	"os"
	"os/signal"
	"pw-ws-test/refactor/client"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"
)

var addr = flag.String("addr", "localhost:8100", "http service address")

func main() {
	flag.Parse()

	client, err := client.NewWebSocketClient(*addr, "frontend")
	if err != nil {
		panic(err)
	}
	logrus.Info("Connecting")

	go func() {
		// write down data every 100 ms
		ticker := time.NewTicker(time.Millisecond * 1500)
		i := 0
		for range ticker.C {
			err := client.Write(i)
			if err != nil {
				logrus.Errorf("error: %s, writing error", err.Error())
			}
			i++
		}
	}()

	// Close connection correctly on exit
	sigs := make(chan os.Signal, 1)

	// `signal.Notify` registers the given channel to
	// receive notifications of the specified signals.
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// The program will wait here until it gets the
	<-sigs
	client.Stop()
	logrus.Info("Goodbye")
}
