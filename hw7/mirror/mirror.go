package mirror

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

// MirroredQuery - for compare pings.
func MirroredQuery() string {
	responses := make(chan string, 3)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		responses <- request("https://asianmirror.lk/", &wg)
	}()
	wg.Add(1)
	go func() {
		responses <- request("https://ec.europa.eu/", &wg)
	}()
	wg.Add(1)
	go func() {
		responses <- request("https://www.nike.com/men", &wg)
	}()
	wg.Wait()
	return <-responses // возврат самого быстрого ответа
}

// request - for ping site.
func request(hostname string,wg * sync.WaitGroup) string {
	defer wg.Done()
	start := time.Now()
	resp, err := http.Get(hostname)
	if err != nil {
		log.Fatal(err)
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body) // to devNull
	resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%.2fs %7d %s\n", time.Since(start).Seconds(), nbytes, hostname)
	return fmt.Sprintf("%.2fs %7d %s\n", time.Since(start).Seconds(), nbytes, hostname)
}

