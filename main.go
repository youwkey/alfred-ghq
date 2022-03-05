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
	ghqPath := os.Getenv("GHQ_PATH")

	if ghqPath == "" {
		item := alfred.NewInvalidItem("Error").Subtitle("require workflow environment variable 'GHQ_PATH'")
		sf.Items().Append(item)
		_ = sf.Output()
		return
	} else if f, err := os.Stat(ghqPath); err != nil {
		item := alfred.NewInvalidItem("Error").Subtitle(ghqPath + " not exists")
		sf.Items().Append(item)
		_ = sf.Output()
		return
	} else if m := f.Mode(); m.IsDir() || 0111&m.Perm() == 0 {
		item := alfred.NewInvalidItem("Error").Subtitle(ghqPath + " not exec")
		sf.Items().Append(item)
		_ = sf.Output()
		return
	}

	var ghqRoot string
	if bytes, err := exec.Command(ghqPath, "root").Output(); err != nil {
		item := alfred.NewInvalidItem("Error").Subtitle(err.Error())
		sf.Items().Append(item)
		_ = sf.Output()
		return
	} else {
		ghqRoot = strings.TrimRight(string(bytes), "\n")
	}

	if bytes, err := exec.Command(ghqPath, "list").Output(); err != nil {
		item := alfred.NewInvalidItem("Error").Subtitle(err.Error())
		sf.Items().Append(item)
		_ = sf.Output()
		return
	} else {
		for _, v := range strings.Fields(string(bytes)) {
			absPath := filepath.Join(ghqRoot, v)
			repoURL := "https://" + v
			modCtrl := alfred.NewModifier().Arg(absPath).Subtitle("Browse in Terminal").Icon(alfred.IconExecutableBinary)
			modAlt := alfred.NewModifier().Arg(repoURL).Subtitle("Open URL").Icon(alfred.IconBookmark)
			item := alfred.
				NewItem(v).
				UID(v).
				Arg(absPath).
				Match(strings.Join(strings.Split(v, "/"), " ")).
				Type(alfred.ItemTypeFile).
				Mods(alfred.NewModifiers().Ctrl(modCtrl).Alt(modAlt)).
				Icon(alfred.NewIconWithType(absPath, alfred.IconTypeFileIcon)).
				QuicklookURL(repoURL)

			sf.Items().Append(item)
		}
		if sf.Items().Length() == 0 {
			sf.Items().Append(alfred.NewInvalidItem("No Results"))
		}
	}

	_ = sf.Output()
}
