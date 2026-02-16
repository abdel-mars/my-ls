package main 

import (
	// "fmt"
	"os"
	"my-ls/internal/app"
)

func main() {
	args := os.Args[1:]
	app.Run(args)
}
