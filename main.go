package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	_ "github.com/aedavelli/go-speedtest/speedtest"
)

var listenPort int
var listenIP string

func init() {
	flag.IntVar(&listenPort, "p", 8080, "SpeedTest listen port")
	flag.StringVar(&listenIP, "i", "", "SpeedTest listen IP")
}

func main() {
	flag.Parse()
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", listenIP, listenPort), nil))
}
