package main

import (
    "fmt"
    "gopkg.in/mgo.v2"
)

func main() {
    s, err := mgo.Dial("mongodb://127.0.0.1:27017")
    if err != nil {
        fmt.Printf("Connection failed: %v\n", err)
        return
    }
    defer s.Close()
    fmt.Println("Connected successfully!")
}