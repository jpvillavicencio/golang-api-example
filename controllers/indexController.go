package controllers

import (
	"fmt"
	"myapp/common"
	"net/http"
)

type IndexController struct {
	common.Controller
}

func (c *IndexController) Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}
