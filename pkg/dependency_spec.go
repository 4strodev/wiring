package pkg

import (
	"log"
	"reflect"

	"github.com/4strodev/wiring/pkg/errors"
)

type abstractionLifeCycle uint8

const (
	SINGLETON abstractionLifeCycle = iota
	TRANSIENT
)

// dependencySpec defines how the abstraction is resolved
type dependencySpec struct {
	container  *wireContainer
	lifeCycle  abstractionLifeCycle
	resolver   any
	instance   any
	returnType reflect.Type
}

func (spec *dependencySpec) Type() reflect.Type {
	return spec.returnType
}

func (spec *dependencySpec) Resolve() (any, error) {
	switch spec.lifeCycle {
	case SINGLETON:
		if spec.instance == nil {
			instance, err := spec.executeResolver()
			if err != nil {
				return nil, err
			}
			if instance == nil {
				log.Println("instance is nil!")
				return nil, errors.NewError("Resolver returned a nil instance")
			}
			spec.instance = instance
		}

		return spec.instance, nil
	case TRANSIENT:
		return spec.executeResolver()
	default:
		return nil, errors.Errorf("abstraction lifecycle not valid")
	}

}

func (spec *dependencySpec) executeResolver() (any, error) {
	resolverArguments, err := spec.arguments()
	if err != nil {
		return nil, err
	}

	resolverValue := reflect.ValueOf(spec.resolver)
	returnedValues := resolverValue.Call(resolverArguments)

	instanceValue := returnedValues[0]
	instance := instanceValue.Interface()

	if len(returnedValues) == 2 {
		errValue := returnedValues[1]
		if errValue.IsNil() {
			err = nil
		} else {
			err = errValue.Interface().(error)
		}
	}

	return instance, err
}

func (spec *dependencySpec) arguments() ([]reflect.Value, error) {
	resolverType := reflect.TypeOf(spec.resolver)
	values := make([]reflect.Value, resolverType.NumIn())

	for i := 0; i < resolverType.NumIn(); i++ {
		value, err := spec.container.resolveType(resolverType.In(i))
		if err != nil {
			return nil, err
		}
		values[i] = reflect.ValueOf(value)
	}
	return values, nil
}

func newSpec(resolver any, lifeCycle abstractionLifeCycle, container *wireContainer) (spec *dependencySpec, err error) {
	spec = new(dependencySpec)
	spec.lifeCycle = lifeCycle
	spec.container = container

	// Get return type of the function
	resolverType := reflect.TypeOf(resolver)
	if resolverType.Kind() != reflect.Func {
		err = errors.NewError("resolver not valid it should be a function")
		return
	}

	numOut := resolverType.NumOut()
	if numOut < 1 || numOut > 2 {
		err = errors.NewError("resolver should return between 1-2")
		return
	}

	returnType := resolverType.Out(0)
	spec.returnType = returnType
	if returnType.Implements(reflect.TypeFor[error]()) {
		err = errors.NewError("error cannot be the first resolver return type")
		return
	}
	spec.resolver = resolver

	if resolverType.NumOut() == 2 {
		// Check if second return type is error
		secondType := resolverType.Out(1)
		if !secondType.Implements(reflect.TypeFor[error]()) {
			err = errors.NewError("second return type of resolver is not an error")
			return
		}
	}

	return spec, nil
}
