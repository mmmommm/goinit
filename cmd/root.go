package cmd

import (
	"embed"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

func CurrentDir () string {
	c, _ := os.Getwd()
	return c
}

var rootCmd = &cobra.Command{
	Use: "goinit",
}

func exitError(msg interface{}) {
	fmt.Fprintln(os.Stderr, msg)
	os.Exit(1)
}

func Execute() {
	rootCmd.Run = func(cmd *cobra.Command, args []string) {
		if err := createFiles(); err != nil {
			exitError(err)
		}
	}
	if err := rootCmd.Execute(); err != nil {
		exitError(err)
	}
}

//go:embed files/*
var local embed.FS

// create .gitignore, LICENSE, README.md, main.go
func createFiles() error {
	fis, err := local.ReadDir("files")
	if err != nil {
		return err
	}
	for _, fi := range fis {
		fmt.Println(fi.Name())
		if fi.Name() == "lint.yml" {
			if err := createActions(fi.Name()); err != nil {
				return err
			}
			return nil
		}
		if fi.Name() == "test.yml" {
			if err := createActions(fi.Name()); err != nil {
				return err
			}
			return nil
		}
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
	return nil
}

// create .github/workflows/test.yaml, lint.yml
func createActions(name string) error {
	actionsPath := filepath.Join(CurrentDir(), ".github", "workflows")
	// i don't know chmod 0777 is best for this.
	if err := os.MkdirAll(filepath.Join(actionsPath), 0777); err != nil {
		return err
	}
	in, err := local.Open(filepath.Join("files", name))
	if err != nil {
		return err
	}
	out, err := os.Create(filepath.Join(actionsPath, filepath.Base(name)))
	if err != nil {
		return err
	}
	if _, err := io.Copy(out, in); err != nil {
		return err
	}
	out.Close()
	in.Close()
	output := filepath.Join(".github", "workflows", filepath.Base(name))
	log.Println("exported", output)

	return nil
}
