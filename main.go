//go:generate go install -v github.com/kevinburke/go-bindata/go-bindata
//go:generate go-bindata -prefix res/ -pkg assets -o assets/assets.go res/Rocket.Chat.lnk
//go:generate go install -v github.com/josephspurrier/goversioninfo/cmd/goversioninfo
//go:generate goversioninfo -icon=res/papp.ico -manifest=res/papp.manifest
package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strings"

	. "github.com/portapps/portapps"
	"github.com/portapps/portapps/pkg/shortcut"
	"github.com/portapps/portapps/pkg/utl"
	"github.com/portapps/rocketchat-portable/assets"
)

type config struct {
	Cleanup bool `yaml:"cleanup" mapstructure:"cleanup"`
}

var (
	app *App
	cfg *config
)

func init() {
	var err error

	// Default config
	cfg = &config{
		Cleanup: false,
	}

	// Init app
	if app, err = NewWithCfg("rocketchat-portable", "Rocket.Chat", cfg); err != nil {
		Log.Fatal().Err(err).Msg("Cannot initialize application. See log file for more info.")
	}
}

func main() {
	utl.CreateFolder(app.DataPath)
	app.Process = utl.PathJoin(app.AppPath, "Rocket.Chat.exe")
	app.Args = []string{
		"--user-data-dir=" + app.DataPath,
	}

	// Cleanup on exit
	if cfg.Cleanup {
		defer func() {
			utl.Cleanup([]string{
				path.Join(os.Getenv("APPDATA"), "Rocket.Chat"),
			})
		}()
	}

	updateSettingsPath := path.Join(app.DataPath, "update.json")
	if _, err := os.Stat(updateSettingsPath); err == nil {
		rawSettings, err := ioutil.ReadFile(updateSettingsPath)
		if err == nil {
			jsonMapSettings := make(map[string]interface{})
			json.Unmarshal(rawSettings, &jsonMapSettings)
			Log.Info().Msgf("Current update settings: %s", jsonMapSettings)

			jsonMapSettings["autoUpdate"] = false
			jsonMapSettings["canUpdate"] = false
			Log.Info().Msgf("New update settings: %s", jsonMapSettings)

			jsonSettings, err := json.Marshal(jsonMapSettings)
			if err != nil {
				Log.Error().Err(err).Msg("Update settings marshal")
			}
			err = ioutil.WriteFile(updateSettingsPath, jsonSettings, 0644)
			if err != nil {
				Log.Error().Err(err).Msg("Write update settings")
			}
		}
	} else {
		fo, err := os.Create(updateSettingsPath)
		if err != nil {
			Log.Error().Err(err).Msg("Cannot create update.json")
		}
		defer fo.Close()
		if _, err = io.Copy(fo, strings.NewReader(`{"autoUpdate":false,"canUpdate":false}`)); err != nil {
			Log.Error().Err(err).Msg("Cannot write to update.json")
		}
	}

	// Copy default shortcut
	shortcutPath := path.Join(os.Getenv("APPDATA"), "Microsoft", "Windows", "Start Menu", "Programs", "Rocket.Chat Portable.lnk")
	defaultShortcut, err := assets.Asset("Rocket.Chat.lnk")
	if err != nil {
		Log.Error().Err(err).Msg("Cannot load asset Rocket.Chat.lnk")
	}
	err = ioutil.WriteFile(shortcutPath, defaultShortcut, 0644)
	if err != nil {
		Log.Error().Err(err).Msg("Cannot write default shortcut")
	}

	// Update default shortcut
	err = shortcut.Create(shortcut.Shortcut{
		ShortcutPath:     shortcutPath,
		TargetPath:       app.Process,
		Arguments:        shortcut.Property{Clear: true},
		Description:      shortcut.Property{Value: "Rocket.Chat Portable by Portapps"},
		IconLocation:     shortcut.Property{Value: app.Process},
		WorkingDirectory: shortcut.Property{Value: app.AppPath},
	})
	if err != nil {
		Log.Error().Err(err).Msg("Cannot create shortcut")
	}
	defer func() {
		if err := os.Remove(shortcutPath); err != nil {
			Log.Error().Err(err).Msg("Cannot remove shortcut")
		}
	}()

	app.Launch(os.Args[1:])
}
