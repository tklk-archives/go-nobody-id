package main

import (
	"os"
	"fmt"
	"image/png"
	"github.com/techknowlogick/go-nobody-id"
)

func main() {
	n := nobody.New(80)
	f, err := os.OpenFile("nobody.png", os.O_WRONLY|os.O_CREATE, 0600)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer f.Close()
    png.Encode(f, n)
}