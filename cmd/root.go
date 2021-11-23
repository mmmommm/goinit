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
	"github.com/spf13/viper"
)

var (
	// Used for flags.
	module string
	v      = "0.1.0"

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
			err := run(cmd, args)
			return err
		},
		Version:          v,
		TraverseChildren: true,
	}
	if err := rootCmd.Execute(); err != nil {
		util.ExitError(err)
	}
}

func run(cmd *cobra.Command, args []string) error {
	// get flag value
	initDefaultModuleFlag()
	m, err := cmd.Flags().GetString("module")
	if err != nil {
		cmd.Println("\"module\" flag declared as non-string. Please correct your code")
		return err
	}
	path := filepath.Join(util.CurrentDir(), args[0])
	if err := MakeDirectory(path); err != nil {
		util.ExitError(err)
	}
	if err := CreateFiles(path); err != nil {
		util.ExitError(err)
	}
	if m != "" {
		RunGoMod(m)
	}
	cmd.Println()
	return nil
}

// create directory that user pass name
func MakeDirectory(path string) error {
	// MkDirAll func is better?
	err := os.Mkdir(path, 0777)
	if err != nil {
		return err
	}
	return nil
}

//go:embed files/*
var local embed.FS

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
	//move test.yml and lint.yml
	actionsPath := filepath.Join(path, ".github", "workflows")
	// i don't know chmod 0777 is best for this usage.
	if err := os.MkdirAll(filepath.Join(actionsPath), 0777); err != nil {
		return err
	}
	if err := os.Rename(filepath.Join(path, "test.yml"), filepath.Join(actionsPath, "test.yml")); err != nil {
		return err
	}
	if err := os.Rename(filepath.Join(path, "lint.yml"), filepath.Join(actionsPath, "lint.yml")); err != nil {
		return err
	}
	return nil
}

func RunGoMod(arg string) error {
	cmd := exec.Command("go", "mod", "init", arg)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + string(output))
	}
	log.Println(string(output))
	return nil
}

// initialize module flag
// FYI https://github.com/spf13/cobra/blob/9e1d6f1c2aa8df64b6a6ba39e92517f68580d653/command.go?_pjax=%23js-repo-pjax-container%2C%20div%5Bitemtype%3D%22http%3A%2F%2Fschema.org%2FSoftwareSourceCode%22%5D%20main%2C%20%5Bdata-pjax-container%5D#L1048
func initDefaultModuleFlag() {
	if rootCmd.Flags().Lookup("module") == nil {
		usage := "module name for"
		if rootCmd.Name() == "" {
			usage += "this command"
		} else {
			usage += rootCmd.Name()
		}
		if rootCmd.Flags().ShorthandLookup("m") == nil {
			rootCmd.Flags().StringVarP(&module, "module", "m", "", usage)
		} else {
			rootCmd.Flags().StringVar(&module, "module", "", usage)
		}
	}
}

func init() {
	rootCmd.Flags().StringP("module", "m", "", "-m option run `go mod init ${argument}`")
	cobra.CheckErr(viper.BindPFlag("module", rootCmd.PersistentFlags().Lookup("module")))
}
