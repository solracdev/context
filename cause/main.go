package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

var ErrFailure = fmt.Errorf("request took too")

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Millisecond)
	defer cancel()

	ctx, cancelCause := context.WithCancelCause(context.Background())
	cancelCause(ErrFailure)

	respCh := make(chan int, 1)
	go func() {
		res, err := http.Get("https://google.com")
		if err != nil {
			log.Fatal(err)
		}
		respCh <- res.StatusCode
	}()

	select {
	case <-ctx.Done():
		fmt.Printf("operation could not complete: %s", context.Cause(ctx))
	case code := <-respCh:
		fmt.Println("status code: ", code)
	}
}
