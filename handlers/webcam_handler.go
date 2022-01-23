package handlers

import (
	"fmt"
	"gocv.io/x/gocv"
)

func WebcamDeviceEstablish(device int) (*gocv.VideoCapture, error) {

	// Avaibles in my machine: 0, 2
	webcam, err := gocv.VideoCaptureDevice(device)

	if err != nil {
		fmt.Println(err)
	}

	webcam.Set(3, float64(320))
	webcam.Set(4, float64(320)*0.75)
	webcam.Set(5, 6)

	return webcam, err
}
