package main

import (
	"log"
	"net/http"

	_ "github.com/aedavelli/go-speedtest/speedtest"
)

func main() {
	log.Fatal(http.ListenAndServe(":8081", nil))
}
