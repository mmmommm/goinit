package cmd

import (
	"embed"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "goinit",
}

func Execute() {
	rootCmd.Run = func(cmd *cobra.Command, args []string) {
		if err := CreateFiles(); err != nil {
			ExitError(err)
		}
	}
	if err := rootCmd.Execute(); err != nil {
		ExitError(err)
	}
}

//go:embed files/*
var local embed.FS

// create .gitignore, .golangci.yml, LICENSE, README.md, main.go, test.yml, lint.yml
func CreateFiles() error {
	fis, err := local.ReadDir("files")
	if err != nil {
		return err
	}
	for _, fi := range fis {
		in, err := local.Open(filepath.Join("files", fi.Name()))
		if err != nil {
			return err
		}
		out, err := os.Create(filepath.Join(CurrentDir(), filepath.Base(fi.Name())))
		if err != nil {
			return err
		}
		if _, err := io.Copy(out, in); err != nil {
			return err
		}
		out.Close()
		in.Close()
		log.Println("exported", filepath.Base(fi.Name()))
	}
	//move test.yml and lint.yml
	actionsPath := filepath.Join(CurrentDir(), ".github", "workflows")
	// i don't know chmod 0777 is best for this usage.
	if err := os.MkdirAll(filepath.Join(actionsPath), 0777); err != nil {
		return err
	}
	if err := os.Rename(filepath.Join(CurrentDir(), "test.yml"), filepath.Join(actionsPath, "test.yml")); err != nil {
		return err
	}
	if err := os.Rename(filepath.Join(CurrentDir(), "lint.yml"), filepath.Join(actionsPath, "lint.yml")); err != nil {
		return err
	}
	return nil
}
