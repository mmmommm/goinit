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
	cfgFile     string
	v           = "0.1.0"

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
	}
	if err := rootCmd.Execute(); err != nil {
		util.ExitError(err)
	}
}

func run(cmd *cobra.Command, args []string) error {
	// get flag value
	mFlag, _ := cmd.Flags().GetString("module")
	vFlag, _ := cmd.Flags().GetBool("version")

	// if version is true return version and return
	if !vFlag {
		path := filepath.Join(util.CurrentDir(), args[0])
		if err := MakeDirectory(path); err != nil {
			util.ExitError(err)
		}
		if err := CreateFiles(path); err != nil {
			util.ExitError(err)
		}
		if mFlag != "" {
			RunGoMod(mFlag)
		}
	} else {
		fmt.Println("goinit version:", v)
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

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.SetConfigName(".goinit")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.Flags().StringP("module", "m", "", "-m option run `go mod init ${argument}`")
	rootCmd.Flags().BoolP("version", "v", false, "-v option show version")
}
