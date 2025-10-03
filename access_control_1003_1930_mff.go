// 代码生成时间: 2025-10-03 19:30:32
package main

import (
    "crypto/sha256"
    "encoding/hex"
    "fmt"
    "net/http"
    "strings"

    "github.com/labstack/echo/v4"
)

// User represents a user with a username and hashed password
type User struct {
    Username string
    Password string
}

// authenticate checks if the provided username and password match with the stored hash
func authenticate(username, password string, user User) bool {
    hash := sha256.Sum256([]byte(password))
    hashStr := hex.EncodeToString(hash[:])
    return user.Username == username && user.Password == hashStr
}

// middleware is an Echo middleware function that checks for user authentication
func middleware(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        authHeader := c.Request().Header.Get("Authorization")
        if authHeader == "" {
            return echo.NewHTTPError(http.StatusUnauthorized, "Authentication required")
        }

        // Split the Authorization header into username and password
        parts := strings.SplitN(authHeader, ":", 2)
        if len(parts) != 2 {
            return echo.NewHTTPError(http.StatusUnauthorized, "Invalid Authorization header")
        }

        username := parts[0]
        password := parts[1]

        // Here we would typically check against a database or another service
        // For this example, we use a hardcoded user
        user := User{Username: "admin", Password: "hashed_password"} // Replace with actual hashing
        if !authenticate(username, password, user) {
            return echo.NewHTTPError(http.StatusUnauthorized, "Invalid credentials")
        }

        return next(c)
    }
}

func main() {
    e := echo.New()

    // Define a route with access control
    e.GET("/secure", func(c echo.Context) error {
        return c.String(http.StatusOK, "Access granted")
    }, middleware)

    // Start the server
    e.Start(":8080")
}
