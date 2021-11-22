package cmd

import (
	"errors"
	"fmt"
	"os/exec"

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
	fmt.Println(arg)
	exec.Command("go", "mod", "init", arg)
	return nil
}


func init() {
	rootCmd.AddCommand(modCmd)
}
