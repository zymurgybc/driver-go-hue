package main

import (
	"fmt"
	"os"
	"os/signal"
)

func main() {

	NewHueDriver()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)

	// Block until a signal is received.
	s := <-c
	fmt.Println("Got signal:", s)

}
