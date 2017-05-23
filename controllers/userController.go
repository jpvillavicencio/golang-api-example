package controllers

import (
	"encoding/json"
	"log"
	"app/common"
	"app/models"
	"net/http"
)

type UserController struct {
	common.Controller
}

func (c *UserController) UserAdd(w http.ResponseWriter, r *http.Request) {
	db := common.Database{}
	db.InitDB()
	var user models.User
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	db.CreateModel(&user)
	log.Print(user)
	c.SendJSON(
		w,
		r,
		user,
		http.StatusCreated,
	)
}

func (c *UserController) UserGetAll(w http.ResponseWriter, r *http.Request) {
	db := common.Database{}
	db.InitDB()
	allUsers := []models.User{}
	db.GetAllModels(&allUsers)
	c.SendJSON(
		w,
		r,
		allUsers,
		http.StatusOK,
	)
}

func (c *UserController) UserGet(w http.ResponseWriter, r *http.Request) {
	db := common.Database{}
	db.InitDB()
	user := models.User{}
	db.GetModel(&user, 1)
	c.SendJSON(
		w,
		r,
		user,
		http.StatusOK,
	)
}
