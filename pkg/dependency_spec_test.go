package pkg

import (
	"reflect"
	"sync"
	"sync/atomic"
	"testing"

	"github.com/4strodev/wiring/pkg/internal/mocks"
	"github.com/stretchr/testify/require"
)

func TestType(t *testing.T) {
	spec, err := newSpec(mocks.ResolverWithMessage("Hello world"), SINGLETON, nil)
	require.NoError(t, err)
	require.Equal(t, reflect.TypeFor[mocks.Abstraction](), spec.Type())
}

func TestSpecResolve(t *testing.T) {
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
	t.Run("should be concurrent save for singleton dependencies", func(t *testing.T) {
		var waitGroup sync.WaitGroup
		var counter atomic.Int32
		resolver := func() mocks.Abstraction {
			counter.Add(1)
			return &mocks.Implementation{
				Message: mocks.DEFAULT_MESSAGE,
			}
		}
		container := New()
		spec, err := newSpec(resolver, SINGLETON, container.(*wireContainer))
		require.NoError(t, err)
		require.NotNil(t, spec)

		for i := 0; i < 5; i++ {
			waitGroup.Add(1)
			go func() {
				defer waitGroup.Done()
				instance, err := spec.Resolve()
				require.NoError(t, err)
				abstraction, ok := instance.(mocks.Abstraction)
				require.True(t, ok)
				_, ok = abstraction.(*mocks.Implementation)
				require.True(t, ok)
			}()
		}

		waitGroup.Wait()
		require.Equal(t, int32(1), counter.Load())
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
