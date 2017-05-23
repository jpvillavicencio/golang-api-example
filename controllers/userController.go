package controllers

import (
	"myapp/common"
	"myapp/models"
	"net/http"
)

type UserController struct {
	common.Controller
}

func (c *UserController) AddUser(w http.ResponseWriter, r *http.Request) {
	// To Do
}

func (c *UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	db := common.Database{}
	db.InitDB("root", "password", "go_api", "")
	allUsers := []models.User{}
	db.GetAllModels(&allUsers)
	c.SendJSON(
		w,
		r,
		allUsers,
		http.StatusOK,
	)
}
