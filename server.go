package main

import (
    "fmt"
    "os"
)


func main() {
    fmt.Println("Hello, World!")
}


// Page struct represents a web page with a title and body content.
type Page struct {
    Title string
    Body  []byte
}

func (p *Page) save() error {
    filename := p.Title + ".txt"
    return os.WriteFile(filename, p.Body, 0600)
}

