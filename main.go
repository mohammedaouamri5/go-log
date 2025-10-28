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

	fmt.Println("====================================")
	log.INIT(
		os.Stdout,
		log.GetSimpelFormatterText(__func()),
	)
}

func main() {
	TestTheme(func() *log.Colors { return nil }, "Default Theme")

	__err := errors.New("some error")
	log.Error(__err.Error())
	log.WithErr(__err).Error("Hello World")

}
