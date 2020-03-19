package main

import (
	"log"
	"os"

   - "github.com/rasanjaya85/goinaction/code/matchers"
	 "github.com/rasanjaya85/goinaction/code/search"
)


func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("President")
}