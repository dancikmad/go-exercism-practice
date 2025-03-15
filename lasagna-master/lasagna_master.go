package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, time int) int {
	if time == 0 {
		time = 2
	}
	return len(layers) * time
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodles int, sauce float64) {
	noodlesCount := 0
	sauceCount := 0

	for _, item := range layers {
		if item == "noodles" {
			noodlesCount += 1
		} else if item == "sauce" {
			sauceCount += 1
		}
	}

	noodles = noodlesCount * 50
	sauce = float64(sauceCount) * 0.20
	return
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendListIngredients []string, ownRecipeIngredients []string) {
	ownRecipeIngredients[len(ownRecipeIngredients)-1] = friendListIngredients[len(friendListIngredients)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) []float64 {
	scaleFactor := float64(portions) / 2.0

	scaled := make([]float64, len(quantities))

	for i, amount := range quantities {
		scaled[i] = amount * scaleFactor
	}
	return scaled
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
