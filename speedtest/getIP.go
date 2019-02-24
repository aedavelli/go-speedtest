package speedtest

import (
	"io"
	"net/http"
	"strings"
)

func getIPHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, strings.Split(r.RemoteAddr, ":")[0])
}

func init() {
	http.HandleFunc("/getIP", getIPHandler)
}
