package main

import (
	"coffeeshop/coffee"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Welcome to the CoffeShop",
		})
	})

	r.GET("/coffee", getCoffee)
	r.Run(":8081")

}

func getCoffee(ctx *gin.Context) {
	coffeelist, _ := coffee.GetCoffees()
	ctx.String(http.StatusOK, "%s", coffeelist)
}
