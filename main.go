package main

import (
	"fmt"
	"time"
)

type User struct {
	FirstName string
	LastName  string
	BirthDate time.Time
}

func (user *User) printFullName() {
	fmt.Println(user.FirstName + " " + user.LastName)
}
func (user *User) getFullName() string {
	return user.FirstName + " " + user.LastName
}

func main() {

	user := User{
		FirstName: "Sasha",
		LastName:  "Ver",
		BirthDate: time.Time{},
	}
	fmt.Println(user)

	user.printFullName()

}
