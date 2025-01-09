package extended

import (
	"github.com/4strodev/wiring/pkg"
	"reflect"
)

type mustContainer struct {
	pkg.Container
}

// Fill implements MustContainer.
func (m *mustContainer) Fill(structure any) {
	err := m.Container.Fill(structure)
	if err != nil {
		panic(err)
	}
}

// HasToken implements MustContainer.
func (m *mustContainer) HasToken(token string) bool {
	return m.Container.HasToken(token)
}

// HasType implements MustContainer.
func (m *mustContainer) HasType(refType reflect.Type) bool {
	return m.Container.HasType(refType)
}

// Resolve implements MustContainer.
func (m *mustContainer) Resolve(value any) {
	err := m.Container.Resolve(value)
	if err != nil {
		panic(err)
	}
}

// ResolveToken implements MustContainer.
func (m *mustContainer) ResolveToken(token string, value any) {
	err := m.Container.ResolveToken(token, value)
	if err != nil {
		panic(err)
	}
}

// Singleton implements MustContainer.
func (m *mustContainer) Singleton(resolver any) {
	err := m.Container.Singleton(resolver)
	if err != nil {
		panic(err)
	}
}

// SingletonToken implements MustContainer.
func (m *mustContainer) SingletonToken(token string, resolver any) {
	err := m.Container.SingletonToken(token, resolver)
	if err != nil {
		panic(err)
	}
}

// Transient implements MustContainer.
func (m *mustContainer) Transient(resolver any) {
	err := m.Container.Transient(resolver)
	if err != nil {
		panic(err)
	}
}

// TransientToken implements MustContainer.
func (m *mustContainer) TransientToken(token string, resolver any) {
	err := m.Container.TransientToken(token, resolver)
	if err != nil {
		panic(err)
	}
}

func Must(container pkg.Container) MustContainer {

	return &mustContainer{}
}
