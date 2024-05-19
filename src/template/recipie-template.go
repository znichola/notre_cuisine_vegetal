package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"os"
	"path/filepath"
	"strings"

	"recipie-template.go/models"
)

func main() {
	recipies := listRecipies("../../database/")

	executeTemplate("home.template.html", "../../_site/index.html", recipies)

	for _, r := range recipies {
		data := extratRecipie("../../database/" + r.Url + ".json")
		executeTemplate("recipie.template.html", "../../_site/"+r.Url+"/index.html", data)
	}
	copyFile("../../static/style.css", "../../_site/style.css")
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

func listRecipies(directory string) []struct {
	Url   string
	Title string
} {
	dir, err := os.Open(directory)
	if err != nil {
		panic(fmt.Errorf("failed to open recipie directory: %w", err))
	}
	defer dir.Close()

	files, err := dir.Readdir(-1)
	if err != nil {
		panic(fmt.Errorf("failed to read recipie directory contents: %w", err))
	}

	var recipes []struct {
		Url   string
		Title string
	}
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".json") {
			url := strings.TrimSuffix(file.Name(), ".json")
			filePath := filepath.Join(directory, file.Name())

			data := extratRecipie(filePath)

			recipes = append(recipes, struct {
				Url   string
				Title string
			}{
				Url:   url,
				Title: data.Title,
			})
		}
	}
	return recipes
}

func copyFile(src string, dist string) {

	srcFile, err := os.Open(src)
	if err != nil {
		panic(err)
	}

	distFile, err := os.Create(dist)
	if err != nil {
		panic(err)
	}

	io.Copy(distFile, srcFile)
}
