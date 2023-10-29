// Copyright 2023 youwkey. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"path/filepath"
	"regexp"
	"strings"

	"github.com/youwkey/alfred-go"
)

var reg = regexp.MustCompile(`([0-9]+)`)

// GHQRepo represents the ghq managed path.
type GHQRepo struct {
	RootPath string
	RepoPath string
	Host     string
	User     string
	Name     string
}

// NewGHQRepo returns a new initialized NewGHQRepo.
func NewGHQRepo(rootPath, repoPath string) *GHQRepo {
	elems := strings.Split(repoPath, "/")

	return &GHQRepo{
		RootPath: rootPath,
		RepoPath: repoPath,
		Host:     elems[0],
		User:     elems[1],
		Name:     elems[2],
	}
}

// AbsPath returns an absolute filepath.
func (g *GHQRepo) AbsPath() string {
	return filepath.Join(g.RootPath, g.RepoPath)
}

// URL returns a repository url.
func (g *GHQRepo) URL() string {
	return "https://" + g.RepoPath
}

// Icon return a repository icon.
func (g *GHQRepo) Icon() *alfred.Icon {
	path := "./assets/repo.svg"

	switch g.Host {
	case "github.com":
		path = "./assets/github.svg"
	case "bitbucket.org":
		path = "./assets/bitbucket.svg"
	}

	return alfred.NewIcon(path)
}

// MatchValue returns alfred match string.
func (g *GHQRepo) MatchValue() string {
	var others []string

	if reg.MatchString(g.Name) {
		others = append(others, reg.FindAllString(g.Name, -1)...)
	}

	return strings.Join(append([]string{g.Host, g.User, g.Name}, others...), " ")
}
