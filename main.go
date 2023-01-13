package main

import (
	"fmt"

	"github.com/Scholak/go-json-validation/utils"
)

type User struct {
	ID       int64  `json:"id" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	Age      int64  `json:"age" validate:"required,min=0,max=256"`
}

func main() {
	// user1 := &User{1, "sebahattin", "scholaksebahattin@gmail.com", "password", 21}
	user2 := &User{Age: -1}

	errors := utils.ValidateStruct(user2)

	if errors != nil {
		for _, err := range errors {
			fmt.Println(err)
		}
	} else {
		fmt.Println("there is no validation error")
	}
}
