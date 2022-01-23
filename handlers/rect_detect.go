package handlers

import (
	"gocv.io/x/gocv"
	"image"
	"image/color"
)

func RectTracking(input gocv.Mat, rect image.Rectangle, color color.RGBA, text string) {
	gocv.Rectangle(&input, rect, color, 3)

	size := gocv.GetTextSize(text, gocv.FontHersheyPlain, 1.2, 2)
	pt := image.Pt(rect.Min.X+(rect.Min.X/2)-(size.X/2), rect.Min.Y-2)
	gocv.PutText(&input, text, pt, gocv.FontHersheyPlain, 1.2, color, 2)
}
