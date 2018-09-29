package main

import (
	"errors"

	"github.com/shabin.hashim/Go_Patterns/factory/impl"
)

var UserNotFoundError = errors.New("User Not Found")

func main() {

	dataStore := impl.CreateDatastore(map[string]string{
		"enginename": "memory",
	})

	dataStore.Name()

}
