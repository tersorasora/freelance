package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
    "github.com/tersorasora/freelance/internal/usecase"
)

type UserHandler struct {
	uuc usecase.UserUsecase
}

func NewUserHandler(router *gin.Engine, uuc usecase.UserUsecase) {
	handler := &UserHandler{uuc}
	router.POST("/register", handler.Register)
	router.POST("/login", handler.Login)
}

func (h *UserHandler) Register(c *gin.Context) {
	type RegisterRequest struct {
		Email    string `json:"email" binding:"required,email"`
		Name     string `json:"name" binding:"required"`
		Password string `json:"password" binding:"required,min=6"`
	}

	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.uuc.RegisterUser(req.Email, req.Name, req.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusCreated, gin.H{
		"message": "Registrasi User Berhasil", 
		"user_id": user.UserID,
		"email": user.Email, 
		"name": user.Name,
	})
}

func (h *UserHandler) Login(c *gin.Context) {
	type LoginRequest struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.uuc.LoginUser(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Login Berhasil", 
		"user_id": user.UserID,
		"email": user.Email,
	})
}