package main

import (
	"fmt"
	"os"

	"github.com/Camburgaler/scholar-utils/pkg/output"
	"github.com/Camburgaler/scholar-utils/pkg/parse"
)

func run() error {
	data, err := parse.Parse()
	if err != nil {
		return err
	}

	// result := transform(data)

	return output.Output(data)
}

func main() {
	fmt.Println("Hello, Scholar!")

	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
