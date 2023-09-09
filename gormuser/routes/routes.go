package routes

import (
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "github.com/Abrargit25/gormuser/handlers"
)

// InitRoutes initializes and configures the API routes
func InitRoutes(router *gin.Engine, db *gorm.DB) {
    // Create handler instances with the database reference
    userHandler := handlers.NewUserHandler(db)
    productHandler := handlers.NewProductHandler(db)

    
    router.GET("/users/:UserID", userHandler.GetUser)
    router.GET("/users/:pageNumber/:limit/:UserID", userHandler.GetUserPage)
    router.POST("/users", userHandler.CreateUser)
    router.GET("/products/:UserID/:ProductID", productHandler.GetProduct)
    router.GET("/products/:UserID", productHandler.GetProductsByUser)
    router.POST("/products", productHandler.CreateProduct)
}
