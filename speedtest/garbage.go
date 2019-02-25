package speedtest

import (
	"crypto/rand"
	"fmt"
	"net/http"
	"strconv"
)

var replayData []byte

func init() {
	http.HandleFunc("/garbage", garbageHandler)
}

func SetReplay(r bool) {
	if r {
		replayData = make([]byte, 1048576)
		_, err := rand.Read(replayData)
		if err != nil {
			fmt.Println("error:", err)
			return
		}
	}
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

	var d []byte
	if len(replayData) == 0 {
		d = make([]byte, 1048576)
		_, err := rand.Read(d)
		if err != nil {
			fmt.Println("error:", err)
			return
		}
	} else {
		d = replayData
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
