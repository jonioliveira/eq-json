package main

import (
	"fmt"
	"os"

	"github.com/jonioliveira/eq-json/pkg/handler"
	"github.com/jonioliveira/eq-json/pkg/validator"
)

func main() {
	args := os.Args[1:]

	if !validator.ValidateArgumentsSize(len(args)) {
		handleErrors(fmt.Errorf("You should pass only two arguments"))
	}

	if ok, path := validator.ValidateFilesPath(args); !ok {
		handleErrors(fmt.Errorf("The path %s is invalid", path))
	}

	filesContent, err := handler.ReadFiles(args)
	if err != nil {
		handleErrors(fmt.Errorf("Error when reading file: %s", err))
	}

	//this can be improved to handle multiple files and not only two
	ok, err := handler.FilesAreEqual(filesContent[args[0]], filesContent[args[0]])
	if err != nil {
		handleErrors(fmt.Errorf("Error when testing files: %e", err))
	}

	fmt.Printf("Files are equal? %t", ok)
}

func handleErrors(err error) {
	fmt.Printf("%s", err)
	os.Exit(1)
}
