package main

import (
	"fmt"
	"html/template"
	"os"

	"recipie-template.go/models"
)

func main() {
	// Read the external HTML file
	htmlFile := "template.html"
	content, err := os.ReadFile(htmlFile)
	if err != nil {
		fmt.Println("Error reading template file:", err)
		return
	}

	// Define the template
	tmpl, err := template.New("externalTemplate").Parse(string(content))
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return
	}

	data := models.Recipie{
		Title:        "Lazagne verte",
		Intro:        "Une variation sur le classique, super vibrante et parfumée. Un véritable arrêt sur image.",
		Ingredients:  []string{"salade", "beurre"},
		Instructions: []string{"foo", "bar"},
	}

	// Execute the template
	outputFile := "output.html"
	output, err := os.Create(outputFile)
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer output.Close()

	err = tmpl.Execute(output, data)
	if err != nil {
		fmt.Println("Error executing template:", err)
		return
	}

	fmt.Println("Template generated successfully!")
}
