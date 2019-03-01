package main // import "server"

import (
	"server/handler"
	"server/mysql/controller"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

type Data struct {
	Image		[]byte		`json:"image"`
	Comment		string    	`json:"comment"`
	User_Name	string    	`json:"user_name"`
	User_ID		int			`json:"user_id"`
	Reply    	[]string   	`json:"reply"`
}

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:8080"},
		AllowMethods: []string{"GET", "POST", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"*"},
	}))

	// 追加!!
	r.Use(static.Serve("/", static.LocalFile("./images", true)))

	r.GET("/images", handler.List)
	r.POST("/images", handler.Upload, controller.CreateFood)
	r.DELETE("/images/:uuid", handler.Delete, controller.DeleteFood)
	r.Run(":8888")
}
