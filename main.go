package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

func main() {
	tmpl, err := template.
		New("").Parse(`
		{{ define "a" }}
		a is for apple
		{{ end }}

		{{ define "b" }}
		b is for banana
		{{ end }}

		{{ define "c" }}
		c is for cherry
		{{ end }}

		{{ template "a" }}
		{{ template "b" }}
		{{ template "c" }}

		reee
		`)
	if err != nil {
		log.Fatal(err)
	}
	tmpl, err = tmpl.
		// New("all").
		Parse(`
		{{- define "a" }}
		a is for Apple
		{{- end -}}

		{{- define "b" }}
		b is for Banana
		{{- end -}}

		{{- define "c" }}
		c is for Cherry
		{{- end -}}

		heehee`)
	fmt.Printf("current template is: \"%s\"%s\n", tmpl.Name(), tmpl.DefinedTemplates())
	err = tmpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatal(err)
	}
}
