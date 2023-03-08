package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	var privileged bool

	flag.BoolVar(&privileged, "privileged", false, "Dummy flag to set application to run privileged")
	flag.Parse()

	if privileged {
		fmt.Println("Hello world privileged")
	} else {
		fmt.Println("Hello world")
	}
	time.Sleep(60)

}
