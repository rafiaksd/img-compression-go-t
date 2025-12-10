package main

import (
	"log"
	"fmt"

	"github.com/disintegration/imaging"
)

func main() {
	src, err := imaging.Open("original.jpg")
	if err != nil {
		log.Fatalf("Failed to open image: %v", err)
	}

	// resize image, width:800px, maintain aspect ratio
	src = imaging.Resize(src, 800, 0, imaging.Lanczos)

	// save resulting image as jpg
	comp_img_name := "compressed.jpg"
	err = imaging.Save(src, comp_img_name)
	if err != nil {
		log.Fatalf("Failed to save image: %v", err)
	}
	
	fmt.Println("Image compressed successfully, saved as:", comp_img_name)
}
