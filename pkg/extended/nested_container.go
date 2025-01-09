package extended

import (
	"github.com/4strodev/wiring/pkg"
	"reflect"
)

type NestedContainer struct {
	parent    pkg.Container
	pkg.Container
}

// Fill implements pkg.Container.
func (n *NestedContainer) Fill(structure any) error {
	err := n.Container.Fill(structure)
	if err != nil {
		return n.parent.Fill(structure)
	}

	return nil
}

// HasToken implements pkg.Container.
func (n *NestedContainer) HasToken(token string) bool {
	return n.Container.HasToken(token) || n.parent.HasToken(token)
}

// HasType implements pkg.Container.
func (n *NestedContainer) HasType(refType reflect.Type) bool {
	return n.Container.HasType(refType) || n.parent.HasType(refType)
}

// Resolve implements pkg.Container.
func (n *NestedContainer) Resolve(value any) error {
	err := n.Container.Resolve(value)
	if err != nil {
		return n.parent.Resolve(value)
	}

	return nil
}

// ResolveToken implements pkg.Container.
func (n *NestedContainer) ResolveToken(token string, value any) error {
	err := n.Container.ResolveToken(token, value)
	if err != nil {
		return n.parent.ResolveToken(token, value)
	}

	return nil
}

// Singleton implements pkg.Container.
func (n *NestedContainer) Singleton(resolver any) error {
	return n.Container.Singleton(resolver)
}

// SingletonToken implements pkg.Container.
func (n *NestedContainer) SingletonToken(token string, resolver any) error {
	return n.Container.SingletonToken(token, resolver)
}

// Transient implements pkg.Container.
func (n *NestedContainer) Transient(resolver any) error {
	return n.Container.Transient(resolver)
}

// TransientToken implements pkg.Container.
func (n *NestedContainer) TransientToken(token string, resolver any) error {
	return n.Container.TransientToken(token, resolver)
}

func NewNested(parent pkg.Container) pkg.Container {
	container := pkg.New()
	return &NestedContainer{
		parent:    parent,
		Container: container,
	}
}
