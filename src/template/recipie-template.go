package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"

	"recipie-template.go/models"
)

func main() {
	recipies := listRecipies("../../database/")

	executeTemplate("home.template.html", "../../static/index.html", recipies)

	for _, r := range recipies {
		data := extratRecipie("../../database/" + r + ".json")
		executeTemplate("recipie.template.html", "../../static/"+r+"/index.html", data)
	}
}

func executeTemplate(templateFile string, outputFile string, data any) {
	tmpl, err := template.ParseFiles(templateFile)
	if err != nil {
		panic(err)
	}

	outputDir := filepath.Dir(outputFile)
	err = os.MkdirAll(outputDir, os.ModePerm)
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

func listRecipies(directory string) []string {
	dir, err := os.Open(directory)
	if err != nil {
		panic(fmt.Errorf("failed to open recipie directory: %w", err))
	}
	defer dir.Close()

	files, err := dir.Readdir(-1)
	if err != nil {
		panic(fmt.Errorf("failed to read recipie directory contents: %w", err))
	}

	var fileNames []string
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".json") {
			name := strings.TrimSuffix(file.Name(), ".json")
			fileNames = append(fileNames, name)
		}
	}
	return fileNames
}
