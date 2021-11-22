package cmd

import (
	"fmt"
	"os"
)

func CurrentDir() string {
	c, _ := os.Getwd()
	return c
}

func ExitError(msg interface{}) {
	fmt.Fprintln(os.Stderr, msg)
	os.Exit(1)
}
