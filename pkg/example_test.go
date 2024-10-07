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

type StructFill struct {
	Field string `wire:",ignore"`
	Impl  Abstraction
}

func Example() {
	var container = wiring.New()
	err := container.Singleton(func() (Abstraction, error) {
		return &Implementation{}, nil
	})
	if err != nil {
		panic(err) 
	}
	var impl Abstraction
	err = container.Resolve(&impl)
	if err != nil {
		panic(err)
	}
	impl.Greet()
	// Output: Hello world
}

func ExampleContainer_transient() {
	var container = wiring.New()
	err := container.Transient(func() (Abstraction, error) {
		// The resolver is executed every time the abstraction is resolved
		fmt.Println("Running resolver")
		return &Implementation{}, nil
	})
	if err != nil {
		panic(err)
	}
	var impl Abstraction
	err = container.Resolve(&impl)
	if err != nil {
		panic(err)
	}
	impl.Greet()

	var otherImpl Abstraction
	err = container.Resolve(&otherImpl)
	if err != nil {
		panic(err)
	}
	otherImpl.Greet()
	// Output:
	// Running resolver
	// Hello world
	// Running resolver
	// Hello world
}

func ExampleContainer_singleton() {
	var container = wiring.New()
	err := container.Singleton(func() (Abstraction, error) {
		fmt.Println("Running resolver")
		return &Implementation{}, nil
	})
	if err != nil {
		panic(err)
	}

	var impl Abstraction
	err = container.Resolve(&impl)
	if err != nil {
		panic(err)
	}
	impl.Greet()

	var otherImpl Abstraction
	err = container.Resolve(&otherImpl)
	if err != nil {
		panic(err)
	}
	otherImpl.Greet()
	// Output:
	// Running resolver
	// Hello world
	// Hello world
}

func ExampleContainer_token() {
	var container = wiring.New()
	err := container.TransientToken("abstraction", func() (Abstraction, error) {
		return &Implementation{}, nil
	})
	if err != nil {
		panic(err)
	}
	var impl Abstraction
	err = container.ResolveToken("abstraction", &impl)
	if err != nil {
		panic(err)
	}
	impl.Greet()
	// Output:
	// Hello world
}

func ExampleContainer_fill() {
	var container = wiring.New()
	err := container.Transient(func() (Abstraction, error) {
		return &Implementation{}, nil
	})
	if err != nil {
		panic(err)
	}
	var impl StructFill
	err = container.Fill(&impl)
	if err != nil {
		panic(err)
	}
	impl.Impl.Greet()
	// Output: Hello world
}
