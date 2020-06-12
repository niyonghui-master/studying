package main

func main() {
	fmt.P
}

func getNumber() int {
	var i int
	go func() {
		i = 5
	}()
	return i
}
