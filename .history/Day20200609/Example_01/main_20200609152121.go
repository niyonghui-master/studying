package main



func getNumber() int {
	var i int
	go func() {
		i = 5
	}()
	return i
}
