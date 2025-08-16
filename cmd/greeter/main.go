package main

import (
	"fmt"

	"sync-leaf/pkg/greetings"
)

func run() string {
	return greetings.Hello()
}

func main() {
	fmt.Println(run())
}
