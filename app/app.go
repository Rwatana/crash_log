package main

import (
    "fmt"
    "log"
    "os"
    "time"
)

func main() {
    logFile := os.Getenv("LOG_FILE")
    if logFile == "" {
        log.Fatal("LOG_FILE environment variable not set")
    }

    file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatalf("Failed to open log file: %v", err)
    }
    defer file.Close()

    logger := log.New(file, "app: ", log.LstdFlags)

    for i := 1; i <= 10; i++ {
        logger.Printf("Log message number %d", i)
        fmt.Printf("Written log message number %d\n", i)
        time.Sleep(2 * time.Second)
    }

    fmt.Println("App finished writing logs.")
}
