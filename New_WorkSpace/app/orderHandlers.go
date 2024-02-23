package app

import (
	"context"

	//"fmt"
	"log"
	"net/http"
	"time"

	"github.com/AdiLambe/TestGoLangCodes/workspace/domain"
	"github.com/AdiLambe/TestGoLangCodes/workspace/dto"
	"github.com/AdiLambe/TestGoLangCodes/workspace/service"
	"github.com/gin-gonic/gin"
)

type OrderHandlers struct {
	service service.OrderService
}

func (oh *OrderHandlers) CreateOrder(c *gin.Context) {
	var newOrder domain.Order

	// Bind the JSON data from the request body to the newOrder struct
	if err := c.BindJSON(&newOrder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	request := &dto.NewOrderRequest{
		Id:          newOrder.Id,
		Name:        newOrder.Name,
		Number:      newOrder.Number,
		Description: newOrder.Description,
		Status:      newOrder.Status,
	}

	// Call the CreateOrder method on your OrderService
	createdOrder, err := oh.service.CreateOrder(request)

	if err != nil {
		c.JSON(err.Code, gin.H{"message": err.Message})
		return
	}

	c.JSON(http.StatusCreated, createdOrder)
}

// Handler Func
func (ch *OrderHandlers) GetOrdersList(c *gin.Context) {

	status := c.Query("status") // to retrieve the value associated with the "Status " key in query parameter.
	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	orders, err := ch.service.GetOrdersList(ctx, status)
	if err != nil {
		c.JSON(err.Code, gin.H{"message": err.Message})
		return
	}
	c.JSON(http.StatusOK, orders)
}

func (ch *OrderHandlers) GetOrder(c *gin.Context) {

	id := c.Param("order_id")
	log.Printf("Received request for Order ID: %s", id)

	order, err := ch.service.GetOrder(id)
	if err != nil {
		log.Printf("Error: %+v", err)
		c.JSON(err.Code, gin.H{"message": err.Message})
		return
	} else {
		c.JSON(http.StatusOK, order)
	}
}
