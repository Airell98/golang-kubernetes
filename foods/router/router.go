package router

import (
	"golang-kubernetes/foods/controller/food_controller"
	"golang-kubernetes/foods/db"
	"golang-kubernetes/foods/events/food_listener"
	"golang-kubernetes/foods/events/food_producer"
	"golang-kubernetes/foods/middleware"

	"github.com/gin-gonic/gin"
)

var PORT = ":8081"

func init() {
	db.InitializeDB()
	food_producer.FoodProducer.SetUpProducer()
}

func StartRouter() {
	route := gin.Default()

	foodRoute := route.Group("/foods")
	{
		foodRoute.Use(middleware.Authentication())
		foodRoute.POST("/", middleware.AdminAuthorization(), food_controller.CreateFood)
		foodRoute.GET("/", middleware.AdminAuthorization(), food_controller.GetAllFoods)

	}

	go food_listener.FoodListener.InitiliazeMainListener()
	route.Run(PORT)
}
