package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"math"
	"os"
)

func main() {
	createLogo()
}

func createLogo() {
	logo := image.NewRGBA(image.Rect(0, 0, 300, 300))

	bgColor := color.RGBA{31, 15, 83, 255}
	draw.Draw(logo, logo.Bounds(), &image.Uniform{bgColor}, image.Point{}, draw.Src)

	drawCat(logo)

	file, err := os.Create("logo/amazing_logo.png")
	if err != nil {
		log.Fatalf("Error creating file amazing_logo.png: %s", err)
	}
	defer file.Close()

	err = png.Encode(file, logo)
	if err != nil {
		log.Fatalf("Error creating logo: %s", err)
	}
}

func drawCat(img *image.RGBA) {
	catColor := color.RGBA{255, 93, 178, 255}
	drawCircle(img, image.Point{150, 150}, 80, catColor)

	earColor := color.RGBA{255, 93, 178, 255}
	drawEarLeft(img, 80, 80, earColor)
	drawEarRight(img, 180, 80, earColor)

	eyeColor := color.RGBA{31, 15, 83, 255}
	drawOval(img, image.Point{120, 130}, 15, 20, eyeColor)
	drawOval(img, image.Point{180, 130}, 15, 20, eyeColor)

	noseColor := color.RGBA{31, 15, 83, 255}
	drawCircle(img, image.Point{150, 170}, 5, noseColor)

	whiskerColor := color.RGBA{31, 15, 83, 255}
	drawWhiskers(img, image.Point{150, 170}, whiskerColor)

	mouthColor := color.RGBA{31, 15, 83, 255}
	drawMouth(img, image.Point{150, 180}, 25, 20, mouthColor)
}

func drawOval(img *image.RGBA, center image.Point, width, height int, clr color.Color) {
	for y := -height; y <= height; y++ {
		for x := -width; x <= width; x++ {
			if float64(x*x)/(float64(width*width)/4)+float64(y*y)/(float64(height*height)/4) <= 1 {
				img.Set(center.X+x, center.Y+y, clr)
			}
		}
	}
}

func drawCircle(img *image.RGBA, center image.Point, radius int, clr color.Color) {
	for y := -radius; y <= radius; y++ {
		for x := -radius; x <= radius; x++ {
			if x*x+y*y <= radius*radius {
				img.Set(center.X+x, center.Y+y, clr)
			}
		}
	}
}

func drawEarLeft(img *image.RGBA, x, y int, clr color.Color) {
	earPoints := []image.Point{
		{x, y + 40},
		{x, y - 40},
		{x + 40, y},
	}
	drawPolygon(img, earPoints, clr)
}

func drawEarRight(img *image.RGBA, x, y int, clr color.Color) {
	earPoints := []image.Point{
		{x, y},
		{x + 40, y - 40},
		{x + 40, y + 40},
	}
	drawPolygon(img, earPoints, clr)
}

func drawPolygon(img *image.RGBA, points []image.Point, clr color.Color) {
	for i := 1; i < len(points); i++ {
		drawLine(img, points[i-1], points[i], clr)
	}
	drawLine(img, points[len(points)-1], points[0], clr)
}

func drawLine(img *image.RGBA, p1, p2 image.Point, clr color.Color) {
	dx := p2.X - p1.X
	dy := p2.Y - p1.Y
	steps := int(math.Max(math.Abs(float64(dx)), math.Abs(float64(dy))))
	if steps == 0 {
		return
	}
	xIncrement := float64(dx) / float64(steps)
	yIncrement := float64(dy) / float64(steps)
	x := float64(p1.X)
	y := float64(p1.Y)
	for i := 0; i <= steps; i++ {
		img.Set(int(x+0.5), int(y+0.5), clr)
		x += xIncrement
		y += yIncrement
	}
}

func drawWhiskers(img *image.RGBA, start image.Point, clr color.Color) {
	drawLine(img, start, image.Point{start.X - 50, start.Y}, clr)
	drawLine(img, start, image.Point{start.X - 50, start.Y - 10}, clr)
	drawLine(img, start, image.Point{start.X - 50, start.Y + 10}, clr)

	drawLine(img, start, image.Point{start.X + 50, start.Y}, clr)
	drawLine(img, start, image.Point{start.X + 50, start.Y + 10}, clr)
	drawLine(img, start, image.Point{start.X + 50, start.Y - 10}, clr)
}

func drawMouth(img *image.RGBA, center image.Point, width, height int, clr color.Color) {
	startAngle := -math.Pi / 5
	endAngle := math.Pi / 5

	for angle := startAngle; angle <= endAngle; angle += 0.01 {
		x := float64(center.X) + float64(width)*math.Sin(angle)
		y := float64(center.Y) + float64(height)*math.Cos(angle)
		img.Set(int(x+0.5), int(y+0.5), clr)
	}
}
