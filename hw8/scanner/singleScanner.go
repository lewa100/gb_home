package scanner

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
	log "github.com/sirupsen/logrus"
)

// SingleScanner -scanner http with one flow.
func SingleScanner() {
	start := time.Now()
	for _, url := range os.Args[1:] {
		fetch(url)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

// fetch - single scanner.
func fetch(url string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	//bytes, err := ioutil.ReadAll(resp.Body)
	nbytes, err := io.Copy(ioutil.Discard, resp.Body) // to devNull
	resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%.2fs %7d %s\n", time.Since(start).Seconds(), nbytes, url)
}
