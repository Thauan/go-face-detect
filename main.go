package main

import (
	"fmt"
	"github.com/Thauan/go-face-detect/handlers"
	"gocv.io/x/gocv"
	"image/color"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 1 {
		fmt.Println("How to run:\n\tfacedetect [camera ID] [classifier XML file]")
		return
	}

	// parse args
	device, _ := strconv.Atoi(os.Args[1])
	xmlFile := "data/haarcascade_frontalface_default.xml"

	webcam, err := handlers.WebcamDeviceEstablish(device)

	fmt.Println(err)

	defer webcam.Close()

	// open display window
	window := gocv.NewWindow("Hello, to face detect 😍🧐")
	defer window.Close()

	// prepare image matrix
	img := gocv.NewMat()
	defer img.Close()

	// color for the rect when faces detected
	blue := color.RGBA{0, 0, 255, 0}

	// load classifier to recognize faces
	classifier := gocv.NewCascadeClassifier()
	defer classifier.Close()

	if !classifier.Load(xmlFile) {
		fmt.Printf("Error reading cascade file: %v\n", xmlFile)
		return
	}

	fmt.Printf("start reading camera device: %v\n", device)

	for {
		if ok := webcam.Read(&img); !ok {
			fmt.Printf("cannot read device %d\n", device)
			return
		}
		if img.Empty() {
			continue
		}

		// detect faces
		rects := classifier.DetectMultiScale(img)
		fmt.Printf("found %d faces\n", len(rects))

		// draw a rectangle around each face on the original image,
		// along with text identifying as "Human"
		for _, r := range rects {
			handlers.RectTracking(img, r, blue, "Human")
		}

		window.ResizeWindow(600, 600)

		window.IMShow(img)
		if window.WaitKey(1) >= 0 {
			break
		}
	}
}
