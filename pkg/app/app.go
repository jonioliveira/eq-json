package app

import (
	"fmt"
	"os"

	"github.com/jonioliveira/eq-json/pkg/handler"
	"github.com/jonioliveira/eq-json/pkg/validator"
)

func Start() error {
	args := os.Args[1:]

	if !validator.ValidateArgumentsSize(len(args)) {
		return fmt.Errorf("You should pass only two arguments")
	}

	if ok, path := validator.ValidateFilesPath(args); !ok {
		return fmt.Errorf("The path %s is invalid", path)
	}

	file1Content, err := handler.ReadFile(args[0])
	if err != nil {
		return fmt.Errorf("Error when reading file %s: %s", args[0], err)
	}

	file2Content, err := handler.ReadFile(args[1])
	if err != nil {
		return fmt.Errorf("Error when reading file %s: %s", args[1], err)
	}

	ok, err := handler.FilesAreEqual(file1Content, file2Content)
	if err != nil {
		return fmt.Errorf("Error when testing files: %e", err)
	}

	fmt.Printf("Files are equal? %t", ok)
	return nil
}
