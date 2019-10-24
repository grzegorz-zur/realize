package realize

import (
	"os"
	"path/filepath"
	"testing"
)

func TestWdir(t *testing.T) {
	expected, err := os.Getwd()
	if err != nil {
		t.Error(err)
	}
	result := Wdir()
	if result != expected {
		t.Error("Expected", filepath.Base(expected), "instead", result)
	}
}

func TestExt(t *testing.T) {
	paths := map[string]string{
		"/test/a/b/c":        "",
		"/test/a/ac.go":      "go",
		"/test/a/ac.test.go": "go",
		"/test/a/ac_test.go": "go",
		"/test/./ac_test.go": "go",
		"/test/a/.test":      "test",
		"/test/a/.":          "",
	}
	for i, v := range paths {
		if ext(i) != v {
			t.Error("Wrong extension", ext(i), v)
		}
	}

}
