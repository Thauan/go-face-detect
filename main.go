package main

import (
	"fmt"
	"image/color"
	"strconv"
	"strings"

	"github.com/Thauan/go-face-detect/handlers"
	"gocv.io/x/gocv"
)

func main() {

	answers := handlers.Checkboxes(
		"Select webcam device?",
		[]string{
			"2",
			"3",
		},
	)
	s := strings.Join(answers, ", ")

	// parse args
	device, _ := strconv.Atoi(s)
	xmlFrontalFace := "data/haarcascade_frontalface_default.xml"
	xmlEye := "data/eye.xml"

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
	red := color.RGBA{255, 0, 0, 0}

	// load classifier to recognize faces
	classifierFrontal := gocv.NewCascadeClassifier()
	classifierEye := gocv.NewCascadeClassifier()

	defer classifierFrontal.Close()
	defer classifierEye.Close()

	if !classifierFrontal.Load(xmlFrontalFace) {
		fmt.Printf("Error reading cascade file: %v\n", xmlFrontalFace)
		return
	}

	if !classifierEye.Load(xmlEye) {
		fmt.Printf("Error reading cascade file: %v\n", xmlEye)
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
		rects := classifierFrontal.DetectMultiScale(img)
		fmt.Printf("found %d faces\n", len(rects))

		// draw a rectangle around each face on the original image,
		// along with text identifying as "Human"
		for _, r := range rects {
			handlers.RectTracking(img, r, 2, blue, "Human")
		}

		// detect faces
		rectsEye := classifierEye.DetectMultiScale(img)
		fmt.Printf("found %d eyes\n", len(rectsEye))

		// draw a rectangle around each face on the original image,
		// along with text identifying as "Human"
		for _, r := range rectsEye {
			handlers.RectTracking(img, r, 4, red, "Eye")
		}

		window.ResizeWindow(600, 600)

		window.IMShow(img)
		if window.WaitKey(1) >= 0 {
			break
		}
	}
}
