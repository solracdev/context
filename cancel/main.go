package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	go func() {
		ticker := time.NewTicker(5 * time.Second)
		<-ticker.C
		pid := os.Getpid()
		_ = syscall.Kill(pid, syscall.SIGTERM)
	}()

	ctx, cancel := context.WithCancel(context.TODO())
	go applicationEnd(cancel)
	<-ctx.Done()
	fmt.Printf("Application DONE by: %v \n", ctx.Err())
}

func applicationEnd(cf context.CancelFunc) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
	sig := <-sigChan
	fmt.Printf("context cancel form signal %v \n", sig)
	cf()
}
