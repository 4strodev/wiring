[![go badge](https://pkg.go.dev/badge/github.com/4strodev/wiring.svg)](https://pkg.go.dev/github.com/4strodev/wiring)
![lint status badge](https://github.com/4strodev/wiring/actions/workflows/lint.yaml/badge.svg)
![test status badge](https://github.com/4strodev/wiring/actions/workflows/test.yaml/badge.svg)
# 🔌 Wiring
A container for DI and autowiring written in Go. The main purpose is to provide an interface and a default implementation
for autowiring. You can use the default implementation or create your own if the requirements of your project are not
satisfied.

## Download

    go get github.com/4strodev/wiring@latest
    

# 🌟 Features
Yet this library tries to be as simple as possible I tried to create the minimum features to cover the majority of usecases.

## Lifecycle
There are two kind of lifecycles

- **Singleton**: These dependencies are instantiated once and then the instance is cached for future resolves.
  - The default implementation is **🧵 concurrently safe**.
- **Transient**: Those are dependencies that are always instantiated every time they are resolved.

Ej.
### Resolve a dependency
```go
package main

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

func main() {
	var container = wiring.New()
	container.Singleton(func() (Abstraction, error) {
        	// This resolver is executed just once
		fmt.Println("Running resolver")
		return &Implementation{}, nil
	})

    	// Resolving dependency
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
```

## Extended containers
The `extended` package contains extensions for containers:

- **Derived**: A derived container allows you to create containers that inherits resolvers from parent containers. Allowing you to use container for
  short living contexts like an http request.
- **Must**: An interface that instead of returning errors panics.

## Token based
Token based dependencies allows you to specify dependencies using a custom token. These are performant, because they are not relying in reflection (at all), and allows you to have
multiple dependencies of the same type but with different tokens.

## Struct filling
Do you have massive dependencies? No problem define a struct with exported fields and let the container fill your struct with the dependencies you need.
```go
package main

import (
    	"fmt"
    	"log"

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

type FillableStruct struct {
    	Greeter         Abstraction // if no tag is specified it is resolved by type
    	TokenBased      Abstraction `wire:"token"`
    	IgnoredReader   io.Reader   `wire:",ignore"`
    	ignoredField    string

}

func main() {
	var container = wiring.New()
	container.Singleton(func() (Abstraction, error) {
        	// This resolver is executed just once
		return &Implementation{}, nil
	})

	container.SingletonToken("token", func() (Abstraction, error) {
        	// This resolver is executed just once
		return &Implementation{}, nil
	})

    	var fillable FillableStruct
    	err := container.Fill(&fillable)
    	if err != nil {
		log.Fatal(err)
    	}

    	fillable.Greeter.Greet()
    	fillable.TokenBased.Greet()
}
```

## Docs
There are more examples on [the documentation](https://pkg.go.dev/github.com/4strodev/wiring)

## Inspiration
This project was heavily inspired by [goloby/container](https://github.com/golobby/container).
