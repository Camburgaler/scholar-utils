package main

import (
	"fmt"
	"os"

	emevdParser "github.com/Camburgaler/scholar-utils/pkg/data/emevd/parse"
	paramParser "github.com/Camburgaler/scholar-utils/pkg/data/param/parse"
	"github.com/Camburgaler/scholar-utils/pkg/output"
	"github.com/Camburgaler/scholar-utils/pkg/transform"
)

func run() error {
	paramData, err := paramParser.Parse()
	if err != nil {
		return err
	}

	emevdData, err := emevdParser.Parse()
	if err != nil {
		return err
	}

	result, err := transform.Transform(paramData, emevdData)
	if err != nil {
		return err
	}

	return output.Output(result, paramData, emevdData)
}

func main() {
	fmt.Println("Hello, Scholar!")

	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
