package main

import (
	"log"
	"net/http"
	//"os"

	"./controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET()
	log.Print(http.ListenAndServe("0.0.0.0:8080", router))
}