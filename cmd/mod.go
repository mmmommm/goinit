package cmd

import (
	"errors"
	"fmt"
	"log"
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
			ExitError(err)
		}
		if err := CreateFiles(); err != nil {
			ExitError(err)
		}
		return nil
	},
}

func runGoMod(arg string) error {
	cmd := exec.Command("go", "mod", "init", arg)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + string(output))
	}
	log.Println(string(output))
	return nil
}

func init() {
	rootCmd.AddCommand(modCmd)
}
