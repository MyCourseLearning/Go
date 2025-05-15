package hello

import (
	"errors"
	"fmt"
)

func SayHello() string {
	fmt.Println(sayHello())
	return "Hello, World!"
}

func sayHello() string {
	return "Hello, World! private!"
}

func ReturnError() error {
	return errors.New("")
}
