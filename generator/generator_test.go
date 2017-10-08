package generator

import (
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
