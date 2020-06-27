package scanner

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

// MultiScanner -scanner http with parallel flows.
func MultiScanner(){
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go mFetch(url, ch)
	}
	for range os.Args[1:] {
		fmt.Print(<-ch) // receive from channel
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

// mFetch - multi scanner.
func mFetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body) // to devNull
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	ch <- fmt.Sprintf("%.2fs %7d %s\n", time.Since(start).Seconds(), nbytes, url)
}
