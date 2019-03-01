package view

import (
	"encoding/json"
	"fmt"
	"net/http"

	"server/mysql/sql_model"
)

type foodsResponse struct {
	Total int           	`json:"total"`
	Foods []*sql_model.Food `json:"foods"`
}

func RenderFoods(w http.ResponseWriter, foods []*sql_model.Food) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	s, err := json.Marshal(foodsResponse{Total: len(foods), Foods: foods})
	if err != nil {
		RenderInternalServerError(w, "cant't encode foods response json")
		return
	}
	fmt.Fprintln(w, string(s))
}

func RenderFood(w http.ResponseWriter, food *sql_model.Food, statusCode int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(statusCode)
	s, err := json.Marshal(food)
	if err != nil {
		RenderInternalServerError(w, "cant't encode food response json")
		return
	}
	fmt.Fprintln(w, string(s))
}