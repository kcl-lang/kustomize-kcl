package main

import (
	"fmt"
	"os"

	"github.com/KusionStack/krm-kcl/pkg/kio"
)

func main() {
	p := kio.NewPipeline(os.Stdin, os.Stdout, true)
	if err := p.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
