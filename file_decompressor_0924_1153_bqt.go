// 代码生成时间: 2025-09-24 11:53:23
package main

import (
    "archive/zip"
    "bytes"
    "echo"
    "fmt"
    "io"
    "log"
    "net/http"
    "os"
    "path/filepath"
    "strings"
)

// DecompressHandler handles the file decompression
func DecompressHandler(c echo.Context) error {
    file, err := c.FormFile("file")
    if err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, "No file selected")
    }
    src, err := file.Open()
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Can't open the file")
    }
    defer src.Close()

    // Create a buffer to store the file contents
    buffer := new(bytes.Buffer)
    _, err = buffer.ReadFrom(src)
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Can't read the file")
    }

    // Define the destination directory for the extracted files
    destDir := "./extracted"
    err = os.MkdirAll(destDir, os.ModePerm)
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Can't create destination directory")
    }

    // Create a zip reader
    zipReader, err := zip.NewReader(bytes.NewReader(buffer.Bytes()), int64(buffer.Len()))
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Invalid zip file")
    }
    defer zipReader.Close()

    // Iterate through the files in the zip.
    for _, zipFile := range zipReader.File {
        fileToWrite, err := zipFile.Open()
        if err != nil {
            return echo.NewHTTPError(http.StatusInternalServerError, "Can't extract file")
        }
        defer fileToWrite.Close()

        // Get the file information
        fileInfo := zipFile.FileInfo()
        fileName := fileInfo.Name()
        if !strings.HasPrefix(fileName, "./") {
            fileName = filepath.Join(destDir, fileName)
        }

        // Check if the file already exists
        if _, err := os.Stat(fileName); os.IsNotExist(err) {
            // Make sure the directory structure is created
            dir := filepath.Dir(fileName)
            if err := os.MkdirAll(dir, os.ModePerm); err != nil {
                return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create directory")
            }
        } else if err != nil {
            return echo.NewHTTPError(http.StatusInternalServerError, "Error stating file")
        }

        // Write the file content to the destination path
        outFile, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.FileMode(fileInfo.Mode()))
        if err != nil {
            return echo.NewHTTPError(http.StatusInternalServerError, "Can't create file")
        }
        defer outFile.Close()
        _, err = io.Copy(outFile, fileToWrite)
        if err != nil {
            return echo.NewHTTPError(http.StatusInternalServerError, "Failed to write file")
        }
    }

    return c.String(http.StatusOK, "File decompressed successfully")
}

func main() {
    e := echo.New()
    e.POST("/decompress", DecompressHandler)
    e.Logger.Fatal(e.Start(":8080"))
}
