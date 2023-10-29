// Copyright 2023 youwkey. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"regexp"
	"strings"
)

var reg = regexp.MustCompile(`([0-9]+)`)

// GHQRepo represents the ghq managed path.
type GHQRepo struct {
	RootPath string
	RepoPath string
}

// NewGHQRepo returns a new initialized NewGHQRepo.
func NewGHQRepo(rootPath, repoPath string) *GHQRepo {
	return &GHQRepo{
		RootPath: rootPath,
		RepoPath: repoPath,
	}
}

// MatchValue returns alfred match string.
func (g GHQRepo) MatchValue() string {
	var others []string

	elems := strings.Split(g.RepoPath, "/")
	host, user, name := elems[0], elems[1], elems[2]

	if reg.MatchString(name) {
		others = append(others, reg.FindAllString(name, -1)...)
	}

	return strings.Join(append([]string{host, user, name}, others...), " ")
}
