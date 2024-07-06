package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "strconv"
    "sync"
)

type Post struct {
    ID   int    `json:"id"`
    Body string `json:"body"`
}

var (
    posts   = make(map[int]Post)
    nextID  = 1
    postsMu sync.Mutex
)
