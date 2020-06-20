package img

import (
	log "github.com/sirupsen/logrus"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

const sBoard = 210
var flagColor = false

// drawCell draws Cell on board.
func drawCell(rectImg *image.RGBA, x, y, size int) {
	rectLine := image.NewRGBA(image.Rect(x, y, x+size, y+size))
	//COLOR
	greyRect := color.RGBA{169, 169, 169, 255}
	blackRect := color.RGBA{0, 0, 0, 230}

	if flagColor{
		draw.Draw(rectImg, rectLine.Bounds(), &image.Uniform{blackRect}, image.ZP, draw.Src)
	}else {
		draw.Draw(rectImg, rectLine.Bounds(), &image.Uniform{greyRect}, image.ZP, draw.Src)
	}
}

// DrawChessBoard draw chess board in png file.
func DrawChessBoard() {
	//COLOR
	board := color.RGBA{139, 69, 19, 255}

	rectImg := image.NewRGBA(image.Rect(0, 0, sBoard, sBoard))

	draw.Draw(rectImg, rectImg.Bounds(), &image.Uniform{board}, image.ZP, draw.Src)
	path := "hw6/img/"
	file, err := os.Create(path + "chessBoard.png")
	if err != nil {
		log.Fatalf("Failed create file: %s", err)
	}
	defer file.Close()
	for i := 5; i < sBoard-10; i += 25 {
		for j := 5; j < sBoard-10; j += 25 {
			//if i == j {
			drawCell(rectImg, j, i, 25)
			//}
			flagColor = !flagColor
		}
		flagColor = !flagColor
	}

	png.Encode(file, rectImg)
}
