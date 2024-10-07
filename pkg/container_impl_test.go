package pkg

import (
	"io"
	"reflect"
	"testing"

	"github.com/4strodev/wiring/pkg/internal/mocks"
	"github.com/stretchr/testify/require"
)

type ComplexAbstraction interface {
	mocks.Abstraction
}

type ComplexImplementation struct {
	mocks.Abstraction
}

func InitializeContainer(t *testing.T) Container {
	var container = New()
	err := container.Singleton(mocks.Resolver)
	// this should make tests fail
	require.Error(t, err)
	err = container.SingletonToken(mocks.TESTING_TOKEN, mocks.TokenResolver)
	require.NoError(t, err)
	return container
}

func TestHasType(t *testing.T) {
	container := InitializeContainer(t)
	require.True(t, container.HasType(reflect.TypeFor[mocks.Abstraction]()))
	require.False(t, container.HasType(reflect.TypeFor[mocks.Implementation]()))
	require.False(t, container.HasType(reflect.TypeFor[io.Reader]()))
	require.False(t, container.HasType(reflect.TypeFor[*mocks.Abstraction]()))
}

func TestHasToken(t *testing.T) {
	container := InitializeContainer(t)
	require.True(t, container.HasToken(mocks.TESTING_TOKEN))
	require.False(t, container.HasToken("another token"))
}

func TestResolve(t *testing.T) {
	t.Run("should return error with no declared abstractions", func(t *testing.T) {
		var err error
		container := New()
		var nonDeclaredAbstraction io.Reader
		err = container.Resolve(&nonDeclaredAbstraction)
		require.Error(t, err)
	})
	t.Run("should accept only pointers", func(t *testing.T) {
		var err error
		var container = InitializeContainer(t)
		var abstraction mocks.Abstraction
		err = container.Resolve(abstraction)
		require.Error(t, err)
	})
	t.Run("should execute resolver passing corresponding arguments", func(t *testing.T) {
		var err error
		var container = InitializeContainer(t)
		err = container.Singleton(func(abstraction mocks.Abstraction) ComplexAbstraction {
			require.NotNil(t, abstraction)
			return ComplexImplementation{Abstraction: abstraction}
		})
		require.NoError(t, err)

		var complexAbstraction ComplexAbstraction
		err = container.Resolve(&complexAbstraction)
		require.NoError(t, err)
		require.NotNil(t, complexAbstraction)
	})
	t.Run("should return error when resolver expects non existing abstraction", func(t *testing.T) {
		var err error
		var container = New()
		err = container.Singleton(func(abstraction mocks.Abstraction) ComplexAbstraction {
			require.NotNil(t, abstraction)
			return ComplexImplementation{Abstraction: abstraction}
		})
		require.NoError(t, err)

		var complexAbstraction ComplexAbstraction
		err = container.Resolve(&complexAbstraction)
		require.Error(t, err)
	})
}

func TestResolveToken(t *testing.T) {
	t.Run("should resolve a dependency by token", func(t *testing.T) {
		var err error
		var container = InitializeContainer(t)
		var value string
		err = container.ResolveToken(mocks.TESTING_TOKEN, &value)
		require.NoError(t, err)
		require.NotZero(t, value)
	})
	t.Run("should return error for not defined token", func(t *testing.T) {
		var err error
		var container = InitializeContainer(t)
		var value string
		err = container.ResolveToken("invalid token", &value)
		require.Error(t, err)
		require.Zero(t, value)
	})
}

// The resolver of a singleton should be called just once
func TestFill(t *testing.T) {
	t.Run("should fill a valid struct", func(t *testing.T) {
		var err error
		container := InitializeContainer(t)

		var fillableStruct mocks.FillableStruct
		err = container.Fill(&fillableStruct)
		require.NoError(t, err)

		// Validate struct fields
		require.NoError(t, fillableStruct.CheckResolvedFields())
		require.Equal(t, mocks.TokenResolver(), fillableStruct.TokenResolved)
		require.Equal(t, mocks.Resolver(), fillableStruct.TypeResolved)
	})
	t.Run("should only accept struct pointers", func(t *testing.T) {
		var err error
		container := InitializeContainer(t)
		var nonStructValue int
		err = container.Fill(&nonStructValue)
		require.Error(t, err)
		err = container.Fill(nonStructValue)
		require.Error(t, err)
	})
	t.Run("should return error when a field has an invalid type or token", func(t *testing.T) {
		var err error
		container := InitializeContainer(t)

		var fillableStruct mocks.InvalidStruct
		err = container.Fill(&fillableStruct)
		require.Error(t, err)
	})
}
