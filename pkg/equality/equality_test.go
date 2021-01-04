package equality

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilesAreEqual(t *testing.T) {
	t.Run("Compare json object to itself", func(t *testing.T) {
		var result interface{}
		obj := []byte(`{"id":"aa", "name":"bb"}`)
		err := json.Unmarshal(obj, &result)
		assert.Nil(t, err)

		ok := JsonObjectsEquals(result, result)
		assert.Equal(t, true, ok)
	})

	t.Run("Compare json object to different object", func(t *testing.T) {
		var result1, result2 interface{}
		obj1 := []byte(`{"id":"aa", "name":"bb"}`)
		obj2 := []byte(`{"id":"aa", "name":"cc"}`)

		err := json.Unmarshal(obj1, &result1)
		assert.Nil(t, err)

		err = json.Unmarshal(obj2, &result2)
		assert.Nil(t, err)

		ok := JsonObjectsEquals(result1, result2)
		assert.Equal(t, false, ok)
	})

	t.Run("Compare json object to different type of object", func(t *testing.T) {
		var result1, result2 interface{}
		obj1 := []byte(`{"id":"aa", "name":"bb"}`)
		obj2 := []byte(`[{"id":"aa", "name":"cc"}]`)

		err := json.Unmarshal(obj1, &result1)
		assert.Nil(t, err)

		err = json.Unmarshal(obj2, &result2)
		assert.Nil(t, err)

		ok := JsonObjectsEquals(result1, result2)
		assert.Equal(t, false, ok)
	})

	t.Run("Compare json array itself", func(t *testing.T) {
		var result1 interface{}
		obj1 := []byte(`[{"id":"aa", "name":"bb"}, {"id":"bb", "name":"cc"}]`)

		err := json.Unmarshal(obj1, &result1)
		assert.Nil(t, err)

		ok := JsonObjectsEquals(result1, result1)
		assert.Equal(t, true, ok)
	})

	t.Run("Compare json array to another json array with same elements but different values", func(t *testing.T) {
		var result1, result2 interface{}
		obj1 := []byte(`[{"id":"aa", "name":"bb"}, {"id":"bb", "name":"cc"}]`)
		obj2 := []byte(`[{"id":"aa", "name":"bb"}, {"id":"dd", "name":"ee"}]`)

		err := json.Unmarshal(obj1, &result1)
		assert.Nil(t, err)

		err = json.Unmarshal(obj2, &result2)
		assert.Nil(t, err)

		ok := JsonObjectsEquals(result1, result2)
		assert.Equal(t, false, ok)
	})

	t.Run("Compare json array to another json array with same size but different elements", func(t *testing.T) {
		var result1, result2 interface{}
		obj1 := []byte(`[{"id":"aa", "name":"bb"}, {"id":"bb", "name":"cc"}]`)
		obj2 := []byte(`[{"id":"aa", "name":"bb"}, {"age":"dd", "status":"ee"}]`)

		err := json.Unmarshal(obj1, &result1)
		assert.Nil(t, err)

		err = json.Unmarshal(obj2, &result2)
		assert.Nil(t, err)

		ok := JsonObjectsEquals(result1, result2)
		assert.Equal(t, false, ok)
	})

	t.Run("Compare json array to another json array with different size", func(t *testing.T) {
		var result1, result2 interface{}
		obj1 := []byte(`[{"id":"aa", "name":"bb"}, {"id":"bb", "name":"cc"}]`)
		obj2 := []byte(`[{"id":"aa", "name":"bb"}, {"age":"dd", "status":"ee"}, {"property":"zz"}]`)

		err := json.Unmarshal(obj1, &result1)
		assert.Nil(t, err)

		err = json.Unmarshal(obj2, &result2)
		assert.Nil(t, err)

		ok := JsonObjectsEquals(result1, result2)
		assert.Equal(t, false, ok)
	})

	t.Run("Compare json map itself", func(t *testing.T) {
		var result1 interface{}
		obj1 := []byte(`{"data": {"id":"aa", "name":"bb"}}`)

		err := json.Unmarshal(obj1, &result1)
		assert.Nil(t, err)

		ok := JsonObjectsEquals(result1, result1)
		assert.Equal(t, true, ok)
	})

	t.Run("Compare json map to another json map with same elements but different values", func(t *testing.T) {
		var result1, result2 interface{}
		obj1 := []byte(`{"data": {"id":"aa", "name":"bb"}}`)
		obj2 := []byte(`{"data": {"id":"bb", "name":"cc"}}`)

		err := json.Unmarshal(obj1, &result1)
		assert.Nil(t, err)

		err = json.Unmarshal(obj2, &result2)
		assert.Nil(t, err)

		ok := JsonObjectsEquals(result1, result2)
		assert.Equal(t, false, ok)
	})

	t.Run("Compare json map to another json map with same size but different elements", func(t *testing.T) {
		var result1, result2 interface{}
		obj1 := []byte(`{"data": {"id":"aa", "name":"bb"}}`)
		obj2 := []byte(`{"data": {"age":"bb", "status":"cc"}}`)

		err := json.Unmarshal(obj1, &result1)
		assert.Nil(t, err)

		err = json.Unmarshal(obj2, &result2)
		assert.Nil(t, err)

		ok := JsonObjectsEquals(result1, result2)
		assert.Equal(t, false, ok)
	})

	t.Run("Compare json map to another json map with different size", func(t *testing.T) {
		var result1, result2 interface{}
		obj1 := []byte(`{"data": {"id":"aa", "name":"bb"}}`)
		obj2 := []byte(`{"data": {"id":"bb", "name":"cc"}, "age":"pp"}`)

		err := json.Unmarshal(obj1, &result1)
		assert.Nil(t, err)

		err = json.Unmarshal(obj2, &result2)
		assert.Nil(t, err)

		ok := JsonObjectsEquals(result1, result2)
		assert.Equal(t, false, ok)
	})

}
