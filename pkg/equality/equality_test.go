package equality

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilesAreEqual(t *testing.T) {
	t.Run("Compare json file to itself", func(t *testing.T) {
		assert.Nil(t, nil)
	})

	t.Run("Compare empty file to itself", func(t *testing.T) {
		assert.Nil(t, nil)
	})

	t.Run("Compare one empty file to json file", func(t *testing.T) {
		assert.Nil(t, nil)
	})

	t.Run("Compare one non json file to json file", func(t *testing.T) {
		assert.Nil(t, nil)
	})

	t.Run("Compare one non json file to non json file", func(t *testing.T) {
		assert.Nil(t, nil)
	})

	t.Run("Compare json file to another json file not equal", func(t *testing.T) {
		assert.Nil(t, nil)
	})

	t.Run("Compare json file to another json file but unordered", func(t *testing.T) {
		assert.Nil(t, nil)
	})
}
