package main

import (
	"log"

	"github.com/hamed0406/go-ascii-art/internal/utils"
)

func main() {
	if err := utils.GenerateTestImage(); err != nil {
		log.Fatal(err)
	}
	log.Println("Test image generated at testdata/test.png")
}
