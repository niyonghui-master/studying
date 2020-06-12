package main

func test(x int) (func(), func()) {
	return func() {
			println(&x, x)
			x += 10
		}, func() {
			println(&x, x)
		}
}

func main()  {
	a, b := test(100)
	a()
	b()
}
