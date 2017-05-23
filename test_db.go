package main

import (
	"log"
	"app/common"
	"app/models"
)

func main() {
	// Initialise database
	db := common.Database{}
	db.InitDB()
	db.DropSchema(models.User{})
	db.InitSchema(models.User{})

	// Create Single User
	user := models.User{Email: "jp", Name: "JP", Age: 24}
	db.CreateModel(&user)

	// Create Multiple Users
	var users []models.User = []models.User{
		models.User{Email: "foobar", Name: "Foo", Age: 21},
		models.User{Email: "helloworld", Name: "Hello", Age: 20},
		models.User{Email: "john", Name: "John", Age: 23},
	}
	for _, user := range users {
		db.CreateModel(&user)
	}

	// Get all users
	allUsers := []models.User{}
	db.GetAllModels(&allUsers)
	for _, user := range allUsers {
		log.Print(user.Name)
	}

	// Get User by ID
	user_by_id := models.User{}
	db.GetModel(&user_by_id, 1)
	log.Print(user_by_id.Name)

	// Delete User by ID
	delete_user := models.User{}
	db.DeleteModel(&delete_user, 1)
}
