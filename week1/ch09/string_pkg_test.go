package ch09

import (
	"strings"
	"testing"
)

func TestStringPkg(t *testing.T) {
	s := "A,B,C"
	parts := strings.Split(s,",")
	t.Log(parts)
	t.Log(strings.Join(parts,"->"))
}