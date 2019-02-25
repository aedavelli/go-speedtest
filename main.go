package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/aedavelli/go-speedtest/speedtest"
)

var listenPort int
var listenIP string
var replay bool

func init() {
	flag.IntVar(&listenPort, "p", 8080, "SpeedTest listen port")
	flag.StringVar(&listenIP, "i", "", "SpeedTest listen IP")
	flag.BoolVar(&replay, "r", false, "Replay same random 1MB data")
}

func main() {
	flag.Parse()
	speedtest.SetReplay(replay)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", listenIP, listenPort), nil))
}
