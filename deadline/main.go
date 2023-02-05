package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	deadLine := time.Date(2023, time.February, 5, 10, 30, 45, 0, time.UTC)
	ctx, cancel := context.WithDeadline(context.TODO(), deadLine)
	defer cancel()
	<-ctx.Done()
	fmt.Println("deadline is here!")
}
