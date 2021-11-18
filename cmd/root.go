package cmd

import (
	"embed"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path"

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

//go:embed files
var local embed.FS

func createFiles() error {
	fis, err := local.ReadDir("files")
	if err != nil {
		log.Fatal(err)
	}
	for _, fi := range fis {
		in, err := local.Open(path.Join("files", fi.Name()))
		if err != nil {
			log.Fatal(err)
		}
		out, err := os.Create(path.Base(fi.Name()))
		if err != nil {
			log.Fatal(err)
		}
		io.Copy(out, in)
		out.Close()
		in.Close()
		log.Println("exported", path.Base(fi.Name()))
	}
	return nil
}

func runGoMod() error {
	pkgName := os.Args[0]
	if err := exec.Command("go", "mod", "init", pkgName).Start(); err != nil {
		return err
	}
	return nil
}
