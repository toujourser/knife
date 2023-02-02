package random

import (
	"regexp"
	"testing"
)

func TestRandString(t *testing.T) {
	pattern := `^[a-zA-Z]+$`
	reg := regexp.MustCompile(pattern)

	randStr := RandString(6)

	if len(randStr) != 6 {
		t.Fatal("length failed")
	}

	if !reg.MatchString(randStr) {
		t.Fatal("MatchString failed")
	}
}
