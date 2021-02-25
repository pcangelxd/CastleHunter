package main

import (
	"./application"
	"os"
	"strings"
)

var (
	path, _ = os.Getwd()
	configPath = strings.Replace(path, "\\", "/", -1) + "/config/config.toml"
)

func main()  {
	app := application.NewApplication(configPath)
	app.Run()
}
