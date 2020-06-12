package main

type Summer interface {
	Sum() int
}

func main() {
	var t *tree
	var s Summer = t

	fmt.Println(t==nil, t.Sum())
}

