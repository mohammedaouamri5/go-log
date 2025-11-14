package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/mohammedaouamri5/go-log/log"
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

type __User struct {
	ID    int
	Name  string
	Email string
}

func TestTheme(__func func() *log.Colors, theme string) {

	__user := &__User{
		ID:    1,
		Name:  "mohammed",
		Email: "mohammed.aouamri@pm.me",
	}
	fmt.Println("====================================")
	log.INIT(
		os.Stdout,
		log.GetSimpelFormatterText(__func()),
	)

	log.Info(theme)
	log.Warn(theme)
	log.Error(theme)
	log.Fatal(theme)

	log.WithMap(obj).Info(theme)
	log.WithMap(obj).Warn(theme)
	log.WithMap(obj).Error(theme)
	log.WithMap(obj).Fatal(theme)

	log.WithObj(__user).Info(theme)
	log.WithObj(__user).Warn(theme)
	log.WithObj(__user).Error(theme)
	log.WithObj(__user).Fatal(theme)

}

func main() {
	TestTheme(func() *log.Colors { return nil }, "Default Theme")

	__err := errors.New("some error")
	log.WithErr(__err).Error("Hello World")
}
