export type Recipie = {
	title: string;
	subtitle: string;
	prepTime?: string;
	cookingTime?: string;
	servs?: string;
	sections: Section[];
};

export type Section = {
	title: string;
	ingredients: Ingredient[];
	instructions: string[];
};

export type Ingredient = {
	qt: number;
	unit: string;
	name: string;
};
