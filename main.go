package main

import (
	"fmt"
	"log"

	"github.com/4strodev/wiring/pkg/wiring"
)

type MyInterface interface {
	Greet() string
}

type MyStruct struct {
}

func (s MyStruct) Greet() string {
	return "WTF!"
}

func main() {
	container := wiring.New()
	container.Singleton(func() (MyInterface, error) {
		return MyStruct{}, nil
	})

	var abstraction MyInterface
	err := container.Resolve(&abstraction)
	if err != nil {
		log.Fatal(err)
	}
	var myOhterValue MyInterface
	err = container.Resolve(&myOhterValue)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(abstraction.Greet())
	fmt.Println(myOhterValue.Greet())
}
