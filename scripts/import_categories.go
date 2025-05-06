package scripts

import (
	"encoding/json"
	"fmt"
	"bl-mockup-server-golang/database"
	"bl-mockup-server-golang/models"
	"os"
)

func ImportCategoriesFromJSON() {
	data, err := os.ReadFile("database/brainlife_mockdata.categories.json")
	if err != nil {
		panic(fmt.Errorf("error reading categories json: %w", err))
	}

	var categories []models.Category
	err = json.Unmarshal(data, &categories)
	if err != nil {
		panic(fmt.Errorf("error unmarshalling categories: %w", err))
	}

	for _, category := range categories {
		if err := database.DB.Create(&category).Error; err != nil {
			fmt.Println("Failed to insert category:", category.Name)
		}
	}
	fmt.Println("Imported", len(categories), "categories successfully.")
}
