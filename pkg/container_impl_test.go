package pkg

import (
	"testing"

	"github.com/4strodev/wiring/pkg/internal/mocks"
	"github.com/stretchr/testify/require"
)

// The resolver of a singleton should be called just once
func TestFill(t *testing.T) {
	var container = New()
	err := container.Singleton(mocks.Resolver)
	require.NoError(t, err)
	err = container.SingletonToken(mocks.TESTING_TOKEN, mocks.TokenResolver)
	require.NoError(t, err)

	var fillableStruct mocks.FillableStruct
	err = container.Fill(&fillableStruct)
	require.NoError(t, err)

	// Validate struct fields
	require.NoError(t, fillableStruct.CheckResolvedFields())
	require.Equal(t, mocks.TokenResolver(), fillableStruct.TokenResolved)
	require.Equal(t, mocks.Resolver(), fillableStruct.TypeResolved)
}
