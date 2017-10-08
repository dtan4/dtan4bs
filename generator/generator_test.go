package generator

import (
	"io/ioutil"
	"os"
	"path"
	"strings"
	"testing"
)

func TestNewGenerator(t *testing.T) {
	testcases := []struct {
		baseDir string
	}{
		{
			baseDir: "/foo/bar/baz",
		},
		{
			baseDir: "",
		},
	}

	for _, tc := range testcases {
		g := NewGenerator(tc.baseDir)

		if got, want := g.baseDir, tc.baseDir; got != want {
			t.Errorf("baseDir got: %s, want baseDir: %s", got, want)
		}
	}
}

func TestGenerate(t *testing.T) {
	baseDir, err := ioutil.TempDir("", "dtan4bs-generator-generate")
	if err != nil {
		t.Fatalf("cannot create tempdir: %s", err)
	}
	defer os.RemoveAll(baseDir)

	g := &Generator{
		baseDir: baseDir,
	}

	testcases := []struct {
		filename string
		dir      string
		metadata map[string]string
	}{
		{
			filename: "README.md",
			dir:      "",
			metadata: map[string]string{},
		},
	}

	for _, tc := range testcases {
		if err := g.Generate(tc.filename, tc.dir, tc.metadata); err != nil {
			t.Errorf("error should not be raised: %s", err)
		}

		fp := path.Join(baseDir, tc.dir, tc.filename)

		s, err := os.Stat(fp)
		if os.IsNotExist(err) {
			t.Errorf("%s was not created: %s", fp, err)
			continue
		}

		if s.IsDir() {
			t.Errorf("%s is not a file", fp)
		}
	}
}

func TestGenerate_error(t *testing.T) {
	baseDir, err := ioutil.TempDir("", "dtan4bs-generator-generate")
	if err != nil {
		t.Fatalf("cannot create tempdir: %s", err)
	}
	defer os.RemoveAll(baseDir)

	g := &Generator{
		baseDir: baseDir,
	}

	testcases := []struct {
		filename string
		dir      string
		metadata map[string]string
		errMsg   string
	}{
		{
			filename: "foobarbaz",
			dir:      "",
			metadata: map[string]string{},
			errMsg:   "template of foobarbaz does not exist",
		},
	}

	for _, tc := range testcases {
		err := g.Generate(tc.filename, tc.dir, tc.metadata)
		if err == nil {
			t.Error("error should be raised")
			continue
		}

		if got, want := err.Error(), tc.errMsg; strings.Contains(got, want) {
			t.Errorf("error message does not contain %q, got: %q", want, got)
		}
	}
}
