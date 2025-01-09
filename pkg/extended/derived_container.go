package extended

import (
	"github.com/4strodev/wiring/pkg"
	"reflect"
)

// DerivedContainer allows you to create containers that inherits resolvers from parent containers.
// This is a fully functional container which can has their own and Scoped 🚀 dependencies.
// Allowing you to use container for short living contexts like an http request.
type DerivedContainer struct {
	parent pkg.Container
	pkg.Container
}

// Fill implements pkg.Container.
func (d *DerivedContainer) Fill(structure any) error {
	err := d.Container.Fill(structure)
	if err != nil {
		return d.parent.Fill(structure)
	}

	return nil
}

// HasToken implements pkg.Container.
func (d *DerivedContainer) HasToken(token string) bool {
	return d.Container.HasToken(token) || d.parent.HasToken(token)
}

// HasType implements pkg.Container.
func (d *DerivedContainer) HasType(refType reflect.Type) bool {
	return d.Container.HasType(refType) || d.parent.HasType(refType)
}

// Resolve implements pkg.Container.
func (d *DerivedContainer) Resolve(value any) error {
	err := d.Container.Resolve(value)
	if err != nil {
		return d.parent.Resolve(value)
	}

	return nil
}

// ResolveToken implements pkg.Container.
func (d *DerivedContainer) ResolveToken(token string, value any) error {
	err := d.Container.ResolveToken(token, value)
	if err != nil {
		return d.parent.ResolveToken(token, value)
	}

	return nil
}

func Derived(parent pkg.Container) pkg.Container {
	container := pkg.New()
	return &DerivedContainer{
		parent:    parent,
		Container: container,
	}
}
