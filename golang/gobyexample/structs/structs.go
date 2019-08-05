package main

import "fmt"

type person struct {
	name string
	age int
}

func main()  {
	fmt.Println(person{"Bob",20})
	fmt.Println(person{name:"Alice",age:20})
	fmt.Println(person{name:"Tom"})
	fmt.Println(person{age:20})
	fmt.Println(&person{name:"pat", age:20})

	s := person{name:"Sean", age:50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)
}
