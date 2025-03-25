package ch02

import "testing"

func TestConst(t *testing.T) {
	const (
		Mon = iota + 1
		Tue
		Wed
		Thu
		Fri
		Sat
		Sun
	)
	t.Log(Mon, Tue, Wed, Thu, Fri, Sat, Sun)
}
