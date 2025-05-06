package scripts

import (
	"encoding/json"
	"fmt"
	"bl-mockup-server-golang/database"
	"bl-mockup-server-golang/models"
	"os"
)

func ImportMetricsFromJSON() {
	data, err := os.ReadFile("database/brainlife_mockdata.metrics.json")
	if err != nil {
		panic(fmt.Errorf("error reading metrics json: %w", err))
	}

	var metrics []models.Metric
	err = json.Unmarshal(data, &metrics)
	if err != nil {
		panic(fmt.Errorf("error unmarshalling metrics: %w", err))
	}

	for _, metric := range metrics {
		if err := database.DB.Create(&metric).Error; err != nil {
			fmt.Println("Failed to insert metric:", metric.Name)
		}
	}
	fmt.Println("Imported", len(metrics), "metrics successfully.")
}
