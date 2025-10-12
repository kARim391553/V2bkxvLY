// 代码生成时间: 2025-10-13 02:01:20
package main

import (
    "fmt"
    "net/http"
    "testing"

    "github.com/labstack/echo"
)

// TestSuite is a struct that will hold the Echo instance for testing
type TestSuite struct {
    e *echo.Echo
}

// SetupTestSuite initializes the Echo instance for testing
func SetupTestSuite() *TestSuite {
    e := echo.New()
    return &TestSuite{e}
}

// TestGetHandler tests the GET request handler
func TestGetHandler(t *testing.T) {
    ts := SetupTestSuite()
    defer ts.e.Close()

    // Define a test handler
    ts.e.GET("/test", func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello, World!")
    })

    // Perform the GET request
    req := echo.NewRequest(http.MethodGet, "/test", nil)
    req.SetHost("example.com")
    res, err := ts.e.ServeHTTP(req, nil)
    if err != nil {
        t.Errorf("Expected nil, got %v", err)
    }

    // Check the status code
    if res.Status != http.StatusOK {
        t.Errorf("Expected status %v, got %v", http.StatusOK, res.Status)
    }

    // Check the response body
    expected := "Hello, World!"
    body, _ := res.Body.String()
    if body != expected {
        t.Errorf("Expected body %v, got %v", expected, body)
    }
}

func main() {
    // Run tests
    testing.Main()
}
