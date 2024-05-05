package models

type Recipie struct {
	Title        string   `json:"title"`
	Intro        string   `json:"intro"`
	PrepTime     *string  `json:"prepTime,omitempty"`
	CookingTime  *string  `json:"cookingTime,omitempty"`
	Servs        *string  `json:"servs,omitempty"`
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
