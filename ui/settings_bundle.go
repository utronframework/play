//+build ignore

package main

import (
	"bytes"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/gu-io/gu/common"
	"github.com/gu-io/gu/common/themes/styleguide"
)

func main() {
	settings, err := getSettings()
	if err != nil {
		panic(err)
	}

	var theme bytes.Buffer

	if err := styleguide.Render(&theme, settings.Theme); err != nil {
		panic(err)
	}

	cssPublicDir := filepath.Join(settings.Public.Path, "css")
	cssPublic := filepath.Join(settings.Public.Path, "css/theme.css")

	if err := os.MkdirAll(cssPublicDir, 0777); err != nil && err != os.ErrExist {
		panic(err)
	}

	themeFile, err := os.Create(cssPublic)
	if err != nil {
		panic(err)
	}

	defer themeFile.Close()

	if _, err := theme.WriteTo(themeFile); err != nil {
		panic(err)
	}
}

func getSettings() (common.Settings, error) {
	var config common.Settings

	// Load settings into configuration.
	if _, err := toml.DecodeFile("./settings.toml", &config); err != nil {
		return config, err
	}

	if err := config.Validate(); err != nil {
		return config, err
	}

	return config, nil
}
