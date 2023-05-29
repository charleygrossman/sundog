package main

import (
	"fmt"
	"marketdata/cmd/eod/internal"
)

func main() {
	if err := internal.RootCmd.Execute(); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
