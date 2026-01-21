package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type Recipe struct {
	Name        string
	Ingredients []string
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	// First, let's think about recipe categories
	recipes := []Recipe{

		//recipes are ai-generated recipes
		{"Fried Egg", []string{"egg", "oil"}},
		{"Boiled Egg", []string{"egg"}},
		{"Scrambled Eggs", []string{"egg", "butter", "milk"}},
		{"Omelette", []string{"egg", "cheese", "butter", "milk"}},
		{"Egg Curry", []string{"egg", "onion", "tomato", "spices"}},
		{"Deviled Eggs", []string{"egg", "mayo", "mustard"}},

		{"Toast", []string{"bread", "butter"}},
		{"Garlic Bread", []string{"bread", "butter", "garlic"}},
		{"Bread Pudding", []string{"bread", "milk", "egg", "sugar"}},
		{"Croutons", []string{"bread", "oil", "herbs"}},
		{"Bruschetta", []string{"bread", "tomato", "basil", "garlic"}},

		{"Fried Rice", []string{"rice", "egg", "carrot", "peas"}},
		{"Rice Pudding", []string{"rice", "milk", "sugar"}},
		{"Rice Bowl", []string{"rice", "vegetables", "soy sauce"}},
		{"Lemon Rice", []string{"rice", "lemon", "mustard seeds"}},

		{"French Fries", []string{"potato", "oil", "salt"}},
		{"Mashed Potatoes", []string{"potato", "butter", "milk"}},
		{"Potato Salad", []string{"potato", "mayo", "onion"}},
		{"Baked Potato", []string{"potato", "butter", "cheese"}},

		{"Tomato Soup", []string{"tomato", "onion", "garlic", "cream"}},
		{"Tomato Pasta", []string{"pasta", "tomato", "garlic", "basil"}},
		{"Tomato Salad", []string{"tomato", "cucumber", "onion", "lemon"}},

		{"Mac & Cheese", []string{"pasta", "cheese", "milk", "butter"}},
		{"Cheese Toastie", []string{"bread", "cheese", "butter"}},
		{"Yogurt Parfait", []string{"yogurt", "fruit", "granola"}},
		{"Creamy Soup", []string{"vegetables", "cream", "butter"}},

		{"Green Salad", []string{"lettuce", "cucumber", "tomato", "dressing"}},
		{"Pasta Salad", []string{"pasta", "vegetables", "dressing"}},
		{"Bean Salad", []string{"beans", "corn", "onion", "lime"}},

		{"Vegetable Stir Fry", []string{"vegetables", "oil", "soy sauce", "garlic"}},
		{"Pasta Aglio Olio", []string{"pasta", "garlic", "olive oil", "chili"}},
		{"Bean Burrito", []string{"tortilla", "beans", "cheese", "salsa"}},
		{"Quesadilla", []string{"tortilla", "cheese", "salsa"}},

		{"Vegetable Soup", []string{"vegetables", "broth", "onion", "carrot"}},
		{"Lentil Soup", []string{"lentils", "onion", "carrot", "celery"}},
		{"Minestrone", []string{"pasta", "beans", "tomato", "vegetables"}},

		{"Pancakes", []string{"flour", "egg", "milk", "baking powder"}},
		{"Waffles", []string{"flour", "egg", "milk", "butter"}},
		{"Oatmeal", []string{"oats", "milk", "honey"}},
		{"Smoothie", []string{"fruit", "yogurt", "honey"}},

		{"Fruit Salad", []string{"fruit", "honey", "lemon"}},
		{"Mug Cake", []string{"flour", "sugar", "cocoa", "milk", "oil"}},
		{"Chocolate Pudding", []string{"milk", "cocoa", "sugar", "cornstarch"}},
		{"Rice Krispie Treats", []string{"cereal", "marshmallow", "butter"}},
	}
	fmt.Println("Pantry Chef")
	fmt.Println("What ingredients do you have? (comma-separated)")
	fmt.Println("Be careful what you type for")
	fmt.Print("Example: egg, bread, milk\n> ")

	input, _ := reader.ReadString('\n')
	fmt.Println(input)

	userIngredients := parseIngredients(input)
	fmt.Printf("\nYou have: %s\n\n", strings.Join(userIngredients, ", "))
	fmt.Print("Thinking")
	fakeAI()
	fmt.Println()
	fmt.Print("Collecting User Input and Searching")
	fakeAI()

	results := findRecipes(userIngredients, recipes)

	fmt.Println("\n\n Recipes you can make:")
	if len(results) == 0 {
		fmt.Println("Sorry, no matching recipes found.")
	} else {
		for _, r := range results {
			fmt.Println("- " + r.Name)
		}
	}
}

func parseIngredients(input string) []string {
	ingredients := strings.Split(input, ",")

	for i := range ingredients {
		ingredients[i] = strings.ToLower(strings.TrimSpace(ingredients[i]))
	}
	return ingredients
}

func fakeAI() {
	for i := 0; i < 10; i++ {
		fmt.Print(".")
		time.Sleep(200 * time.Millisecond)
	}
}

func findRecipes(userIngredients []string, recipes []Recipe) []Recipe {
	var matched []Recipe

	for _, recipe := range recipes {
		if hasAllIngredients(userIngredients, recipe.Ingredients) {
			matched = append(matched, recipe)
		}
	}
	return matched
}

func hasAllIngredients(userIngredients, required []string) bool {
	ingredientMap := make(map[string]bool)
	for _, ing := range userIngredients {
		ingredientMap[ing] = true
	}

	for _, need := range required {
		if !ingredientMap[need] {
			return false
		}
	}
	return true
}
