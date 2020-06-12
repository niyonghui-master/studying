package main

type Summer interface {
	Sum() int
}

func main() {
	var t *tree
	var s Summer = t

	fmt.Println(t==nil, t.Sum())



var i ints
var s Sumer = i // nil value can satisfy interface
fmt.Println(s==nil, s.Sum()) // true, 0
}

