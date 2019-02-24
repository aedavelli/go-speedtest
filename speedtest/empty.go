package speedtest

import (
	"io"
	"io/ioutil"
	"net/http"
)

func emptyHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		io.Copy(ioutil.Discard, r.Body)
	}
	w.Header().Add("Cache-Control", "no-store, no-cache, must-revalidate, max-age=0")
	w.Header().Add("Cache-Control", "post-check=0, pre-check=0")
	w.Header().Add("Pragma", "no-cache")
	w.Header().Add("Connection", "keep-alive")
	w.Header().Add("Content-Type", "text/html; charset=UTF-8")
	w.Header().Add("Keep-Alive", "timeout=5, max=99")
}

func init() {
	http.HandleFunc("/empty", emptyHandler)
}
