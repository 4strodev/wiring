package pkg

import (
	"io"
	"testing"

	"github.com/4strodev/wiring/pkg/internal/mocks"
	"github.com/stretchr/testify/require"
)

func InitializeContainer(t *testing.T) Container {
	var container = New()
	err := container.Singleton(mocks.Resolver)
	require.NoError(t, err)
	err = container.SingletonToken(mocks.TESTING_TOKEN, mocks.TokenResolver)
	require.NoError(t, err)
	return container
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
}
