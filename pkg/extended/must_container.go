package extended

import (
	"reflect"
)

// MustContainer is a container which instead of returing errors it panics
// use it for prototyping or in circumstances where you know that resolvers
// are extremerly well configured.
type MustContainer interface {
	// Type based injection

	// Singleton sets a dependency as a [wiring.] dependency.
	// Once the abstraction is instanciated this instance will be cached and
	// will no longer create new instances
	Singleton(resolver any)
	// Transient sets a dependency as a transient dependency.
	// Every time the container is asked to resolve an abstraction
	// the container will create a new instance of that dependency
	Transient(resolver any)
	// Resolve given a pointer to value it will be resolved and the container
	// will update the referenced value with the instance resolved
	Resolve(value any)

	// Token based injection

	// SingletonToken same as Singleton but instead of using the type to identify
	// the implementation it uses the token
	SingletonToken(token string, resolver any)
	// TransientToken same as Transient but instead of using the type to identify
	// the implementation it uses the token
	TransientToken(token string, resolver any)
	// Gets the instance associated with the provided token
	ResolveToken(token string, value any)

	// Fill gets a struct pointer and resolves their fields, if the field needs to be resolved by token
	// you can use the 'wire' tag with the token that is associated with. If the field needs to be ignored
	// use the ignore param -> wire:",ignore". Unexported fields will be ignored
	Fill(structure any)

	// Check if the container has a resolver for that type
	HasType(refType reflect.Type) bool
	// Check if the container has a resolver for that token
	HasToken(token string) bool
}

