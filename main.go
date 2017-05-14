package main

import (
	"flag"
	"fmt"
)

var (
	outputDirectory string
)

func init() {
	flag.StringVar(&outputDirectory, "d", "", "output directory")

	flag.Parse()
}

func main() {
	fmt.Printf("outputDirectory: %s\n", outputDirectory)
}
