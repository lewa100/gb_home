package chess

import (
	"fmt"
	"math/rand"
	"time"
)

type Chess struct {
	x int
	y int
}


const(
	// Max Pos on Chess Board.
	maxBoard = 8
	// Shift for min position = 1.
	shift = 1
)
//  Chess Board.
/*
	1 2 3 4 5 6 7 8
1	x x x x x x x x
2	x x x x x x x x
3	x x x x x x x x
4	x x x x x x x x
5	x x x x x x x x
6	x x x x x x x x
7	x x x x x x x x
8	x x x x x x x x
*/

// For work random.
func init() {
	rand.Seed(time.Now().UnixNano())
}

// GetPos - get Random position on board.
func (ch *Chess) GetPos() {
	ch.x = rand.Intn(maxBoard) + 1
	ch.y = rand.Intn(maxBoard) + 1
}

// PrintPos - Output position to console.
func (ch *Chess) PrintPos() {
	fmt.Println("Real Horse position on board: [ X : ", ch.x, ", Y : ", ch.y, "]")
}

// True positions on board.
var options []Chess

// SearchForOptions - Search For Options for move a horse.
func (ch *Chess) SearchForOptions() []Chess {
	a := [4]int{1,2,-1,-2}
	b := [4]int{2,1,-2,-1}

	for i:=0; i< 4; i++{
		if i % 2 == 0 {
			addInOptions(Chess{ch.x+a[i] , ch.y+b[i]} )
			addInOptions(Chess{ch.x+ng(a[i]) , ch.y+b[i]} )
		}else{
			addInOptions(Chess{ch.x+a[i] , ch.y+b[i]} )
			addInOptions(Chess{ch.x+a[i] , ch.y+ng(b[i])} )
		}
	}
	return options
}
