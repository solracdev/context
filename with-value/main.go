package main

import (
	"context"
	"fmt"
)

func main() {
	ctxTest := context.WithValue(context.TODO(), "test", "value")
	ctxTest = context.WithValue(ctxTest, "test2", "value2")
	ctxTest = context.WithValue(ctxTest, "test3", "value3")

	ctxKey := fmt.Sprintf("%v", ctxTest)
	fmt.Println(ctxKey)
}
