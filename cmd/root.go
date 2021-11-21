package cmd

import (
	"bytes"
	"embed"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
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
		if err := runGoMod(); err != nil {
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
		if fi.Name() == "ci.yaml" {
			if err := createActions() {
				return err
			}
			return nil
		}
		in, err := local.Open(filepath.Join("files", fi.Name()))
		if err != nil {
			return err
		}
		out, err := os.Create(filepath.Join(CurrentDir(), "example", filepath.Base(fi.Name())))
		if err != nil {
			return err
		}
		_, err := io.Copy(out, in)
		if err != nil {
			return err
		}
		out.Close()
		in.Close()
		log.Println("exported", filepath.Base(fi.Name()))
	}
	return nil
}

// create .github/workflows/ci.yaml
func createActions() error {
	actionsPath := filepath.Join(CurrentDir(), "example", ".github", "workflows", "ci.yaml")
	file, err := local.ReadFile(filepath.Join("files", "ci.yaml"))
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Join(CurrentDir(), "example", ".github", "workflows"), 0777); err != nil {
		return err
	}
	out, err := os.Create(actionsPath)
	if err != nil {
		return err
	}
	_, err := io.Copy(out, bytes.NewReader(file))
	if err != nil {
		return err
	}
	out.Close()
	log.Println("exported .github/workflows/ci.yaml")
	return nil
}

func runGoMod() error {
	exec.Command("go", "mod", "init", os.Args[0])
	return nil
}
