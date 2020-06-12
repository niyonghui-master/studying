package main

type S struct{}

func main() {
    var x S
    y := &x
	w := identity(y)
	
	fmt.
}

func identity(z *S) *S {
    return z
}