package main

import "fmt"

//Sorting the name-list according to the length of the individual names and grouping them into slices

var names = []string{"Katrina", "Evan", "Neil", "Adam", "Martin", "Matt",
	"Emma", "Isabella", "Emily", "Madison",
	"Ava", "Olivia", "Sophia", "Abigail",
	"Elizabeth", "Chloe", "Samantha",
	"Addison", "Natalie", "Mia", "Alexis"}

func main() {

	var maxlen int

	for _,name := range names{
		if l:= len(name); l > maxlen{
			maxlen = l
		}
	}

	v := make([][]string, maxlen-1)

	for i := 0; i < maxlen; i++{
		for j,_ := range names{
			if len(names[j]) == i{
				v[i-1] = append(v[i-1],names[j])
			}
		}
	}

	fmt.Printf("%q",v)


}

