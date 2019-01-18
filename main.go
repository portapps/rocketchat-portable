//go:generate go install -v github.com/kevinburke/go-bindata/go-bindata
//go:generate go-bindata -prefix res/ -pkg assets -o assets/assets.go res/Rocket.Chat.lnk
//go:generate go install -v github.com/josephspurrier/goversioninfo/cmd/goversioninfo
//go:generate goversioninfo -icon=res/papp.ico
package main

import (
	_ "github.com/kevinburke/go-bindata"

	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strings"

	. "github.com/portapps/portapps"
	"github.com/portapps/rocketchat-portable/assets"
)

func init() {
	Papp.ID = "rocketchat-portable"
	Papp.Name = "Rocket.Chat"
	Init()
}

func main() {
	Papp.AppPath = AppPathJoin("app")
	Papp.DataPath = CreateFolder(AppPathJoin("data"))
	Papp.Process = PathJoin(Papp.AppPath, "Rocket.Chat.exe")
	Papp.Args = nil
	Papp.WorkingDir = Papp.AppPath

	updateSettingsPath := path.Join(Papp.DataPath, "update.json")
	if _, err := os.Stat(updateSettingsPath); err == nil {
		rawSettings, err := ioutil.ReadFile(updateSettingsPath)
		if err == nil {
			jsonMapSettings := make(map[string]interface{})
			json.Unmarshal(rawSettings, &jsonMapSettings)
			Log.Info("Current update settings:", jsonMapSettings)

			jsonMapSettings["autoUpdate"] = false
			jsonMapSettings["canUpdate"] = false
			Log.Info("New settings:", jsonMapSettings)

			jsonSettings, err := json.Marshal(jsonMapSettings)
			if err != nil {
				Log.Error("Update settings marshal:", err)
			}
			err = ioutil.WriteFile(updateSettingsPath, jsonSettings, 0644)
			if err != nil {
				Log.Error("Write Update settings:", err)
			}
		}
	} else {
		fo, err := os.Create(updateSettingsPath)
		if err != nil {
			Log.Error("Cannot create update.json:", err)
		}
		defer fo.Close()
		if _, err = io.Copy(fo, strings.NewReader(`{"autoUpdate":false,"canUpdate":false}`)); err != nil {
			Log.Error("Cannot write to update.json:", err)
		}
	}

	// Copy default shortcut
	shortcutPath := path.Join(os.Getenv("APPDATA"), "Microsoft", "Windows", "Start Menu", "Programs", "Rocket.Chat Portable.lnk")
	defaultShortcut, err := assets.Asset("Rocket.Chat.lnk")
	if err != nil {
		Log.Error("Cannot load asset Rocket.Chat.lnk:", err)
	}
	err = ioutil.WriteFile(shortcutPath, defaultShortcut, 0644)
	if err != nil {
		Log.Error("Cannot write default shortcut:", err)
	}

	// Update default shortcut
	err = CreateShortcut(WindowsShortcut{
		ShortcutPath:     shortcutPath,
		TargetPath:       Papp.Process,
		Arguments:        WindowsShortcutProperty{Clear: true},
		Description:      WindowsShortcutProperty{Value: "Rocket.Chat Portable by Portapps"},
		IconLocation:     WindowsShortcutProperty{Value: Papp.Process},
		WorkingDirectory: WindowsShortcutProperty{Value: Papp.AppPath},
	})
	if err != nil {
		Log.Error("Cannot create shortcut:", err)
	}

	Launch(os.Args[1:])

	_ = os.Remove(shortcutPath)
}
