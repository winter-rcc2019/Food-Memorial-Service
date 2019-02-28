package sql_model

import (
	//"time"

	"../db"
	"github.com/gin-gonic/gin"
)

type Memory struct {
	//ID        	int       `json:"id"`
	//UUID      	string    	`json:"uuid"` //IRANAIYO
	Image		File		`json:"image"`
	Comment		string    	`json:"comment"`
	User_Name	string    	`json:"user_name"`
	Repry    	string    	`json:"Repry"`
	//CreatedAt 	time.Time `json:"created_at"`
}


