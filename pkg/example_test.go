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
	Field string `wire:"ignore"`
	Impl  Implementation
}

func Example() {
	var container = wiring.New()
	container.Singleton(func() (Abstraction, error) {
		return &Implementation{}, nil
	})
	var impl Implementation
	container.Resolve(&impl)
	impl.Greet()
	// Output: Hello world
}

func ExampleContainer_transient() {
	var container = wiring.New()
	container.Transient(func() (Abstraction, error) {
		// The resolver is executed every time the abstraction is resolved
		fmt.Println("Running resolver")
		return &Implementation{}, nil
	})
	var impl Abstraction
	container.Resolve(&impl)
	impl.Greet()

	var otherImpl Abstraction
	container.Resolve(&otherImpl)
	otherImpl.Greet()
	// Output:
	// Running resolver
	// Hello world
	// Running resolver
	// Hello world
}

func ExampleContainer_singleton() {
	var container = wiring.New()
	container.Singleton(func() (Abstraction, error) {
		// The resolver is executed every time the abstraction is resolved
		fmt.Println("Running resolver")
		return &Implementation{}, nil
	})
	var impl Abstraction
	container.Resolve(&impl)
	impl.Greet()

	var otherImpl Abstraction
	container.Resolve(&otherImpl)
	otherImpl.Greet()
	// Output:
	// Running resolver
	// Hello world
	// Hello world
}

func ExampleContainer_token() {
	var container = wiring.New()
	container.TransientToken("abstraction", func() (Abstraction, error) {
		return &Implementation{}, nil
	})
	var impl Abstraction
	container.ResolveToken("abstraction", &impl)
	impl.Greet()
	// Output:
	// Hello world
}

func ExampleContainer_fill() {
	var container = wiring.New()
	container.Transient(func() (Abstraction, error) {
		return &Implementation{}, nil
	})
	var impl StructFill
	container.Fill(&impl)
	impl.Impl.Greet()
	// Output: Hello world
}
