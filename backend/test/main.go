package main

import (
	"fmt"
	"reflect"
)

type user struct {
	name  string
	email string
}

type errJson struct {
	status string
	error  string
	errors any
}

func main() {
	userData := &user{}
	userData.name = "ayoub"
	errJson := &errJson{}

	validateData(userData, errJson)

	fmt.Printf("errJson: %v\n", errJson)
}

func validateData(userData *user, errJson *errJson) {
	userErrorJson := user{}
	if userData.name == "" {
		userErrorJson.name = "Name is requierd"
	}
	if userData.email == "" {
		userErrorJson.name = "Email is requierd"
	}

	if !reflect.DeepEqual(userErrorJson, user{}) {
		errJson.errors = userErrorJson
	}
}
