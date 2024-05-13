package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

type CustomError struct {
	Message    string
	Code       string
	HttpStatus int
}

func main() {
	// Download the errors.json file from the repository
	response, err := http.Get("https://raw.githubusercontent.com/wwi21seb-projekt/error-domain/main/errors/errors.json")
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	var errors []CustomError
	if err := json.NewDecoder(response.Body).Decode(&errors); err != nil {
		panic(err)

	}

	// Generate errors.go
	outPath := filepath.Join("./", "errors.go")
	out, err := os.Create(outPath)
	if err != nil {
		panic(err)
	}
	defer out.Close()

	_, _ = out.WriteString("// CODE GENERATED BY generate.go - DO NOT EDIT\n")
	_, _ = out.WriteString("package errors\n\n")

	_, _ = out.WriteString("var (\n")
	for _, e := range errors {
		_, _ = out.WriteString(fmt.Sprintf("\t%s = &CustomError{\n\t\tMessage: \"%s\",\n\t\tCode: \"%s\",\n\t\tHttpStatus: %d,\n\t}\n", e.Code, e.Message, e.Code, e.HttpStatus))
	}
	_, _ = out.WriteString(")\n")

	// Run go fmt on the generated file

}