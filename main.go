package main

import (
	"fmt"
	"learn-go/greeting"
)

func main() {
	message := greeting.Greet("Gopher")
	fmt.Println(message)
}
