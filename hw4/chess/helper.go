package chess

// addInOptions - add in true positions.
func addInOptions(findPos Chess){
	if findPos.x >0 && findPos.y > 0 &&
		findPos.x < maxBoard + 1 && findPos.y < maxBoard + 1{
		options = append(options, findPos)
	}
}

// ng - negative Meaning n.
func ng(n int) int{
	return n * -1
}
