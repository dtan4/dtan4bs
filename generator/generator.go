package generator

import (
	"io/ioutil"
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

	fp := path.Join(g.baseDir, dir, filename)

	if err := ioutil.WriteFile(fp, data, 0644); err != nil {
		return errors.Wrapf(err, "failed to create %q", fp)
	}

	return nil
}
