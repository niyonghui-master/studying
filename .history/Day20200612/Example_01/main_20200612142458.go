package main

type Summer interface {
	Sum() int
}

func main() {
	var i ints
	var s Summer = i // nil value can satisfy interface
	fmt.Println(s==nil, s.Sum()) // true, 0
}

