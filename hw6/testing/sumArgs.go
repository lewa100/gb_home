package testing

// SumArgs for addition of arguments.
func SumArgs(args ... int) int{
	sm := 0
	for _, v := range args {
		sm += v
	}
	return sm
}
