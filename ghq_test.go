// Copyright 2023 youwkey. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"testing"
)

func TestGHQRepo_MatchValue(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		path string
		want string
	}{
		{name: "standard", path: "a/b/c", want: "a b c"},
		{name: "with single numbers", path: "a/b/c100", want: "a b c100 100"},
		{name: "with multi numbers", path: "a/b/c100-200", want: "a b c100-200 100 200"},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			repo := NewGHQRepo("", test.path)
			if got := repo.MatchValue(); got != test.want {
				t.Errorf("MatchValue(): got = %s, want = %s", got, test.want)
			}
		})
	}
}
