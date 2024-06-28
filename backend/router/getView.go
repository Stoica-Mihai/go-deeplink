package router

import (
	"html/template"
	"net/http"
	"path"
)

func GetView(viewName string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		templatePath := path.Join(VIEWSPATH, viewName)
		tmpl, err := template.ParseFiles(templatePath)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		err = tmpl.Execute(w, nil)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
