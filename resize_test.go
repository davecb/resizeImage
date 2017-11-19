package main

import (
	"testing"
	"github.com/nfnt/resize"
	"time"
	"fmt"
)

func TestResize(t *testing.T) {
	var tests = [] struct {
		x uint
		y uint
	} {
		{100, 100},
		{200, 200},
		{400, 400},
		{800, 800},
		{1600, 1600},
		{3200, 3200},


	}
	img, _ := getImage()
	fmt.Print("#date      time         decode   resize   encode   size  file     rc  op\n")
	for _, test := range tests {
		initial := time.Now()
		resize.Resize(test.x, test.y, img, resize.NearestNeighbor)
		resizeTime :=  time.Since(initial)
		reportPerformance(initial, 0, resizeTime, 0,
			int64(test.x * test.x), 200)
	}
}


