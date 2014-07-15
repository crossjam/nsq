package templates

import (
	"html/template"
	"log"
)

var sources map[string]string
var T *template.Template

func init() {
	T = template.New("_top_")
}

func registerTemplate(name string, source string) {
	if sources == nil {
		sources = make(map[string]string)
	}
	_, ok := sources[name]
	if ok {
		log.Fatalf("ERROR: duplicate template %s", name)
	}
	sources[name] = source
}

func Parse() {
	log.Printf("INFO: parsing templates")
	for name, source := range sources {
		log.Printf("INFO: parsing template %s", name)
		_, err := T.New(name).Parse(source)
		if err != nil {
			log.Fatalf("ERROR: failed to parse template %s - %s", name, err)
		}
	}
}

func ReplaceTemplate(name string, source string) {
	if sources == nil {
		sources = make(map[string]string)
	}
	_, ok := sources[name]
	if ok {
		log.Printf("INFO: replaced template %s", name)
	}
	sources[name] = source
}
