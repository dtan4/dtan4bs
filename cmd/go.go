package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/dtan4/dtan4bs/generator"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var (
	goFiles = []string{
		"LICENSE",
		"README.md",
		"main.go",
	}
)

// goCmd represents the go command
var goCmd = &cobra.Command{
	Use:   "go",
	Short: "Generate Go project template",
	RunE:  doGo,
}

func doGo(cmd *cobra.Command, args []string) error {
	// TODO: set baseDir via option
	baseDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return errors.Wrap(err, "failed to get current directory")
	}

	g := generator.NewGenerator(baseDir)

	for _, filename := range goFiles {
		if err := g.Generate(filename, "", map[string]string{}); err != nil {
			return errors.Wrap(err, "failed to generate Go project files")
		}
	}

	fmt.Println("successfully generated")

	return nil
}

func init() {
	RootCmd.AddCommand(goCmd)
}
