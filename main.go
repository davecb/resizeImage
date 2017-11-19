package main

import (
	"fmt"
	"image/jpeg"
	"log"
	"os"
	"time"
	"image"

	"github.com/nfnt/resize"
)

func main() {
	// start
	img, length := getImage()
	//set up server
	// serve
	//	interpret
	//	resize || send to resize queue
	// 	resize (if not queued)
	//https://github.com/nfnt/resize
	initial := time.Now()
	m := resize.Resize(100, 100, img, resize.NearestNeighbor)
	resizeTime := time.Since(initial)

	// return results
	saveResizedImage(m)
	fmt.Print("#date      time         decode   resize   encode   size  file     rc  op\n")
	reportPerformance(initial, 0, resizeTime, 0, length, 200)

}
func saveResizedImage(m image.Image) {
	out, err := os.Create("test_qsized.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	// write new image to file
	err = jpeg.Encode(out, m, nil)
	if err != nil {
		log.Fatal("write failure %v\n", err)
		os.Exit(0)
	}
	fmt.Println("success...")
}

// get an image to resize
func getImage() (image.Image, int64) {
	file, err := os.Open("01.jpg")
	if err != nil {
		log.Fatal(err)
	}
	// decode jpeg into image.Image
	img, err := jpeg.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	fi, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}

	return img, fi.Size()
}

// reportPerformance in standard format
func reportPerformance(initial time.Time, decode, resize,
	encode time.Duration, length int64, rc int) {

	fmt.Printf("%s %f %f %f %d %s %d RESIZE\n",
		initial.Format("2006-01-02 15:04:05.000"),
		decode.Seconds(), resize.Seconds(), encode.Seconds(), length, "file.img",
		rc)
}
