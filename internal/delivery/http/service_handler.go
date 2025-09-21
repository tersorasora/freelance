package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tersorasora/freelance/internal/delivery/middleware"
    "github.com/tersorasora/freelance/internal/usecase"
)

type ServiceHandler struct {
	suc usecase.ServiceUseCase
}

func NewServiceHandler(router *gin.Engine, suc usecase.ServiceUseCase) {
	handler := &ServiceHandler{suc}
	router.GET("/services", handler.GetAllServices)
	router.GET("/services/search", handler.SearchServices)
	router.GET("/total_services", handler.GetTotalServices)
	
	// protected routes
	authGroup := router.Group("/")
	authGroup.Use(middleware.AuthMiddleware()) 
	{
		authGroup.POST("/services/create", handler.CreateService)
		authGroup.GET("/services/my", handler.GetMyServices)
		authGroup.DELETE("/services/:service_id", handler.DeleteService)
	}
}

func (h *ServiceHandler) GetAllServices(c *gin.Context) {
	services, err := h.suc.GetAllServices()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"services": services})
}

func (h *ServiceHandler) SearchServices(c *gin.Context) {
	fielID := c.Query("field_id")
	serviceName := c.Query("service_name")

	services, err := h.suc.SearchServices(serviceName, fielID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"services": services})
}

// Protected routes

func (h *ServiceHandler) CreateService(c *gin.Context) {
	currentUserID := c.GetString("user_id") // from middleware
	if currentUserID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "not logged in"})
		return
	}

	type serviceRequest struct {
		ServiceName string  `json:"service_name" binding:"required"`
		Description string  `json:"description" binding:"required"`
		Price       float64 `json:"price" binding:"required"`
		Period      string  `json:"period" binding:"required"`
		FieldID     string  `json:"field_id" binding:"required"`
	}

	var req serviceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	service, err := h.suc.CreateService(req.ServiceName, req.Description, req.Price, req.Period, req.FieldID, currentUserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Service created successfully",
		"service": service,
	})
}

func (h *ServiceHandler) GetMyServices(c *gin.Context) {
	currentUserID := c.GetString("user_id") // from middleware
	if currentUserID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "not logged in"})
		return
	}

	services, err := h.suc.GetMyServices(currentUserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"services": services})
}

func (h *ServiceHandler) DeleteService(c *gin.Context) {
	currentUserID := c.GetString("user_id") // from middleware
	if currentUserID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "not logged in"})
		return
	}

	serviceID := c.Param("service_id")
	if serviceID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "service_id is required"})
		return
	}

	err := h.suc.DeleteService(serviceID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Service deleted successfully"})
}

func (h *ServiceHandler) GetTotalServices(c *gin.Context) {
	total, err := h.suc.GetTotalServices()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"total_services": total})
}