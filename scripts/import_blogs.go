package scripts

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"bl-mockup-server-golang/database"
	"bl-mockup-server-golang/models"
)

// Struct trung gian để parse JSON gốc
type RawBlog struct {
	ID      string          `json:"id"`
	Title   string          `json:"title"`
	Author  models.Author   `json:"author"`
	DateRaw json.RawMessage `json:"date"`
	Views   int             `json:"views"`
	Tags    []string        `json:"tags"`
	Excerpt string          `json:"excerpt"`
	Content string          `json:"content"`
	Image   string          `json:"image"`
}

func ImportBlogsFromJSON() {
	data, err := os.ReadFile("database/brainlife_mockdata.blogs.json")
	if err != nil {
		panic(fmt.Errorf("error reading json file: %w", err))
	}

	var rawBlogs []RawBlog
	err = json.Unmarshal(data, &rawBlogs)
	if err != nil {
		panic(fmt.Errorf("error unmarshalling: %w", err))
	}

	var blogs []models.Blog
	for _, rb := range rawBlogs {
		var dateParsed time.Time
	
		// Case 1: {"$date": "..."}
		var wrapped struct {
			Date string `json:"$date"`
		}
		if err := json.Unmarshal(rb.DateRaw, &wrapped); err == nil {
			dateParsed, _ = time.Parse(time.RFC3339, wrapped.Date)
		} else {
			// Case 2: "..."
			var plain string
			if err := json.Unmarshal(rb.DateRaw, &plain); err == nil {
				dateParsed, _ = time.Parse(time.RFC3339, plain)
			} else {
				fmt.Println("⚠️ Không parse được date:", string(rb.DateRaw))
				dateParsed = time.Now()
			}
		}
	
		blogs = append(blogs, models.Blog{
			CustomID: rb.ID,
			Title:    rb.Title,
			Author:   rb.Author,
			Date:     dateParsed,
			Views:    rb.Views,
			Tags:     rb.Tags,
			Excerpt:  rb.Excerpt,
			Content:  rb.Content,
			Image:    rb.Image,
		})
	}
	
	for _, blog := range blogs {
		if err := database.DB.Create(&blog).Error; err != nil {
			fmt.Println("Failed to insert:", blog.Title, err)
		}
	}
	fmt.Println("✅ Imported", len(blogs), "blogs successfully.")
}
