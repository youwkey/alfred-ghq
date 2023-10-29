// Copyright 2023 youwkey. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// main
package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/youwkey/alfred-go"
)

// check path errors.
var (
	ErrPathEmpty     = errors.New("path not set")
	ErrPathNotExists = errors.New("path not exists")
	ErrPathNotExec   = errors.New("path not exec")
)

// CheckExec check path is executable file.
func CheckExec(path string) error {
	if path == "" {
		return ErrPathEmpty
	} else if f, err := os.Stat(path); err != nil {
		return fmt.Errorf("%w '%s'", ErrPathNotExists, path)
	} else if m := f.Mode(); m.IsDir() || 0o111&m.Perm() == 0 {
		return fmt.Errorf("%w '%s'", ErrPathNotExec, path)
	}

	return nil
}

func getIcon(host string) *alfred.Icon {
	path := "./assets/repo.svg"

	switch host {
	case "github.com":
		path = "./assets/github.svg"
	case "bitbucket.org":
		path = "./assets/bitbucket.svg"
	}

	return alfred.NewIcon(path)
}

func main() {
	sf := alfred.ScriptFilter{}
	ghqPath := os.Getenv("GHQ_PATH")

	if err := CheckExec(ghqPath); err != nil {
		item := alfred.NewInvalidItem("Error").Subtitle(err.Error())
		sf.Items().Append(item)
		_ = sf.Output()

		return
	}

	bytes, err := exec.Command(ghqPath, "root").Output()
	if err != nil {
		item := alfred.NewInvalidItem("Error").Subtitle(err.Error())
		sf.Items().Append(item)
		_ = sf.Output()

		return
	}

	ghqRoot := strings.TrimRight(string(bytes), "\n")

	bytes, err = exec.Command(ghqPath, "list").Output()
	if err != nil {
		item := alfred.NewInvalidItem("Error").Subtitle(err.Error())
		sf.Items().Append(item)
		_ = sf.Output()

		return
	}

	for _, v := range strings.Fields(string(bytes)) {
		absPath := filepath.Join(ghqRoot, v)
		repoURL := "https://" + v
		modCmd := alfred.NewModifier().Arg(absPath).Subtitle("Browse in Terminal").Icon(alfred.IconExecutableBinary)
		modAlt := alfred.NewModifier().Arg(repoURL).Subtitle("Open in Finder").Icon(alfred.IconFinder)
		item := alfred.
			NewItem(v).
			UID(v).
			Arg(absPath).
			Match(strings.Join(strings.Split(v, "/"), " ")).
			Type(alfred.ItemTypeFile).
			Mods(alfred.NewModifiers().Cmd(modCmd).Alt(modAlt)).
			Icon(getIcon(strings.Split(v, "/")[0])).
			QuicklookURL(repoURL)

		sf.Items().Append(item)
	}

	if sf.Items().Length() == 0 {
		sf.Items().Append(alfred.NewInvalidItem("No Results"))
	}

	_ = sf.Output()
}
