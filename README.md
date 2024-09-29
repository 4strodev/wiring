# Wiring
A container for DI and autowiring written in Go. The main purpose is to provide an interface and a default implementation
for autowiring. You can use the default implementation or create your own if the requirements of your project are not
satisfied.

## Download

    go get github.com/4strodev/wiring@latest

## Usage
`Container` is the base interface, it provides all the methods necessary to start using DI in your project.

1. Instantiate a container. To do this wiring provides a default implementation defined in the `impl/wiring` package.
```go
// TODO obsolete example
package main

import (
	"fmt"
	"reflect"

	"github.com/4strodev/wiring/impl/wiring"
	"github.com/4strodev/wiring/pkg"
)

type MyInterface interface {
	MyMethod()
}

type MyInterfaceImpl struct {
}

func (impl MyInterfaceImpl) MyMethod() {
	fmt.Println("Hello")
}

type MyStruct struct {
	MyValue      MyInterface
	MyOtherValue MyInterface `wire:"ignore"`
}

func main() {
	var container pkg.Container = wiring.New()
	var myInterfaceInstance MyInterface
	var myStructInstance MyStruct

	var resolver = pkg.FuncResolver[MyInterfaceImpl](func(pkg.ReadContainer) (MyInterfaceImpl, error) {
		return MyInterfaceImpl{}, nil
	})

	container.Singleton(reflect.TypeFor[MyInterface](), resolver)

	err := container.Resolve(&myInterfaceInstance)
	if err != nil {
		panic(err)
	}

	myInterfaceInstance.MyMethod()
	container.Fill(&myStructInstance)
	myStructInstance.MyValue.MyMethod()
}
```
