package main

type Student struct {
	name string
}

func main() {
	m := make(map[string]Student)
	m["people"].name = "Alice"
}
