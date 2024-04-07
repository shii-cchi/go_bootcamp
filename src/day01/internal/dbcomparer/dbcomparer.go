package dbcomparer

import (
	"day01/internal/dbreader"
	"fmt"
)

func Compare(oldData, newData dbreader.Recipes) {
	compareCakeNames(oldData, newData)

	for _, oldCake := range oldData.Cakes {
		for _, newCake := range newData.Cakes {
			if oldCake.Name == newCake.Name {
				compareCookingTime(oldCake, newCake)

				compareIngredients(oldCake, newCake)

				break
			}
		}
	}
}

func compareCakeNames(oldData, newData dbreader.Recipes) {
	for _, newCake := range newData.Cakes {
		found := false

		for _, oldCake := range oldData.Cakes {
			if newCake.Name == oldCake.Name {
				found = true
				break
			}
		}

		if !found {
			fmt.Printf("ADDED cake \"%s\"\n", newCake.Name)
		}
	}

	for _, oldCake := range oldData.Cakes {
		found := false

		for _, newCake := range newData.Cakes {
			if oldCake.Name == newCake.Name {
				found = true
				break
			}
		}

		if !found {
			fmt.Printf("REMOVED cake \"%s\"\n", oldCake.Name)
		}
	}
}

func compareCookingTime(oldCake, newCake dbreader.Cake) {
	if oldCake.Time != newCake.Time {
		fmt.Printf("CHANGED cooking time for cake \"%s\" - \"%s\" instead of \"%s\"\n", oldCake.Name, newCake.Time, oldCake.Time)
	}
}

func compareIngredients(oldCake, newCake dbreader.Cake) {
	oldIngredientsMap := make(map[string]dbreader.Ingredient)
	newIngredientsMap := make(map[string]dbreader.Ingredient)

	for _, ingredient := range oldCake.Ingredients {
		oldIngredientsMap[ingredient.Name] = ingredient
	}
	for _, ingredient := range newCake.Ingredients {
		newIngredientsMap[ingredient.Name] = ingredient
	}

	compareIngredientsName(oldCake, newCake, oldIngredientsMap, newIngredientsMap)

	compareIngredientsUnitAndCount(oldCake, newIngredientsMap)
}

func compareIngredientsName(oldCake, newCake dbreader.Cake, oldIngredientsMap, newIngredientsMap map[string]dbreader.Ingredient) {
	for _, newIngredient := range newCake.Ingredients {
		_, exists := oldIngredientsMap[newIngredient.Name]
		if !exists {
			fmt.Printf("ADDED ingredient \"%s\" for cake \"%s\"\n", newIngredient.Name, newCake.Name)
		}
	}

	for _, oldIngredient := range oldCake.Ingredients {
		_, exists := newIngredientsMap[oldIngredient.Name]
		if !exists {
			fmt.Printf("REMOVED ingredient \"%s\" for cake \"%s\"\n", oldIngredient.Name, oldCake.Name)
		}
	}

}

func compareIngredientsUnitAndCount(oldCake dbreader.Cake, newIngredientsMap map[string]dbreader.Ingredient) {
	for _, oldIngredient := range oldCake.Ingredients {
		newIngredient, exists := newIngredientsMap[oldIngredient.Name]
		if exists {
			if oldIngredient.Unit != newIngredient.Unit {
				if !(oldIngredient.Unit != "" && newIngredient.Unit == "") {
					fmt.Printf("CHANGED unit for ingredient \"%s\" for cake \"%s\" - \"%s\" instead of \"%s\"\n", oldIngredient.Name, oldCake.Name, newIngredient.Unit, oldIngredient.Unit)
				} else {
					fmt.Printf("REMOVED unit \"%s\" for ingredient \"%s\" for cake \"%s\"\n", oldIngredient.Unit, oldIngredient.Name, oldCake.Name)
				}
			} else {
				if oldIngredient.Count != newIngredient.Count {
					fmt.Printf("CHANGED unit count for ingredient \"%s\" for cake \"%s\" - \"%s\" instead of \"%s\"\n", oldIngredient.Name, oldCake.Name, newIngredient.Count, oldIngredient.Count)
				}
			}
		}
	}
}
