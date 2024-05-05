package main

import (
	"fmt"
	"html/template"
	"os"

	"recipie-template.go/models"
)

func main() {

	tmpl := template.Must(template.ParseFiles("template.html"))

	data := models.Recipie{
		Title: "Lazagne verte",
		Intro: "Une variation sur le classique, super vibrante et parfumée. Un véritable arrêt sur image.",
		Ingredients: []string{
			"1 pack of Vegan sausages",
			"1 Onion",
			"2 cloves Garlic",
			"1 tsp Fennel seeds",
			"1 tsp Chilli flakes",
			"Salt and pepper to taste",
			"1 knob vegan butter",
			"200g Kale",
			"200g Spinach",
			"1 handful Basil (plus more for the lasagne)",
			"1 handful Parsley",
			"1 tsp Onion powder",
			"1 tsp Garlic powder",
			"150ml Veg stock",
			"1 clove garlic",
			"1 cup Peas",
			"1/4 cup Walnuts",
			"1/4 cup Nutritional yeast",
			"50g Vegan butter",
			"50g Flour",
			"300ml Soy milk",
			"1/2 tsp Nutmeg",
			"1/2 cup nutritional yeast",
			"Lasagne sheets",
			"Nutritional yeast to sprinkle on top",
		},
		Instructions: []string{
			"Start by taking a packet of vegan sausages out their casings",
			"Fry them off with 1 large white onion, 2 garlic cloves, 1 tsp fennel seeds and 1 tsp chilli flakes until golden and crispy",
			"Now for the green sauce, melt some vegan butter and add kale, spinach, handful of basil & parsley along with some onion and garlic powder in some veg stock",
			"Add to a blender & blitz with peas, garlic, walnuts and of course Nooch",
			"Next our creamy béchamel sauce. Heat in a large pan equal parts vegan butter and flour to form a roux then add soy milk until desired consistency with some Nooch and 1/2 tsp nutmeg.",
			"Then it's time to layer it up. Whatever way you want just make sure to top with béchamel, some Nooch and whack in the oven at 200c for about 40-50mins",
		},
	}

	outputFile := "../../static/index.html"
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
