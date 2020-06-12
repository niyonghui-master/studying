package main

import "fmt"

type student struct {
	Name string
	Age  int
}

func pase_student() {
	m := make(map[string]*student)
	stus := []student{
		{
			Name: "Ni",
			Age:  21,
		},
		{
			Name: "Yong",
			Age:  22,
		},
		{
			Name: "Hui",
			Age:  23,
		},
	}

	for i, s := range stus {
		m[s.Name] = &stus[i]
	}

	fmt.Println(m)

	for k, v := range m {
		fmt.Println(k, " => ", v.Name)
	}
}

func main()  {
	pase_student()
}
