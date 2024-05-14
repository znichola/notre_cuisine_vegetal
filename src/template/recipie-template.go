package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"os"

	"recipie-template.go/models"
)

func main() {

	tmpl := template.Must(template.ParseFiles("template.html"))

	data := extratRecipie("../../database/green_lazagna.json")

	outputFile := "../../static/index.html"
	output, err := os.Create(outputFile)
	if err != nil {
		panic(err)
	}
	defer output.Close()

	err = tmpl.Execute(output, data)
	if err != nil {
		panic(err)
	}

	fmt.Println("Template generated successfully!")
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
