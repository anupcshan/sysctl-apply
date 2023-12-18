package main

import (
	"flag"
	"log"

	"github.com/lorenzosaino/go-sysctl"
)

func main() {
	sysctlFile := flag.String("sysctl", "/etc/sysctl.conf", "sysctl file")
	flag.Parse()

	if err := sysctl.LoadConfigAndApply(*sysctlFile); err != nil {
		log.Fatal(err)
	}
}
