package main

import (
	"fmt"
	"os"

	"kcl-lang.io/krm-kcl/pkg/kio"
)

func main() {
	p := kio.NewPipeline(os.Stdin, os.Stdout, true)
	if err := p.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
