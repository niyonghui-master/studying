package main

type Student struct {
	name string
}

func main() {
	m := make
	m := map[string]Student{"people": {"Bob"}}
	m["people"].name = "Alice"
}
