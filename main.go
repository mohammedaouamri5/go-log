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
	pwd string
}
func (__user *__User) ToMap() map[string]any {
	return map[string]any{
		"ID":    __user.ID,
		"Name":  __user.Name,
		"Email": __user.Email,
		"pwd": __user.pwd,
	}
}

func TestTheme(__func func() *log.Colors, theme string) {

	__user := &__User{
		ID:    1,
		Name:  "mohammed",
		Email: "mohammed.aouamri@pm.me",
		pwd: "1234",
	}
	fmt.Println("====================================")
	log.INIT(
		os.Stdout,
		log.NewStructuredLoggerFormatter(__func()),
	)


	__err := errors.New("some error")

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
	

	log.WithErr(__err).Info(theme)
	log.WithErr(__err).Warn(theme)
	log.WithErr(__err).Error(theme)
	log.WithErr(__err).Fatal(theme)
	

	log.WithField("Bo3kzjfv" , 938).WithField("User", __user).Info(theme)
	log.WithField("Bo3kzjfv" , 938).WithField("User", *__user).Info(theme)
}

func main() {
	TestTheme(func() *log.Colors { return nil }, "Default Theme")

	__err := errors.New("some error")
	log.WithErr(__err).Error("Hello World")
}
