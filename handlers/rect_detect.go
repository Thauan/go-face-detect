package handlers

import (
	"gocv.io/x/gocv"
	"image"
	"image/color"
)

func RectTracking(input gocv.Mat, rect image.Rectangle, divisor int, color color.RGBA, text string) {
	gocv.Rectangle(&input, rect, color, 3)

	size := gocv.GetTextSize(text, gocv.FontHersheyPlain, 1.2, 2)
	pt := image.Pt(rect.Min.X+(rect.Min.X/divisor)-(size.X/divisor), rect.Min.Y-divisor)
	gocv.PutText(&input, text, pt, gocv.FontHersheyPlain, 1.2, color, 2)
}
