package main

import (
	"log"
	"fmt"
	"os"

	"github.com/disintegration/imaging"
)

func main() {
	// expect 2 args, input/output file
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <input-image> <output-image>")
		return
	}

	input:=os.Args[1]
	output:=os.Args[2]


	src, err := imaging.Open(input)
	if err != nil {
		log.Fatalf("Failed to open image: %v", err)
	}

	// resize image, width:800px, maintain aspect ratio
	src = imaging.Resize(src, 800, 0, imaging.Lanczos)

	// save resulting image as jpg
	comp_img_name := output
	err = imaging.Save(src, comp_img_name)
	if err != nil {
		log.Fatalf("Failed to save image: %v", err)
	}
	
	fmt.Println("Image compressed successfully, saved as:", comp_img_name)
}
