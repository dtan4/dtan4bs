package generator

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
