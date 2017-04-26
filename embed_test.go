package goef

import (
	"testing"
)

func TestFileToStringLiteral(t *testing.T) {
	sl, err := FileToStringLiteral("hello.txt")
	if err != nil {
		t.Error(err)
	}
	t.Log(sl)
}
