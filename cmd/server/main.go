package main

import (
	"log"
    handler "github.com/tersorasora/freelance/internal/delivery/http"
    "github.com/tersorasora/freelance/internal/repository"
    "github.com/tersorasora/freelance/internal/usecase"
	"github.com/tersorasora/freelance/internal/entity"

    "github.com/joho/godotenv"
    "github.com/gin-gonic/gin"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
	"github.com/gin-contrib/cors"
    "os"
)

func main() {
	// Load .env
    err := godotenv.Load()
    if err != nil {
        log.Println("No .env file found, using system env")
    }

    dsn := os.Getenv("DATABASE_URL")
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect database:", err)
    }

    // Auto-migrate
    db.AutoMigrate(&entity.User{})

    // Setup layers
    // User
    userRepo := repository.NewUserRepository(db)
    userUC := usecase.NewUserUsecase(userRepo)

    // Field
    fieldRepo := repository.NewFieldRepository(db)
    fieldUc := usecase.NewFieldUseCase(fieldRepo)

    // Service
    serviceRepo := repository.NewServiceRepository(db)
    serviceUC := usecase.NewServiceUseCase(serviceRepo)

    // Setup HTTP server
    r := gin.Default()
    r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Your React app URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))
    handler.NewUserHandler(r, userUC)
    handler.NewFieldHandler(r, fieldUc)
    handler.NewServiceHandler(r, serviceUC)

    log.Println("Server running on :8080")
    r.Run(":8080")
}