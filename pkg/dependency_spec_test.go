package pkg

import (
	"reflect"
	"testing"

	"github.com/4strodev/wiring/pkg/internal/mocks"
	"github.com/stretchr/testify/require"
)

func TestType(t *testing.T) {
	spec, err := newSpec(mocks.ResolverWithMessage("Hello world"), SINGLETON, nil)
	require.NoError(t, err)
	require.Equal(t, reflect.TypeFor[mocks.Abstraction](), spec.Type())
}

func TestResolve(t *testing.T) {
	t.Run("should execute resolver just once on singleton", func(t *testing.T) {
		var counter = 0
		resolver := func() mocks.Abstraction {
			counter++
			return &mocks.Implementation{
				Message: mocks.DEFAULT_MESSAGE,
			}
		}
		container := New()
		spec, err := newSpec(resolver, SINGLETON, container.(*wireContainer))
		require.NoError(t, err)
		require.NotNil(t, spec)

		for i := 0; i < 5; i++ {
			instance, err := spec.Resolve()
			require.NoError(t, err)
			abstraction, ok := instance.(mocks.Abstraction)
			require.True(t, ok)
			_, ok = abstraction.(*mocks.Implementation)
			require.True(t, ok)
		}
		require.Equal(t, 1, counter)
	})
	t.Run("should execute resolver every time an abstraction is resolved on transient", func(t *testing.T) {
		var counter = 0
		resolver := func() mocks.Abstraction {
			counter++
			return &mocks.Implementation{
				Message: mocks.DEFAULT_MESSAGE,
			}
		}
		container := New()
		spec, err := newSpec(resolver, TRANSIENT, container.(*wireContainer))
		require.NoError(t, err)
		require.NotNil(t, spec)

		for i := 0; i < 5; i++ {
			instance, err := spec.Resolve()
			require.NoError(t, err)
			abstraction, ok := instance.(mocks.Abstraction)
			require.True(t, ok)
			_, ok = abstraction.(*mocks.Implementation)
			require.True(t, ok)
		}
		require.Equal(t, 5, counter)
	})
}

func TestExecuteResolver(t *testing.T) {
	container := New()
	spec, err := newSpec(mocks.Resolver, TRANSIENT, container.(*wireContainer))
	require.NoError(t, err)
	require.NotNil(t, spec)

	instance, err := spec.executeResolver()
	require.NoError(t, err)
	abstraction, ok := instance.(mocks.Abstraction)
	require.True(t, ok)
	_, ok = abstraction.(*mocks.Implementation)
	require.True(t, ok)
}