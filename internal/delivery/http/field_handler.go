package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tersorasora/freelance/internal/usecase"
)

type FieldHandler struct {
	fuc usecase.FieldUseCase
}

func NewFieldHandler(router *gin.Engine, fuc usecase.FieldUseCase) {
	handler := &FieldHandler{fuc}
	router.POST("/fields/create", handler.CreateField)
	router.GET("/fields", handler.GetAllFields)
	router.GET("/fields/:field_id", handler.GetFieldByID)
	router.DELETE("/fields/:field_id", handler.DeleteField)
	router.GET(("/total_fields"), handler.GetTotalFields)
}

func (h *FieldHandler) CreateField(c *gin.Context) {
	type FieldRequest struct {
		Name string `json:"name" binding:"required"`
	}

	fmt.Println("lewat")

	var req FieldRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("lewat1")

	field, err := h.fuc.CreateField(req.Name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Field created successfully",
		"field id": field.FieldID,
		"name": field.FieldName,
	})
}

func (h *FieldHandler) GetAllFields(c *gin.Context) {
	fields, err := h.fuc.GetAllFields()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"fields": fields})
}

func (h *FieldHandler) GetFieldByID(c *gin.Context) {
	field, err := h.fuc.GetFieldByID(c.Param("field_id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Field not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"field": field})
}

func (h *FieldHandler) DeleteField(c *gin.Context) {
	err := h.fuc.DeleteField(c.Param("field_id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Field deleted successfully"})
}

func (h *FieldHandler) GetTotalFields(c *gin.Context) {
	total, err := h.fuc.GetTotalFields()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"total_fields": total})
}