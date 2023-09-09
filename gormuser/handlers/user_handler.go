package handlers

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "github.com/Abrargit25/gormuser/models"
)

// UserHandler struct to hold the database instance
type UserHandler struct {
    DB *gorm.DB
}

// NewUserHandler creates a new UserHandler with the given database instance
func NewUserHandler(db *gorm.DB) *UserHandler {
    return &UserHandler{DB: db}
}

// GetUser retrieves a specific user by UserID
func (h *UserHandler) GetUser(c *gin.Context) {
    userID := c.Param("UserID")
    var user models.User
    if err := h.DB.Where("user_id = ?", userID).First(&user).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }
    c.JSON(http.StatusOK, user)
}

// GetUserPage retrieves a paginated list of users
func (h *UserHandler) GetUserPage(c *gin.Context) {
    pageNumber, _ := strconv.Atoi(c.Param("pageNumber"))
    limit, _ := strconv.Atoi(c.Param("limit"))
    offset := (pageNumber - 1) * limit

    var users []models.User
    if err := h.DB.Offset(offset).Limit(limit).Find(&users).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users"})
        return
    }

    c.JSON(http.StatusOK, users)
}

// CreateUser creates a new user
func (h *UserHandler) CreateUser(c *gin.Context) {
    var users []models.User

    // Generate and insert 100 users
    for i := 101; i <= 2000; i++ {
        user := models.User{
            UserID: strconv.Itoa(i), // Unique UserID for each user
            Email:  "user" + strconv.Itoa(i) + "@example.com",
        }
        users = append(users, user)
    }

    if err := h.DB.Create(&users).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create users"})
        return
    }

    c.JSON(http.StatusCreated, users)
}
