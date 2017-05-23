package controllers

import (
	"fmt"
	"app/common"
	"net/http"
)

type IndexController struct {
	common.Controller
}

func (c *IndexController) Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}
