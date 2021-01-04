package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidatorArguments(t *testing.T) {
	t.Run("No args", func(t *testing.T) {
		ok := ValidateArgumentsSize(0)
		assert.Equal(t, false, ok)
	})

	t.Run("Less then needed args", func(t *testing.T) {
		ok := ValidateArgumentsSize(1)
		assert.Equal(t, false, ok)
	})

	t.Run("Right number of args", func(t *testing.T) {
		ok := ValidateArgumentsSize(2)
		assert.Equal(t, true, ok)
	})

	t.Run("More than needed Args", func(t *testing.T) {
		ok := ValidateArgumentsSize(3)
		assert.Equal(t, false, ok)
	})
}

func TestValidatorFilePath(t *testing.T) {
	t.Run("Valid file path", func(t *testing.T) {
		ok, path := ValidateFilesPath([]string{"../../assets/file1.json"})
		assert.Equal(t, true, ok)
		assert.Equal(t, "", path)
	})

	t.Run("Invalid file path", func(t *testing.T) {
		ok, path := ValidateFilesPath([]string{"../../assets/notExists.json"})
		assert.Equal(t, false, ok)
		assert.Equal(t, "../../assets/notExists.json", path)
	})

	t.Run("One valid and one invalid file path", func(t *testing.T) {
		ok, path := ValidateFilesPath([]string{"../../assets/file1.json", "../../assets/notExists.json"})
		assert.Equal(t, false, ok)
		assert.Equal(t, path, "../../assets/notExists.json")
	})
}
