// Page 17
// This runs e1.7-e1.9. Each commit in version control is one exercise finished
// and may have slightly modified this file.
package main

import (
	"fmt"
	"github.com/ncruz8991/gopl.io/ch1/exercises/e1.8"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(e1_8.ParseUrl(url))
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		// Instead of using ioutil.ReadAll, this copies the bytes straight
		// to the standard output.
		_, err = io.Copy(os.Stdout, resp.Body)
		fmt.Printf("---\nReceived response with status: %s\n", resp.Status)

		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}
