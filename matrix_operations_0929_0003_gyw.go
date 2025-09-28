// 代码生成时间: 2025-09-29 00:03:21
package main

import (
    "net/http"
    "strconv"
    "fmt"
    "log"
    "github.com/labstack/echo/v4"
)

// Matrix represents a 2D slice of integers.
type Matrix [][]int

// MatrixAdd adds two matrices and returns the result.
func MatrixAdd(a, b Matrix) (Matrix, error) {
    if len(a) != len(b) || len(a[0]) != len(b[0]) {
        return nil, fmt.Errorf("matrices have different dimensions")
    }
    rows := len(a)
    cols := len(a[0])
    result := make(Matrix, rows)
    for i := 0; i < rows; i++ {
        result[i] = make([]int, cols)
        for j := 0; j < cols; j++ {
            result[i][j] = a[i][j] + b[i][j]
        }
    }
    return result, nil
}

// MatrixSub subtracts two matrices and returns the result.
func MatrixSub(a, b Matrix) (Matrix, error) {
    if len(a) != len(b) || len(a[0]) != len(b[0]) {
        return nil, fmt.Errorf("matrices have different dimensions")
    }
    rows := len(a)
    cols := len(a[0])
    result := make(Matrix, rows)
    for i := 0; i < rows; i++ {
        result[i] = make([]int, cols)
        for j := 0; j < cols; j++ {
            result[i][j] = a[i][j] - b[i][j]
        }
    }
    return result, nil
}

// MatrixMul multiplies two matrices and returns the result.
func MatrixMul(a, b Matrix) (Matrix, error) {
    if len(a[0]) != len(b) {
        return nil, fmt.Errorf("matrices cannot be multiplied due to dimension mismatch")
    }
    rowsA := len(a)
    colsA := len(a[0])
    rowsB := len(b)
    colsB := len(b[0])
    result := make(Matrix, rowsA)
    for i := 0; i < rowsA; i++ {
        result[i] = make([]int, colsB)
        for j := 0; j < colsB; j++ {
            result[i][j] = 0
            for k := 0; k < colsA; k++ {
                result[i][j] += a[i][k] * b[k][j]
            }
        }
    }
    return result, nil
}

// MatrixTranspose transposes a matrix and returns the result.
func MatrixTranspose(m Matrix) Matrix {
    rows := len(m)
    cols := len(m[0])
    result := make(Matrix, cols)
    for i := 0; i < cols; i++ {
        result[i] = make([]int, rows)
        for j := 0; j < rows; j++ {
            result[i][j] = m[j][i]
        }
    }
    return result
}

// startServer starts the Echo HTTP server with the defined routes.
func startServer() {
    e := echo.New()
    e.GET("/add", matrixAddHandler)
    e.GET("/sub", matrixSubHandler)
    e.GET("/mul", matrixMulHandler)
    e.GET("/transpose", matrixTransposeHandler)
    
    e.Logger.Fatal(e.Start(":8080"))
}

// matrixAddHandler handles GET requests for matrix addition.
func matrixAddHandler(c echo.Context) error {
    // Retrieve matrix data from query parameters.
    // Implement your own logic to parse and validate the query parameters.
    // For simplicity, this example is omitted.
    
    // Perform matrix addition.
    a := Matrix{{1, 2}, {3, 4}}
    b := Matrix{{5, 6}, {7, 8}}
    result, err := MatrixAdd(a, b)
    if err != nil {
        return err
    }

    // Return the result as JSON.
    return c.JSON(http.StatusOK, result)
}

// matrixSubHandler handles GET requests for matrix subtraction.
func matrixSubHandler(c echo.Context) error {
    // Similar to matrixAddHandler.
    
    a := Matrix{{1, 2}, {3, 4}}
    b := Matrix{{5, 6}, {7, 8}}
    result, err := MatrixSub(a, b)
    if err != nil {
        return err
    }

    return c.JSON(http.StatusOK, result)
}

// matrixMulHandler handles GET requests for matrix multiplication.
func matrixMulHandler(c echo.Context) error {
    // Similar to matrixAddHandler.
    
    a := Matrix{{1, 2}, {3, 4}}
    b := Matrix{{5, 7}, {6, 8}}
    result, err := MatrixMul(a, b)
    if err != nil {
        return err
    }

    return c.JSON(http.StatusOK, result)
}

// matrixTransposeHandler handles GET requests for matrix transpose.
func matrixTransposeHandler(c echo.Context) error {
    // Similar to matrixAddHandler.
    
    m := Matrix{{1, 2, 3}, {4, 5, 6}}
    result := MatrixTranspose(m)

    return c.JSON(http.StatusOK, result)
}

func main() {
    startServer()
}