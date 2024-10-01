package pkg

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/4strodev/wiring/pkg/errors"
)

type wireContainer struct {
	typeMapping  map[reflect.Type]*dependencySpec
	tokenMapping map[string]*dependencySpec
}

func New() Container {
	return &wireContainer{
		typeMapping:  make(map[reflect.Type]*dependencySpec),
		tokenMapping: make(map[string]*dependencySpec),
	}
}

// SingletonToken implements pkg.Container.
func (w *wireContainer) SingletonToken(token string, resolver any) error {
	spec, err := newSpec(resolver, SINGLETON, w)
	if err != nil {
		return err
	}
	spec.lifeCycle = SINGLETON
	w.tokenMapping[token] = spec
	return nil
}

// TransientToken implements pkg.Container.
func (w *wireContainer) TransientToken(token string, resolver any) error {
	spec, err := newSpec(resolver, TRANSIENT, w)
	if err != nil {
		return err
	}
	spec.lifeCycle = TRANSIENT
	w.tokenMapping[token] = spec
	return nil
}

// Fill implements pkg.Container.
func (w *wireContainer) Fill(structure any) error {
	baseType := reflect.TypeOf(structure)
	baseValue := reflect.ValueOf(structure)
	if baseType.Kind() != reflect.Pointer || baseType.Elem().Kind() != reflect.Struct {
		errors.NewError("fill requires a struct pointer")
	}

	structType := baseType.Elem()
	structValue := baseValue.Elem()
	nFields := structType.NumField()
	for i := 0; i < nFields; i++ {
		var instance any
		var err error
		fieldType := structType.Field(i)
		if !fieldType.IsExported() {
			continue
		}
		fieldValue := structValue.Field(i)

		tagValue := fieldType.Tag.Get(WIRE_TAG)
		tagParams := strings.SplitN(tagValue, ",", 2)
		if len(tagParams) == 2 && "ignore" == tagParams[1] {
			continue
		}

		// Handling token resolved strategy
		var spec *dependencySpec
		if tagParams[0] != "" {
			tokenTag := tagParams[0]
			spec, err := w.getSpecForToken(tokenTag)
			if err != nil {
				return err
			}
			instance, err = spec.Resolve()
			if err != nil {
				return err
			}
			fieldValue.Set(reflect.ValueOf(instance))
			continue
		}

		// Handling type resolving strategy
		spec, err = w.getSpec(fieldType.Type)
		if err != nil {
			err = errors.Errorf("error resolving field '%s': %w", fieldType.Name, err)
			return err
		}
		instance, err = spec.Resolve()
		if err != nil {
			return nil
		}
		fieldValue.Set(reflect.ValueOf(instance))
	}

	return nil
}

// Resolve implements pkg.Container.
func (w *wireContainer) Resolve(abstraction any) error {
	abstractionVal := reflect.ValueOf(abstraction)
	if abstractionVal.Kind() != reflect.Pointer {
		return errors.NewError("abstranction must be a pointer to an interface")
	}
	if !abstractionVal.Elem().CanSet() {
		return errors.NewError("cannot set value of abstraction")
	}

	abstractionType := abstractionVal.Elem().Type()
	var instance any

	spec, err := w.getSpec(abstractionType)
	if err != nil {
		return err
	}
	instance, err = spec.Resolve()
	if err != nil {
		return err
	}

	if !reflect.TypeOf(instance).Implements(abstractionVal.Type().Elem()) {
		return errors.NewError(fmt.Sprintf("worng resolver for type %v", abstractionType))
	}

	abstractionVal.Elem().Set(reflect.ValueOf(instance))
	return nil
}

// ResolveWithToken implements pkg.Container.
func (w *wireContainer) ResolveToken(token string, abstraction any) error {
	abstractionVal := reflect.ValueOf(abstraction)
	if abstractionVal.Kind() != reflect.Pointer {
		return errors.NewError("abstranction must be a pointer to an interface")
	}
	abstractionType := abstractionVal.Elem().Type()
	var instance any

	if !abstractionVal.Elem().CanSet() {
		return errors.NewError("cannot set value of abstraction")
	}

	spec, err := w.getSpecForToken(token)
	if err != nil {
		return err
	}
	instance, err = spec.Resolve()
	if err != nil {
		return err
	}

	if !reflect.TypeOf(instance).Implements(abstractionVal.Type().Elem()) {
		return errors.NewError(fmt.Sprintf("worng resolver for type %v", abstractionType))
	}

	abstractionVal.Elem().Set(reflect.ValueOf(instance))
	return nil
}

// Singleton sets a resolver for the provided type with a singleton lifecycle. If a previous resolver was set
// the new one overrides the previous one.
func (w *wireContainer) Singleton(resolver any) error {
	// May be here we can check if the resolver is valid
	spec, err := newSpec(resolver, SINGLETON, w)
	if err != nil {
		return nil
	}
	spec.lifeCycle = SINGLETON

	intfceType := spec.Type()
	w.typeMapping[intfceType] = spec
	return nil
}

// Singleton sets a resolver for the provided type with a transient lifecycle. If a previous resolver was set
// the new one overrides the previous one.
func (w *wireContainer) Transient(resolver any) error {
	// May be here we can check if the resolver is valid
	spec, err := newSpec(resolver, TRANSIENT, w)
	if err != nil {
		return nil
	}
	spec.lifeCycle = TRANSIENT

	intfceType := spec.Type()
	w.typeMapping[intfceType] = spec
	return nil
}

func (w *wireContainer) getSpecForToken(token string) (*dependencySpec, error) {
	spec, abstractionDefined := w.tokenMapping[token]
	if !abstractionDefined {
		message := fmt.Sprintf("resolver for token '%s' not set", token)
		return &dependencySpec{}, errors.NewError(message)
	}
	return spec, nil
}

func (w *wireContainer) getSpec(reflectType reflect.Type) (*dependencySpec, error) {
	spec, abstractionDefined := w.typeMapping[reflectType]
	if !abstractionDefined {
		message := fmt.Sprintf("resolver for type '%s' not set", reflectType.String())
		return &dependencySpec{}, errors.NewError(message)
	}
	return spec, nil
}

func (w *wireContainer) resolveType(reflectionType reflect.Type) (any, error) {
	spec, exists := w.typeMapping[reflectionType]
	if !exists {
		return nil, errors.Errorf("resolver not defined for %s", reflectionType.String())
	}
	return spec.Resolve()
}
