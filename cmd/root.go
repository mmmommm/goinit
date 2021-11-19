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
			return
		}
		if err := createDotPrefixFiles(); err != nil {
			return
		}
		// if err := runGoMod(); err != nil {
		// 	exitError(err)
		// }
	}

	if err := rootCmd.Execute(); err != nil {
		exitError(err)
	}
}

//go:embed files
var local embed.FS

func createFiles() error {
	currentDir, err := os.Getwd()
	if err != nil {
		return err
	}

	if err != nil {
		return err
	}
	fis, err := local.ReadDir("files")
	if err != nil {
		log.Fatal(err)
	}
	for _, fi := range fis {
		in, err := local.Open(filepath.Join("files", fi.Name()))
		if err != nil {
			log.Fatal(err)
		}
		out, err := os.Create(filepath.Join(currentDir, "example", filepath.Base(fi.Name())))
		if err != nil {
			log.Fatal(err)
		}
		io.Copy(out, in)
		out.Close()
		in.Close()
		log.Println("exported", filepath.Base(fi.Name()))
	}
	return nil
}

func createDotPrefixFiles() error {
	currentDir, err := os.Getwd()
	if err != nil {
		return err
	}
	actionsPath := filepath.Join(currentDir, "example", ".github", "workflows", "ci.yaml")
	gitignorePath := filepath.Join(currentDir, "example", ".gitignore")

	actionIn, err := os.Open(filepath.Join(currentDir, "files", ".github", "workflows", "ci.yaml"))
	if err != nil {
		return err
	}
	gitignoreIn, err := os.Open(filepath.Join(currentDir, ".gitignore"))
	if err != nil {
		return err
	}
	if err := os.MkdirAll(actionsPath, os.ModePerm); err != nil {
		return err
	}
	actionOut, err := os.Create(actionsPath)
	if err != nil {
		return err
	}
	gitignoreOut, err := os.Create(gitignorePath)
	if err != nil {
		return err
	}
	io.Copy(actionOut, actionIn)
	io.Copy(gitignoreOut, gitignoreIn)
	return nil
}

// func runGoMod() error {
// 	pkgName := os.Args[0]
// 	if err := exec.Command("go", "mod", "init", pkgName).Start(); err != nil {
// 		return err
// 	}
// 	return nil
// }
