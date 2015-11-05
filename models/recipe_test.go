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

func TestNullId(t *testing.T){
	recipe := Recipe{}
	id := recipe.GetId()
	if id != "null_id" {
		t.Errorf("expected special id `null_id` for empty Name")
	}
}

func TestExampleId(t *testing.T){
	recipe := Recipe{}
	recipe.Name = "This is a test"
	id := recipe.GetId()
	if id != "ce114e4501d2f4e2dcea3e17b546f339" {
		t.Errorf("expected md5 for Name <<%s>>, got <<%s>>",recipe.Name,id)
	}
}



