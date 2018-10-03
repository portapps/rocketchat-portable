//go:generate go install -v github.com/kevinburke/go-bindata/go-bindata
//go:generate go-bindata -pkg assets -o assets/assets.go res/Rocket.Chat.lnk
//go:generate go install -v github.com/josephspurrier/goversioninfo/cmd/goversioninfo
//go:generate goversioninfo -icon=res/papp.ico
package main

import (
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
	Papp.DataPath = AppPathJoin("data")
	Papp.Process = PathJoin(Papp.AppPath, "Rocket.Chat.exe")
	Papp.Args = nil
	Papp.WorkingDir = Papp.AppPath

	fo, err := os.Create(path.Join(Papp.DataPath, "update.json"))
	if err != nil {
		Log.Error("Cannot create update.json:", err)
	}
	defer fo.Close()

	_, err = io.Copy(fo, strings.NewReader(`{
  "autoUpdate": false
}`))
	if err != nil {
		Log.Error("Cannot write to update.json:", err)
	}

	shortcutPath := path.Join(os.Getenv("APPDATA"), "Microsoft", "Windows", "Start Menu", "Programs", "Rocket.Chat Portable.lnk")

	// Copy default shortcut
	defaultShortcut, err := assets.Asset("res/Rocket.Chat.lnk")
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
		Description:      "Rocket.Chat Portable by Portapps",
		IconLocation:     Papp.Process,
		WorkingDirectory: Papp.AppPath,
	})
	if err != nil {
		Log.Error("Cannot create shortcut:", err)
	}

	Launch(os.Args[1:])
}
