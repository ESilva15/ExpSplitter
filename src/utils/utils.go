package utils

import (
	"bytes"
	"html/template"
	"log"
	"os"
	"path"
)

func ReadFile(path string) []byte {
	file, err := os.ReadFile(path)
	if err != nil {
		log.Printf("Unable to open file: %s\n", path)
	}
	return file
}

func HtmlTemplate(file string, data map[string]any) template.HTML {
	templateName := path.Base(file)
	tmpl, err := template.New(templateName).ParseFiles(file)
	if err != nil {
		panic(err)
	}

	// Use a bytes.Buffer to store the rendered output
	var renderedOutput bytes.Buffer

	// Render the template
	err = tmpl.ExecuteTemplate(&renderedOutput, templateName, data)
	if err != nil {
		log.Fatalf("Template execution: %v", err)
	}

	// Now, renderedOutput holds the rendered template
	return template.HTML(renderedOutput.String())
}
