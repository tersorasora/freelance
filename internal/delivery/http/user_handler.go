package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tersorasora/freelance/internal/auth"
	"github.com/tersorasora/freelance/internal/delivery/middleware"
	"github.com/tersorasora/freelance/internal/usecase"
)

type UserHandler struct {
	uuc usecase.UserUsecase
}

func NewUserHandler(router *gin.Engine, uuc usecase.UserUsecase) {
	handler := &UserHandler{uuc}
	router.POST("/register", handler.Register)
	router.POST("/login", handler.Login)
	router.GET("/user/:user_id", handler.GetUser)
	router.GET("/total_users", handler.GetTotalUsers)
	
	// protected routes
	authGroup := router.Group("/")
	authGroup.Use(middleware.AuthMiddleware())
	{
		authGroup.GET("/profile", handler.GetProfile)
		authGroup.DELETE("/user/:user_id", handler.DeleteUser)
	}	
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

	token, err := auth.GenerateToken(user.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login Berhasil", 
		"token": token,
		"user" : gin.H{
			"user_id": user.UserID,
			"email": user.Email,
			"name": user.Name,
		},
	})
}

func (h *UserHandler) GetUser(c *gin.Context) {
	userID := c.Param("user_id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id is required"})
		return
	}

	user, err := h.uuc.GetUser(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User found",
		"user_id": user.UserID,
		"email": user.Email,
		"name": user.Name,
		"balance": user.Balance,
		"role_id": user.RoleID,
	})
}

func (h *UserHandler) GetProfile(c *gin.Context) {
	userID := c.GetString("user_id") // from middleware
	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "not logged in"})
		return
	}

	user, err := h.uuc.GetUser(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errorwagu": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User found",
		"token": c.GetString("token"),
		"user_id": user.UserID,
		"email":   user.Email,
		"name":    user.Name,
		"balance": user.Balance,
		"role_id": user.RoleID,
	})
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	userID := c.Param("user_id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id is required"})
		return
	}
	
	err := h.uuc.DeleteUser(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User berhasil dihapus"})
}

func (h *UserHandler) GetTotalUsers(c *gin.Context) {
	total, err := h.uuc.GetTotalUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"total_users": total})
}