package main

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/Abrargit25/gormuser/routes"
    "github.com/Abrargit25/gormuser/database"
)

func main() {
    // Initialize the database
    db := database.InitDB()

    // Create a Gin router
    r := gin.Default()

    // Initialize routes
    routes.InitRoutes(r, db)

    // Run the web server
    port := ":8080" // Change the port as needed
    if err := r.Run(port); err != nil {
        fmt.Println(err)
        return
    }

    fmt.Printf("Server is running on %s\n", port)
}
