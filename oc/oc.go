package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/zhouxiaoxiang/oriole-oc/handle"
)

const (
	length = 3
	usage  = "Error: only used in oriole-service."
)

func main() {
	flag.Parse()
	info := flag.Args()

	if len(info) == length {
		handle.Handle(info...)
	} else {
		fmt.Println(usage)
		os.Exit(0)
	}
}
