package delivery

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/szczynk/Assignment2/models"
	"gorm.io/gorm"
)

type orderRoutes struct {
	ouc models.OrderUsecase
}

func NewOrderRoute(handlers *gin.Engine, ouc models.OrderUsecase) {
	route := &orderRoutes{ouc}

	handler := handlers.Group("/orders")
	{
		handler.GET("/", route.Fetch)
		handler.POST("/", route.Store)
		handler.GET("/:id", route.GetByID)
		handler.PUT("/:id", route.Update)
		handler.DELETE("/:id", route.Delete)
	}
}

// Fetch godoc
// @Summary      Fetch orders
// @Description  get orders
// @Tags         orders
// @Accept       json
// @Produce      json
// @Success      200	{object}	[]models.Order
// @Failure      500	{object}	ErrorResponse
// @Router       /orders [get]
func (route *orderRoutes) Fetch(c *gin.Context) {
	var (
		orders []models.Order
		err    error
	)

	err = route.ouc.Fetch(c, &orders)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": orders})
}

// Store godoc
// @Summary      Create an order
// @Description  create and store an order
// @Tags         orders
// @Accept       json
// @Produce      json
// @Param        message  body  models.Order true  "Order"
// @Success      200  {object}  models.Order
// @Failure      400  {object}	ErrorResponse
// @Failure      500  {object}	ErrorResponse
// @Router       /orders [post]
func (route *orderRoutes) Store(c *gin.Context) {
	var (
		order models.Order
		err   error
	)

	err = c.ShouldBindJSON(&order)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	err = route.ouc.Store(
		c.Request.Context(),
		&order,
	)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": order})
}

// GetByID godoc
// @Summary      Show an order
// @Description  get an order by ID
// @Tags         orders
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Order ID"
// @Success      200  {object}  models.Order
// @Failure      404  {object}  ErrorResponse
// @Failure      500  {object}	ErrorResponse
// @Router       /orders/{id}	[get]
func (route *orderRoutes) GetByID(c *gin.Context) {
	id := c.Param("id")

	var (
		order models.Order
		err   error
	)

	err = route.ouc.GetByID(c, &order, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Not Found",
				"message": fmt.Sprintf("Order Data with id %s not found", id),
			})
			return
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": order})
}

// Update godoc
// @Summary      Update an order
// @Description  update an order by ID
// @Tags         orders
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Order ID"
// @Success      200  {string}  string
// @Failure      400  {object}	ErrorResponse
// @Failure      404  {object}	ErrorResponse
// @Failure      500  {object}	ErrorResponse
// @Router       /orders/{id} [put]
func (route *orderRoutes) Update(c *gin.Context) {
	id := c.Param("id")

	var (
		order models.Order
		err   error
	)

	err = c.ShouldBindJSON(&order)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	err = route.ouc.Update(
		c.Request.Context(),
		&order,
		id,
	)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Not Found",
				"message": fmt.Sprintf("Order Data with id %s not found", id),
			})
			return
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": order})
}

// Delete godoc
// @Summary      Delete an order
// @Description  delete an order by ID
// @Tags         orders
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Order ID"
// @Success      200  {string}  string
// @Failure      404  {object}	ErrorResponse
// @Failure      500  {object}	ErrorResponse
// @Router       /orders/{id} [delete]
func (route *orderRoutes) Delete(c *gin.Context) {
	id := c.Param("id")

	var (
		order models.Order
		err   error
	)

	err = route.ouc.Delete(c, &order, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Not Found",
				"message": fmt.Sprintf("Order Data with id %s not found", id),
			})
			return
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"data": fmt.Sprintf("Order with id %s has been deleted", id),
		},
	)
}
