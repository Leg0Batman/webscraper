package storage

import (
    "context"
    "cloud.google.com/go/storage"
    "fmt"
)

func SaveToCloud(data string) {
    ctx := context.Background()
    client, err := storage.NewClient(ctx)
    if err != nil {
        fmt.Println("Error creating cloud storage client:", err)
        return
    }
    defer client.Close()

    bucket := client.Bucket("your-bucket-name")
    obj := bucket.Object("data.txt")
    w := obj.NewWriter(ctx)
    defer w.Close()

    if _, err := w.Write([]byte(data)); err != nil {
        fmt.Println("Error writing to cloud storage:", err)
    }
}