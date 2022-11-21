package middleware

import (
	"golang-kubernetes/orders/domain/order_domain"
	"golang-kubernetes/orders/service/order_service"
	"golang-kubernetes/orders/utils/error_utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func OrderAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		userData := c.MustGet("userData").(jwt.MapClaims)

		userId := int32(userData["id"].(float64))

		var orderReq order_domain.Order

		orderId, err := orderReq.GetorderIdParam(c)

		if err != nil {
			c.AbortWithStatusJSON(err.Status(), err)
			return
		}

		order, err := order_service.OrderService.GetOrderById(*orderId)

		if err != nil {
			c.AbortWithStatusJSON(err.Status(), err)
			return
		}

		if order.Status == "PAID" {
			theErr := error_utils.NewNotAuthorized("cannot update order amount of paid order")
			c.AbortWithStatusJSON(theErr.Status(), theErr)
			return
		}

		if order.UserId != userId {
			theErr := error_utils.NewNotAuthorized("you're not authorized to access this data")
			c.AbortWithStatusJSON(theErr.Status(), theErr)
			return
		}

		c.Next()

	}
}
