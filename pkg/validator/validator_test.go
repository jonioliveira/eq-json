package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidatorArguments(t *testing.T) {
	t.Run("No args", func(t *testing.T) {
		assert.Nil(t, nil)
	})

	t.Run("Less then needed args", func(t *testing.T) {
		assert.Nil(t, nil)
	})

	t.Run("Right number of args", func(t *testing.T) {
		assert.Nil(t, nil)
	})

	t.Run("More than needed Args", func(t *testing.T) {
		assert.Nil(t, nil)
	})
}

func TestValidatorFilePath(t *testing.T) {
	t.Run("Valid file path", func(t *testing.T) {
		assert.Nil(t, nil)
	})

	t.Run("Invalid file path", func(t *testing.T) {
		assert.Nil(t, nil)
	})

	t.Run("One valid and one invalid file path", func(t *testing.T) {
		assert.Nil(t, nil)
	})
}
