package main

import "fmt"

type Personinfo struct {
	ID string
	Name string
	age int
}

func main() {
	var personDB map[string] Personinfo
	personDB = make(map[string] Personinfo)

	personDB["wu"] = Personinfo{"w","Wu",12}
	personDB["jian"] = Personinfo{"j","jian",16}
	
	person, judge := personDB["wu"]

	if judge {
		fmt.Println("find:", person.Name, "age:", person.age)
	} else {
		fmt.Println("not find.")
	}
}
