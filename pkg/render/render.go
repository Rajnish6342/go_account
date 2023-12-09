package render

import (
	"fmt"
	"net/http"
	"text/template"
)

func RenderTemplate(w http.ResponseWriter, templateId string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + templateId)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error while parsing template", err)
		return
	}
}
