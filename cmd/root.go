package cmd

import (
	"embed"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/mmmommm/goinit/util"
	"github.com/spf13/cobra"
)

var (
	//go:embed files/*
	local embed.FS

	// Used for flags.
	module string
	v      = "0.1.3"

	rootCmd = &cobra.Command{
		Use:   "goinit",
		Short: "Initialize configuration generator for Go",
		Long:  `goinit is a cli tool to create initialize configuration file to begin go programing.`,
	}
)

func Execute() {
	rootCmd = &cobra.Command{
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("directory name is required")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if args[0] == "goinit" {
				return errors.New("don't use \"goinit\" for directory name, please use another one")
			}

			err := run(cmd, args)
			return err
		},
		Version:          v,
		TraverseChildren: true,
	}
	rootCmd.PersistentFlags().StringVarP(&module, "module", "m", "", "-m")
	if err := rootCmd.Execute(); err != nil {
		util.ExitError(err)
	}
}

func run(cmd *cobra.Command, args []string) error {
	m, err := cmd.Flags().GetString("module")
	if err != nil {
		return errors.New("\"module\" flag declared as non-string. Please correct your code")
	}
	path := filepath.Join(util.CurrentDir(), args[0])
	if err := MakeDirectory(path); err != nil {
		util.ExitError(err)
	}
	if err := CreateFiles(path); err != nil {
		util.ExitError(err)
	}
	if m != "" {
		RunGoMod(m, path)
	}
	return nil
}

// create directory that user pass name
func MakeDirectory(path string) error {
	err := os.Mkdir(path, 0777)
	if err != nil {
		return err
	}
	return nil
}

// create .gitignore, .golangci.yml, LICENSE, README.md, main.go, test.yml, lint.yml
func CreateFiles(path string) error {
	fis, err := local.ReadDir("files")
	if err != nil {
		return err
	}
	for _, fi := range fis {
		in, err := local.Open(filepath.Join("files", fi.Name()))
		if err != nil {
			return err
		}
		out, err := os.Create(filepath.Join(path, filepath.Base(fi.Name())))
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
	actionsPath := filepath.Join(path, ".github", "workflows")
	// i don't know chmod 0777 is best for this usage.
	if err := os.MkdirAll(filepath.Join(actionsPath), 0777); err != nil {
		return err
	}

	//move test.yml and lint.yml
	if err := os.Rename(filepath.Join(path, "test.yml"), filepath.Join(actionsPath, "test.yml")); err != nil {
		return err
	}
	if err := os.Rename(filepath.Join(path, "lint.yml"), filepath.Join(actionsPath, "lint.yml")); err != nil {
		return err
	}
	return nil
}

func RunGoMod(arg, path string) error {
	if err := os.Chdir(path); err != nil {
		return err
	}
	cmd := exec.Command("go", "mod", "init", arg)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + string(output))
	}
	log.Println(string(output))
	return nil
}
