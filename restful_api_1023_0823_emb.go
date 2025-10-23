// 代码生成时间: 2025-10-23 08:23:36
 * documentation, and follows Go best practices for maintainability
 * and scalability.
 */

package main

import (
    "fmt"
    "net/http"
    "os"
    "time"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
)

// Main function to start the Echo server.
func main() {
    // Create a new Echo instance.
    e := echo.New()

    // Middleware to handle logging
    e.Use(middleware.Logger())
    // Middleware to handle recovery from panics
    e.Use(middleware.Recover())

    // Define route for GET request to /hello
    e.GET("/hello", helloHandler)

    // Start the server.
    // Default is at localhost:8080, change if necessary.
    if err := e.Start(":8080"); err != nil {
        fmt.Println("Error starting server: ", err)
    }
}

// Handler function for the /hello endpoint.
func helloHandler(c echo.Context) error {
    // Get the name parameter from the query string.
    name := c.QueryParam("name")
    if name == "" {
        // Default to saying hello to the world if no name is provided.
        name = "World"
    }

    // Return a JSON response.
    return c.JSON(http.StatusOK, map[string]string{
        "message": fmt.Sprintf("Hello, %s!", name),
    })
}