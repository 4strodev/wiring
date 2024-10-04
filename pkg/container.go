// contains the wire implementation of interfaces defined on pkg
package pkg

import "reflect"

const WIRE_TAG = "wire"

type Container interface {
	// Type based injection

	// Singleton sets a dependency as a [wiring.] dependency.
	// Once the abstraction is instanciated this instance will be cached and
	// will no longer create new instances
	Singleton(resolver any) error
	// Transient sets a dependency as a transient dependency.
	// Every time the container is asked to resolve an abstraction
	// the container will create a new instance of that dependency
	Transient(resolver any) error
	// Resolve given a pointer to value it will be resolved and the container
	// will update the referenced value with the instance resolved
	Resolve(value any) error

	// Token based injection

	// SingletonToken same as Singleton but instead of using the type to identify
	// the implementation it uses the token
	SingletonToken(token string, resolver any) error
	// TransientToken same as Transient but instead of using the type to identify
	// the implementation it uses the token
	TransientToken(token string, resolver any) error
	// Gets the instance associated with the provided token
	ResolveToken(token string, value any) error

	// Other container utilities
	Fill(structure any) error

	// Check if the container has a resolver for that type
	HasType(refType reflect.Type) bool
	// Check if the container has a resolver for that token
	HasToken(token string) bool
}
