package handlers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "github.com/Abrargit25/gormuser/models"
)

// ProductHandler struct to hold the database instance
type ProductHandler struct {
    DB *gorm.DB
}

// NewProductHandler creates a new ProductHandler with the given database instance
func NewProductHandler(db *gorm.DB) *ProductHandler {
    return &ProductHandler{DB: db}
}

// GetProduct retrieves a specific product by UserID and ProductID
func (h *ProductHandler) GetProduct(c *gin.Context) {
    userID := c.Param("UserID")
    productID := c.Param("ProductID")

    var product models.Product
    if err := h.DB.Where("user_id = ? AND product_id = ?", userID, productID).First(&product).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
        return
    }

    c.JSON(http.StatusOK, product)
}

// GetProductsByUser retrieves a list of products by UserID
func (h *ProductHandler) GetProductsByUser(c *gin.Context) {
    userID := c.Param("UserID")

    var products []models.Product
    if err := h.DB.Where("user_id = ?", userID).Find(&products).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve products"})
        return
    }

    c.JSON(http.StatusOK, products)
}

// CreateProduct creates a new product
func (h *ProductHandler) CreateProduct(c *gin.Context) {
    var product models.Product
    if err := c.BindJSON(&product); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
        return
    }

    if err := h.DB.Create(&product).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
        return
    }

    c.JSON(http.StatusCreated, product)
}
