package speedtest

import (
	"crypto/rand"
	"fmt"
	"net/http"
	"strconv"
)

func init() {
	http.HandleFunc("/garbage", garbageHandler)
}

func garbageHandler(w http.ResponseWriter, r *http.Request) {
	// Download follows...
	w.Header().Add("Content-Description", "File Transfer")
	w.Header().Add("Content-Type", "application/octet-stream")
	w.Header().Add("Content-Disposition", "attachment; filename=random.dat")
	w.Header().Add("Content-Transfer-Encoding", "binary")
	// Never cache me
	w.Header().Add("Cache-Control", "no-store, no-cache, must-revalidate, max-age=0")
	w.Header().Add("Cache-Control", "post-check=0, pre-check=0")
	w.Header().Add("Pragma", "no-cache")
	d := make([]byte, 1048576)
	_, err := rand.Read(d)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	cs := 4
	if r.Form.Get("ckSize") != "" {
		cs, _ = strconv.Atoi(r.Form.Get("ckSize"))
	}

	if cs > 100 {
		cs = 100
	}

	for i := 0; i < cs; i++ {
		w.Write(d)
	}
}
