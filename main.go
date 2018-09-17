//go:generate go install -v github.com/josephspurrier/goversioninfo/cmd/goversioninfo
//go:generate goversioninfo -icon=res/papp.ico
package main

import (
	"io"
	"os"
	"path"
	"strings"

	. "github.com/portapps/portapps"
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

	Launch(os.Args[1:])
}
