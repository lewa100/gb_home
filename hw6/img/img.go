package img

import (
	log "github.com/sirupsen/logrus"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

const (
	wRect = 200
	hRect = 200
)

// hLine draws a horizontal line.
func hLine(rectImg *image.RGBA, color color.RGBA, x, y, length , width int) {
	rectLine := image.NewRGBA(image.Rect(x, y, x+length, y+width))
	draw.Draw(rectImg, rectLine.Bounds(), &image.Uniform{color}, image.ZP, draw.Src)
}

// vLine draws a veritcal line.
func vLine(rectImg *image.RGBA, color color.RGBA, x, y, length , width int) {
	rectLine := image.NewRGBA(image.Rect(x, y, x+width, y+length))
	draw.Draw(rectImg, rectLine.Bounds(), &image.Uniform{color}, image.ZP, draw.Src)
}

// CreateImg create image with lines.
func CreateImg() {
	green := color.RGBA{0, 255, 0, 255}
	rectImg := image.NewRGBA(image.Rect(0, 0, wRect, hRect))

	draw.Draw(rectImg, rectImg.Bounds(), &image.Uniform{green}, image.ZP, draw.Src)
	path := "hw6/img/"
	file, err := os.Create(path + "rectangle.png")
	if err != nil {
		log.Fatalf("Failed create file: %s", err)
	}
	defer file.Close()
	for i:= 5; i < wRect; i++{
		if i % 25 == 0{
			hLine(rectImg,color.RGBA{0,10,200,255}, 5,i, 190, 5)
			vLine(rectImg,color.RGBA{0,10,200,255}, i ,5, 190, 5)
		}
	}

	png.Encode(file, rectImg)
}
