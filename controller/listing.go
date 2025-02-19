package controller

import (
	"errors"
	"log"
	"strconv"
	"github.com/dancankarani/delha-frontend/database"
	"github.com/dancankarani/delha-frontend/model"
	"github.com/dancankarani/delha-frontend/utilities"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

var db = database.ConnectDB()

// 📌 Create a new listing with image upload
func CreateListing(c *fiber.Ctx) error {
	db.AutoMigrate(&model.Listing{})
	listing := new(model.Listing)

	// Parse the form data
	form, err := c.MultipartForm()
	if err != nil {
		log.Printf("Error parsing multipart form: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse form data",
		})
	}

	// Save image and get URL
	url, err := utilities.SaveFile(c, "image")
	if err != nil {
		log.Printf("Error saving file: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to save file",
		})
	}
	listing.ImageURL = url

	// Get title
	if len(form.Value["title"]) > 0 {
		listing.Title = form.Value["title"][0]
	} else {
		return errors.New("title field is missing")
	}

	// Get description
	if len(form.Value["description"]) > 0 {
		listing.Description = form.Value["description"][0]
	}

	// Get location
	if len(form.Value["location"]) > 0 {
		listing.Location = form.Value["location"][0]
	} else {
		return errors.New("location field is missing")
	}

	// Get size
	if len(form.Value["size"]) > 0 {
		sizeStr := form.Value["size"][0]
		size, err := strconv.ParseFloat(sizeStr, 64)
		if err != nil {
			log.Println("Error converting size:", err)
			return errors.New("invalid size value")
		}
		listing.Size = size
	} else {
		return errors.New("size field is missing")
	}

	// Get price
	if len(form.Value["price"]) > 0 {
		priceStr := form.Value["price"][0]
		price, err := strconv.ParseFloat(priceStr, 64)
		if err != nil {
			log.Println("Error converting price:", err)
			return errors.New("invalid price value")
		}
		listing.Price = price
	} else {
		return errors.New("price field is missing")
	}

	// Assign ID
	listing.ID = uuid.New()

	// Save to database
	if err := db.Create(&listing).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not create listing",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(listing)
}

// 📌 Get all listings
func GetListings(c *fiber.Ctx) error {
	var listings []model.Listing
	if err := db.Find(&listings).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not fetch listings",
		})
	}
	return c.JSON(listings)
}

// 📌 Get a single listing
func GetListing(c *fiber.Ctx) error {
	id := c.Params("id")
	var listing model.Listing
	if err := db.First(&listing, "id = ?", id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Listing not found",
		})
	}
	return c.JSON(listing)
}

// 📌 Update a listing
func UpdateListing(c *fiber.Ctx) error {
	id := c.Params("id")
	var listing model.Listing

	if err := db.First(&listing, "id = ?", id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Listing not found",
		})
	}

	form, err := c.MultipartForm()
	if err != nil {
		log.Printf("Error parsing multipart form: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse form data",
		})
	}

	// Check if a new image is uploaded
	if _, exists := form.File["image"]; exists {
		url, err := utilities.SaveFile(c, "image")
		if err != nil {
			log.Printf("Error saving file: %v", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to save file",
			})
		}
		listing.ImageURL = url
	}

	// Update fields
	if len(form.Value["title"]) > 0 {
		listing.Title = form.Value["title"][0]
	}
	if len(form.Value["description"]) > 0 {
		listing.Description = form.Value["description"][0]
	}
	if len(form.Value["location"]) > 0 {
		listing.Location = form.Value["location"][0]
	}
	if len(form.Value["size"]) > 0 {
		size, err := strconv.ParseFloat(form.Value["size"][0], 64)
		if err == nil {
			listing.Size = size
		}
	}
	if len(form.Value["price"]) > 0 {
		price, err := strconv.ParseFloat(form.Value["price"][0], 64)
		if err == nil {
			listing.Price = price
		}
	}

	db.Save(&listing)

	return c.JSON(listing)
}

// 📌 Delete a listing
func DeleteListing(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := db.Delete(&model.Listing{}, "id = ?", id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not delete listing",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Listing deleted successfully",
	})
}
