package generator

import (
	"io/ioutil"
	"os"
	"path"

	"github.com/pkg/errors"
)

const (
	templatePrefix = "_template/"
)

// Generator represents the file generator
type Generator struct {
	baseDir string
}

// NewGenerator creates new Generator object
func NewGenerator(baseDir string) *Generator {
	return &Generator{
		baseDir: baseDir,
	}
}

// Generate generates the given file from template
func (g *Generator) Generate(filename, dir string, metadata map[string]string) error {
	data, err := Asset(templatePrefix + filename)
	if err != nil {
		return errors.Wrapf(err, "template of %q is not found", filename)
	}

	d := path.Join(g.baseDir, dir)
	if _, err := os.Stat(d); os.IsNotExist(err) {
		if err2 := os.MkdirAll(d, 0755); err2 != nil {
			return errors.Wrapf(err2, "failed to create %q", d)
		}
	}

	fp := path.Join(d, filename)

	if err := ioutil.WriteFile(fp, data, 0644); err != nil {
		return errors.Wrapf(err, "failed to create %q", fp)
	}

	return nil
}
