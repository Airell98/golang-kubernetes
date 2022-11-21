package food_controller

import (
	"fmt"
	"golang-kubernetes/foods/domain/food_domain"
	"golang-kubernetes/foods/events/food_producer"
	"golang-kubernetes/foods/service/food_service"
	"golang-kubernetes/foods/utils/error_utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateFood(c *gin.Context) {
	var foodReq food_domain.Food

	if err := c.ShouldBindJSON(&foodReq); err != nil {
		theErrr := error_utils.NewUnprocessibleEntityError("invalid json body")
		fmt.Println(err)
		c.JSON(theErrr.Status(), theErrr)
		return
	}

	res, err := food_service.FoodService.CreateFood(&foodReq)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	go food_producer.FoodProducer.CreateFood("CREATE-FOOD", res)
	c.JSON(http.StatusOK, res)
}

func GetAllFoods(c *gin.Context) {
	foods, err := food_service.FoodService.GetAllFoods()

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, foods)
}
