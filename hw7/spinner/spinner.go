package spinner

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Spinner - Spinner work 10 seconds.
func Spinner() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var wg sync.WaitGroup
	wg.Add(1)
	go spinner(50*time.Millisecond, &wg, ctx)
	wg.Wait()
}

//spinner - goroutine.
func spinner(delay time.Duration, wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()
	for {
		for _, r := range "-\\|/" {
			fmt.Printf("%c\r", r)
			time.Sleep(delay)
		}
		select {
		case <-ctx.Done():
			return
		default:
		}
	}
}
