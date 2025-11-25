package main

import (
	"os"

	"github.com/mohammedaouamri5/go-log/log"
)
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

func main() {
	log.INIT(
		os.Stdout,
		log.NewStructuredLoggerFormatter(nil),
	)	
  	__user := &__User{
		ID:    67,
		Name:  "mohammed",
		Email: "mohammed.aouamri5@email.me",
		pwd:   "9/11",
	}

	log.WithField("User", __user).Info("Logging User")
} 
