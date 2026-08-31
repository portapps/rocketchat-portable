//go:generate go install -v github.com/josephspurrier/goversioninfo/cmd/goversioninfo
package main

import (
	_ "embed"
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/portapps/portapps/v3"
	"github.com/portapps/portapps/v3/pkg/files"
	"github.com/portapps/portapps/v3/pkg/log"
	"github.com/portapps/portapps/v3/pkg/shortcut"
)

//go:embed res/Rocket.Chat.lnk
var defaultShortcut []byte

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
	if err := os.MkdirAll(app.DataPath, 0o755); err != nil {
		log.Fatal().Err(err).Msg("Cannot create data path")
	}
	app.Process = filepath.Join(app.AppPath, "Rocket.Chat.exe")
	app.Args = []string{
		"--user-data-dir=" + app.DataPath,
	}

	// Cleanup on exit
	if cfg.Cleanup {
		defer func() {
			files.Cleanup(filepath.Join(os.Getenv("APPDATA"), "Rocket.Chat"))
		}()
	}

	updateSettingsPath := filepath.Join(app.DataPath, "update.json")
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
	shortcutPath := filepath.Join(os.Getenv("APPDATA"), "Microsoft", "Windows", "Start Menu", "Programs", "Rocket.Chat Portable.lnk")
	err := os.WriteFile(shortcutPath, defaultShortcut, 0644)
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
