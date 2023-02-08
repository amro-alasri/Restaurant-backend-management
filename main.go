package main

import (
	"os"

	"golang-restaurant-management/middleware"
	"golang-restaurant-management/routes"

	"github.com/gin-gonic/gin"
)

// var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func main() {

	// configuration of the port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	// create the gin router
	router := gin.New()
	router.Use(gin.Logger())

	// make UserRoute middleware before authentication middleware to all user to register
	routes.UserRoutes(router)
	// authenticate middleware
	router.Use(middleware.Authentication())

	// the routers
	routes.FoodRoutes(router)
	routes.MenuRoutes(router)
	routes.TableRoutes(router)
	routes.OrderRoutes(router)
	routes.OrderItemRoutes(router)
	routes.InvoiceRoutes(router)

	// run the server
	router.Run(":" + port)
}
