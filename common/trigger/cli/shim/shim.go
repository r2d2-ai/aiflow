package main

import (
	"fmt"
	"os"

	"github.com/r2d2-ai/AIflow/common/trigger/cli"
)

func main() {

	result, err := cli.Invoke()

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
		os.Exit(1)
	}

	fmt.Fprintf(os.Stdout, "%s", result)
	os.Exit(0)
}
