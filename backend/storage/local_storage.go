package storage

import (
    "os"
    "fmt"
)

func SaveLocally(data string) {
    file, err := os.Create("data.txt")
    if err != nil {
        fmt.Println("Error creating file:", err)
        return
    }
    defer file.Close()

    file.WriteString(data)
}