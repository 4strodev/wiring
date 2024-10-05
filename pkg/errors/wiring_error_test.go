package errors

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWrapError(t *testing.T) {
	var err = errors.New("some unexpected error")
	wiringError := WrapError(err)
	require.Equal(t, err, wiringError.error)
	require.Equal(t, "some unexpected error", wiringError.error.Error())
	require.Equal(t, "some unexpected error", wiringError.Error())
}

func TestErrorf(t *testing.T) {
	t.Run("should format basic values", func(t *testing.T) {
		wiringError := Errorf("some unexpected error: %s", "some message")
		require.Equal(t, "some unexpected error: some message", wiringError.Error())
	})
	t.Run("should wrap errors like fmt.Errorf", func(t *testing.T) {
		var err = errors.New("internal error")
		wiringError := Errorf("some unexpected error: %w", err)
		require.ErrorIs(t, wiringError, err)
	})
}
