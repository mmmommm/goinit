package cmd

import (
	"errors"
	"os/exec"

	"github.com/spf13/cobra"
)

// worldCmd represents the world command
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
		runGoMod(args[0])
		return nil
	},
}

func runGoMod(arg string) error {
	exec.Command("go", "mod", "init", arg)
	return nil
}


func init() {
	rootCmd.AddCommand(modCmd)
}
