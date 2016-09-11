// httpRd.go
package main

import (
	"bytes"
	"net/http"
	"os/exec"
)

func GetIP() (string, error) {
	cmd := exec.Command("/bin/bash", "-c", `/sbin/ifconfig -a|grep inet|grep -v 127.0.0.1|grep -v inet6|awk '{print $2}'|tr -d "addr:"`)
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Start()
	if err != nil {
		return "", err
	}

	err = cmd.Wait()
	if err != nil {
		return "", err
	}

	return out.String(), nil
}

func redirect(w http.ResponseWriter, r *http.Request) {
	localhost, err := GetIP()
	if err != nil {
		w.WriteHeader(200)
		w.Write([]byte(err.Error()))
		return
	} else {
		http.Redirect(w, r, "https://"+localhost, 302)
	}
}

func main() {
	http.ListenAndServe(":80", http.HandlerFunc(redirect))
}
