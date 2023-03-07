package main

import (
	"flag"
)

func main() {
	var privileged bool

	flag.BoolVar(&privileged, "privileged", false, "Dummy flag to set application to run privileged")
}
