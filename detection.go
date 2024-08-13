package main

import (
	"image"
	"math"
)

func calculateSimilarity(img1, img2 image.Image) float64 {
	bounds := img1.Bounds()
	width, height := bounds.Dx(), bounds.Dy()

	var diff float64
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r1, g1, b1, _ := img1.At(x, y).RGBA()
			r2, g2, b2, _ := img2.At(x, y).RGBA()
			diff += math.Abs(float64(r1)-float64(r2)) +
				math.Abs(float64(g1)-float64(g2)) +
				math.Abs(float64(b1)-float64(b2))
		}
	}

	return 1 - (diff / (float64(width*height) * 3 * 65535))
}

func DetectImage(template, source image.Image, threshold float64) []image.Point {
	templateBounds := template.Bounds()
	sourceBounds := source.Bounds()

	var matches []image.Point

	// Slide the template over the source image
	for y := 0; y <= sourceBounds.Max.Y-templateBounds.Max.Y; y++ {
		for x := 0; x <= sourceBounds.Max.X-templateBounds.Max.X; x++ {
			subImg := image.NewRGBA(image.Rect(0, 0, templateBounds.Dx(), templateBounds.Dy()))
			for dy := 0; dy < templateBounds.Dy(); dy++ {
				for dx := 0; dx < templateBounds.Dx(); dx++ {
					subImg.Set(dx, dy, source.At(x+dx, y+dy))
				}
			}
			similarity := calculateSimilarity(template, subImg)

			if similarity >= threshold {
				matches = append(matches, image.Point{X: x, Y: y})
			}
		}
	}

	return matches
}
