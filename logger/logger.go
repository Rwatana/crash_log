package main

import (
    "bufio"
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

    // Wait until the log file exists
    for {
        if _, err := os.Stat(logFile); err == nil {
            break
        }
        fmt.Println("Waiting for log file to be created...")
        time.Sleep(1 * time.Second)
    }

    file, err := os.Open(logFile)
    if err != nil {
        log.Fatalf("Failed to open log file: %v", err)
    }
    defer file.Close()

    reader := bufio.NewReader(file)
    fmt.Println("Logger started reading logs.")

    for {
        line, err := reader.ReadString('\n')
        if err != nil {
            // If EOF, wait and retry
            if err.Error() == "EOF" {
                time.Sleep(1 * time.Second)
                continue
            }
            log.Fatalf("Error reading log file: %v", err)
        }
        fmt.Printf("Read log: %s", line)
    }
}