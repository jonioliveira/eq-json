package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/jonioliveira/eq-json/pkg/equality"
)

func FilesAreEqual(file1, file2 []byte) (bool, error) {
	var result1, result2 interface{}

	var err error
	err = json.Unmarshal(file1, &result1)
	if err != nil {
		return false, fmt.Errorf("Error mashalling string 1 :: %s", err.Error())
	}
	err = json.Unmarshal(file2, &result2)
	if err != nil {
		return false, fmt.Errorf("Error mashalling string 2 :: %s", err.Error())
	}

	return equality.JsonObjectsEquals(result1, result2), nil
}

func ReadFile(path string) ([]byte, error) {
	jsonFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	if !json.Valid(byteValue) {
		return nil, fmt.Errorf("Invalid json file")
	}

	return byteValue, nil
}
