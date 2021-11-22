package cmd

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
)

var modCmd = &cobra.Command{
	Use:   "p",
	Short: "initialize package name",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires package name")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := runGoMod(args[0]); err != nil {
			return err
		}
		return nil
	},
}

func runGoMod(arg string) error {
	p := filepath.Join(os.Getenv("GOPATH"), "pkg", "mod", arg)
	cmd := exec.Command("go", "mod", "init", p)
	_, err := cmd.Output()
	if err != nil {
		fmt.Print(err)
	}
	return nil
}


func init() {
	rootCmd.AddCommand(modCmd)
}
