package models

import "testing"

func TestRecipeExample(t *testing.T) {
	recipe, err := GetExample()

	if err != nil {
		t.Errorf("expected recipe: " + err.Error())
	}
	if recipe.Name == "" {
		t.Errorf("expected recipe Name, got empty Name")
	}
	if len(recipe.Instructions) == 0 {
		t.Errorf("expected recipe instructions, got none")
	}
	if len(recipe.Ingredients) == 0 {
		t.Errorf("expected recipe ingredients, got none")
	}
	if len(recipe.Tags) == 0 {
		t.Errorf("expected recipe tags, got none")
	}
}
