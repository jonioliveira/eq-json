package handler

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandlerReadFile(t *testing.T) {
	t.Run("File exists and is json", func(t *testing.T) {
		bytes, error := ReadFiles([]string{"../../assets/file1.json"})
		assert.Nil(t, error)
		assert.NotNil(t, bytes)
		assert.Equal(t, len(bytes), 1)
	})

	t.Run("File don't exists", func(t *testing.T) {
		_, error := ReadFiles([]string{""})
		assert.NotNil(t, error)
	})

	t.Run("File exists and isn't json", func(t *testing.T) {
		_, error := ReadFiles([]string{"../../assets/test/test1.txt"})
		assert.NotNil(t, error)
		assert.Equal(t, error.Error(), "Invalid json file with path ../../assets/test/test1.txt")
	})

	t.Run("File exists and is empty", func(t *testing.T) {
		_, error := ReadFiles([]string{"../../assets/test/empty.json"})
		assert.NotNil(t, error)
		assert.Equal(t, error.Error(), "Invalid json file with path ../../assets/test/empty.json")
	})
}

func TestFilesAreEqual(t *testing.T) {
	t.Run("Compare json file to itself", func(t *testing.T) {
		path := "../../assets/file1.json"
		files, err := ReadFiles([]string{path})
		assert.Nil(t, err)

		ok, err := FilesAreEqual(files[path], files[path])
		assert.Equal(t, true, ok)
		assert.Nil(t, err)
	})

	t.Run("Compare empty file to itself", func(t *testing.T) {
		path := "../../assets/test/empty.json"
		files, err := ReadFiles([]string{path})
		assert.Equal(t, err.Error(), fmt.Sprintf("Invalid json file with path %s", path))

		ok, err := FilesAreEqual(files[path], files[path])
		assert.Equal(t, false, ok)
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "Error mashalling file 1: unexpected end of JSON input")
	})

	t.Run("Compare one empty file to json file", func(t *testing.T) {
		paths := []string{"../../assets/file1.json", "../../assets/test/empty.json"}
		files, err := ReadFiles(paths)
		assert.Equal(t, err.Error(), fmt.Sprintf("Invalid json file with path %s", paths[1]))

		ok, err := FilesAreEqual(files[paths[0]], files[paths[1]])
		assert.Equal(t, false, ok)
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "Error mashalling file 2: unexpected end of JSON input")
	})

	t.Run("Compare one non json file to json file", func(t *testing.T) {
		paths := []string{"../../assets/file1.json", "../../assets/test/test1.txt"}
		files, err := ReadFiles(paths)
		assert.Equal(t, err.Error(), fmt.Sprintf("Invalid json file with path %s", paths[1]))

		ok, err := FilesAreEqual(files[paths[0]], files[paths[1]])
		assert.Equal(t, false, ok)
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "Error mashalling file 2: unexpected end of JSON input")
	})

	t.Run("Compare one non json file to non json file", func(t *testing.T) {
		path := "../../assets/test/test1.txt"
		files, err := ReadFiles([]string{path})
		assert.Equal(t, err.Error(), fmt.Sprintf("Invalid json file with path %s", path))

		ok, err := FilesAreEqual(files[path], files[path])
		assert.Equal(t, false, ok)
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "Error mashalling file 1: unexpected end of JSON input")
	})

	t.Run("Compare json file to another json file not equal", func(t *testing.T) {
		paths := []string{"../../assets/file1.json", "../../assets/test/file.json"}
		files, err := ReadFiles(paths)
		assert.Nil(t, err)

		ok, err := FilesAreEqual(files[paths[0]], files[paths[1]])
		assert.Equal(t, false, ok)
		assert.Nil(t, err)
	})

	t.Run("Compare json file to another json file but unordered", func(t *testing.T) {
		paths := []string{"../../assets/file1.json", "../../assets/file2.json"}
		files, err := ReadFiles(paths)
		assert.Nil(t, err)

		ok, err := FilesAreEqual(files[paths[0]], files[paths[1]])
		assert.Equal(t, true, ok)
		assert.Nil(t, err)
	})
}
