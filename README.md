# Wiring
A container for DI and autowiring written in Go. The main purpose is to provide an interface and a default implementation
for autowiring. You can use the default implementation or create your own if the requirements of your project are not
satisfied.

## Download

    go get github.com/4strodev/wiring@latest

## Usage
`Container` is the base interface, it provides all the methods necessary to start using DI in your project.

1. Instantiate a container. To do this wiring provides a default implementation defined in the `pkg` package.

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

### Fill a struct
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
    Greeter         Abstraction
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

## Thanks
This project was heavily inspired by [goloby/container](https://github.com/golobby/container).
