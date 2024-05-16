package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"os"

	"recipie-template.go/models"
)

func main() {
	data := extratRecipie("../../database/green_lazagna.json")

	executeTemplate("recipie.template.html", "../../static/recipie/index.html", data)
}

func executeTemplate(templateFile string, outputFile string, data any) {
	tmpl, err := template.ParseFiles(templateFile)
	if err != nil {
		panic(err)
	}

	output, err := os.Create(outputFile)
	if err != nil {
		panic(err)
	}
	defer output.Close()

	err = tmpl.Execute(output, data)
	if err != nil {
		panic(err)
	}

	fmt.Println(templateFile, "generated successfully!")
}

func extratRecipie(file string) models.Recipie {
	content, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	var payload models.Recipie
	err = json.Unmarshal(content, &payload)
	if err != nil {
		panic(err)
	}

	return payload
}
