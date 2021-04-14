package main

import (
	"github.com/youwkey/alfred-go"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	sf := alfred.ScriptFilter{}
	sf.SetEmptyTitle("No Results", "")
	ghqPath := os.Getenv("GHQ_PATH")

	if ghqPath == "" {
		sf.Append(alfred.Item{
			Title:    "Error",
			Subtitle: "require workflow environment variable 'GHQ_PATH'",
			Valid:    false,
		})
		sf.Output()
		return
	} else if f, err := os.Stat(ghqPath); err != nil {
		sf.Append(alfred.Item{
			Title:    "Error",
			Subtitle: ghqPath + " not exists",
			Valid:    false,
		})
		sf.Output()
		return
	} else if m := f.Mode(); m.IsDir() || 0111&m.Perm() == 0 {
		sf.Append(alfred.Item{
			Title:    "Error",
			Subtitle: ghqPath + " not directory",
			Valid:    false,
		})
		sf.Output()
		return
	}

	var ghqRoot string
	if bytes, err := exec.Command(ghqPath, "root").Output(); err != nil {
		sf.Append(alfred.Item{
			Title:    "Error",
			Subtitle: err.Error(),
			Valid:    false,
		})
		sf.Output()
		return
	} else {
		ghqRoot = strings.TrimRight(string(bytes), "\n")
	}

	if bytes, err := exec.Command(ghqPath, "list").Output(); err != nil {
		sf.Append(alfred.Item{
			Title:    "Error",
			Subtitle: err.Error(),
			Valid:    false,
		})
		sf.Output()
		return
	} else {
		for _, v := range strings.Fields(string(bytes)) {
			absPath := filepath.Join(ghqRoot, v)
			repoURL := "https://" + v
			sf.Append(alfred.Item{
				Uid:   v,
				Title: v,
				Arg:   absPath,
				Match: strings.Join(strings.Split(v, "/"), " "),
				Type:  alfred.File,
				Mods: map[alfred.ModKey]alfred.Mod{
					alfred.Ctrl: {
						Arg:      absPath,
						Subtitle: "Browse in Terminal",
						Icon:     alfred.IconExecutableBinary,
					},
					alfred.Alt: {
						Arg:      repoURL,
						Subtitle: "Open URL",
						Icon:     alfred.IconBookmark,
					},
				},
				Icon: &alfred.Icon{
					Type: alfred.FileIcon,
					Path: absPath,
				},
				QuicklookURL: repoURL,
			})
		}
	}

	sf.Output()
}
