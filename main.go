//go:generate go install -v github.com/kevinburke/go-bindata/v4/go-bindata
//go:generate go-bindata -prefix res/ -pkg assets -o assets/assets.go res/Rocket.Chat.lnk
//go:generate go install -v github.com/josephspurrier/goversioninfo/cmd/goversioninfo
//go:generate goversioninfo -icon=res/papp.ico -manifest=res/papp.manifest
package main

import (
	"encoding/json"
	"io"
	"os"
	"path"
	"strings"

	"github.com/portapps/portapps/v3"
	"github.com/portapps/portapps/v3/pkg/log"
	"github.com/portapps/portapps/v3/pkg/shortcut"
	"github.com/portapps/portapps/v3/pkg/utl"
	"github.com/portapps/rocketchat-portable/assets"
)

type config struct {
	Cleanup bool `yaml:"cleanup" mapstructure:"cleanup"`
}

var (
	app *portapps.App
	cfg *config
)

func init() {
	var err error

	// Default config
	cfg = &config{
		Cleanup: false,
	}

	// Init app
	if app, err = portapps.NewWithCfg("rocketchat-portable", "Rocket.Chat", cfg); err != nil {
		log.Fatal().Err(err).Msg("Cannot initialize application. See log file for more info.")
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
		rawSettings, err := os.ReadFile(updateSettingsPath)
		if err == nil {
			jsonMapSettings := make(map[string]interface{})
			json.Unmarshal(rawSettings, &jsonMapSettings)
			log.Info().Msgf("Current update settings: %s", jsonMapSettings)

			jsonMapSettings["autoUpdate"] = false
			jsonMapSettings["canUpdate"] = false
			log.Info().Msgf("New update settings: %s", jsonMapSettings)

			jsonSettings, err := json.Marshal(jsonMapSettings)
			if err != nil {
				log.Error().Err(err).Msg("Update settings marshal")
			}
			err = os.WriteFile(updateSettingsPath, jsonSettings, 0644)
			if err != nil {
				log.Error().Err(err).Msg("Write update settings")
			}
		}
	} else {
		fo, err := os.Create(updateSettingsPath)
		if err != nil {
			log.Error().Err(err).Msg("Cannot create update.json")
		}
		defer fo.Close()
		if _, err = io.Copy(fo, strings.NewReader(`{"autoUpdate":false,"canUpdate":false}`)); err != nil {
			log.Error().Err(err).Msg("Cannot write to update.json")
		}
	}

	// Copy default shortcut
	shortcutPath := path.Join(os.Getenv("APPDATA"), "Microsoft", "Windows", "Start Menu", "Programs", "Rocket.Chat Portable.lnk")
	defaultShortcut, err := assets.Asset("Rocket.Chat.lnk")
	if err != nil {
		log.Error().Err(err).Msg("Cannot load asset Rocket.Chat.lnk")
	}
	err = os.WriteFile(shortcutPath, defaultShortcut, 0644)
	if err != nil {
		log.Error().Err(err).Msg("Cannot write default shortcut")
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
		log.Error().Err(err).Msg("Cannot create shortcut")
	}
	defer func() {
		if err := os.Remove(shortcutPath); err != nil {
			log.Error().Err(err).Msg("Cannot remove shortcut")
		}
	}()

	defer app.Close()
	app.Launch(os.Args[1:])
}
