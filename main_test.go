package main_test

import (
	"errors"
	"fmt"
	"testing"

	main "github.com/youwkey/alfred-ghq"
)

func TestCheckExec(t *testing.T) {
	t.Parallel()

	type Test struct {
		in  string
		err error
	}

	tests := []Test{
		{in: "", err: main.ErrPathEmpty},
		{in: "testdata/invalid", err: main.ErrPathNotExists},
		{in: "testdata/permission_0666", err: main.ErrPathNotExec},
		{in: "testdata/permission_0777", err: nil},
	}

	for i, test := range tests {
		i, test := i, test
		t.Run(fmt.Sprintf("#%d:CheckExec", i), func(t *testing.T) {
			t.Parallel()
			err := main.CheckExec(test.in)
			if !errors.Is(err, test.err) {
				t.Errorf("#%d: got err: %v want err: %v", i, err, test.err)
			}
		})
	}
}
