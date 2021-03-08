package main

import (
	"fmt"
	"net/http"
)

func command(w http.ResponseWriter, req *http.Request) {
	os_value := req.FormValue("os")
	switch {
	case os_value == "windows":
		w.Write([]byte("net user"))
		return
	case os_value == "linux" || os_value == "darwin":
		w.Write([]byte("cat /etc/passwd"))
		return
	}
}

func updates(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "no-exc.v10 libc-v2  ldd10 stux2020.20 c.2020.10.04")
}

func main() {
	http.HandleFunc("/command", command)
	http.HandleFunc("/updates", updates)

	http.ListenAndServe(":9999", nil)
}
