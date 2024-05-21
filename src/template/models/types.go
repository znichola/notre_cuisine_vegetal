package models

type Recipie struct {
	Title       string     `json:"title"`
	SubTitle    string     `json:"subtitle"`
	PrepTime    *string    `json:"prepTime,omitempty"`
	CookingTime *string    `json:"cookingTime,omitempty"`
	Servs       *string    `json:"servs,omitempty"`
	Sections    []Sections `json:"sections"`
}

type Sections struct {
	Title        string   `json:"title"`
	Ingredients  []string `json:"ingredients"`
	Instructions []string `json:"instructions"`
}

// keep it simple for now

// type IngredientSection struct {
// 	Title       string       `json:"title"`
// 	Ingredients []Ingredient `json:"ingredients"`
// }

// type Ingredient struct {
// 	Qt   string `json:"qt"`
// 	Unit string `json:"unit"`
// 	Name string `json:"name"`
// }
