package main

import (
	"fmt"
	"github.com/mohammedaouamri5/go-log/log"
	"os"
)

var obj = map[string]interface{}{
	"name":   "mohamed",
	"age":    22,
	"active": true,
	"hobbies": []interface{}{
		"coding", "music", "gaming", 13, 3.37, true,
	},
	"data": []byte{1, 2, 3, 4, 5},
	"profile": map[string]interface{}{
		"lang": "go",
		"os":   "arch btw",
	},
}

func TestTheme(__func func() *log.Colors, theme string) {
	Logger := log.Logger{
		Writer:    os.Stdout,
		Formatter: log.GetSimpelFormatterText("/home/mohammedaouamri/DEV/go-log/", __func()),
	}
	fmt.Println("====================================")
	Logger.Info(theme)
	Logger.Warn(theme)
	Logger.Error(theme)
	Logger.WithMap(obj).Info("Hello World")
}

func main() {
	TestTheme(log.MonokaiTheme, "Monokai Theme")
	TestTheme(log.DraculaTheme, "Dracula Theme")
	TestTheme(log.GruvboxTheme, "Gruvbox Theme")
	TestTheme(log.SolarizedDarkTheme, "Solarized Dark Theme")
	TestTheme(log.NordTheme, "Nord Theme")
	TestTheme(log.OneDarkTheme, "One Dark Theme")
	TestTheme(log.TokyoNightTheme, "Tokyo Night Theme")
	TestTheme(log.AyuDarkTheme, "Ayu Dark Theme")
	TestTheme(log.LightTheme, "Light Theme")
}

