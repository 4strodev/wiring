package pkg_test

import (
	"fmt"

	wiring "github.com/4strodev/wiring/pkg"
)

type Abstraction interface {
	Greet()
}

type Implementation struct {
}

func (i *Implementation) Greet() {
	fmt.Println("Hello world")
}

func ExmpleNew() {
	var container = wiring.New()
	container.Singleton(func() (Abstraction, error) {
		return &Implementation{}, nil
	})
	var impl Implementation
	container.Resolve(&impl)
	impl.Greet()
	// Output: Hello world
}

func ExampleContainer_ResolveToken() {

}

func ExampleContainer_Fill() {

}
