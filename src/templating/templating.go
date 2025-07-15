package templating

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	fp "path/filepath"
	"time"
)

func createTemplateEngine(name string) *template.Template {
	funcMap := template.FuncMap{
		"formatDate": func(ts int64) string {
			return time.Unix(ts, 0).Format("02-Jan-2006")
		},
		"formatPrice": func(v float32) string {
			return fmt.Sprintf("%.2f", v)
		},
	}

	tmpl := template.New(name).Funcs(funcMap)

	return tmpl
}

func HtmlTemplate(file string, data map[string]any) template.HTML {
	tmplEngName := fp.Base(file)
	tmplEngine := createTemplateEngine(tmplEngName)

	tmpl, err := tmplEngine.ParseFiles(file)
	if err != nil {
		log.Println("error in templating")
		panic(err)
	}

	// Use a bytes.Buffer to store the rendered output
	var renderedOutput bytes.Buffer

	// Render the template
	err = tmpl.ExecuteTemplate(&renderedOutput, tmplEngName, data)
	if err != nil {
		log.Fatalf("Template execution: %v", err)
	}

	// Now renderedOutput holds the rendered template
	return template.HTML(renderedOutput.String())
}
