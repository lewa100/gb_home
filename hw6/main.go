package main

import (
	hp "github.com/gb_home/hw6/http"
	"github.com/gb_home/hw6/img"
	"net/http"
)

func main() {
	img.CreateImg()

	//Hello
	fs := http.HandlerFunc(hp.Hello)
	http.Handle("/hello", fs)
	http.ListenAndServe(":80", nil)

	//in folder testing write QECalculator and SumArgs.

	// Draw ches board
	img.DrawChessBoard()
}
