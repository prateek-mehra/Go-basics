package main

import "fmt"

func main(){


	//Arrays
	a := [2]string{"hello","world"}

	fmt.Println(a)

	var mat[2][3]string

	for i := 0; i < 2; i++{
		for j := 0; j < 3; j++{
			fmt.Sprint("row %d - column %d", i+1, j+1)
		}
	}

	fmt.Printf("%q",mat)


	//Slices
	p := []int{2,3,5,7,9}

	fmt.Println(p)

	fmt.Println(p[:3])

	cities := make([]string,3)

	cities[0] = "Kanpur"
	cities[1] = "Allahabad"
	cities[2] = "Agra"

	fmt.Printf("%q",cities)

	cities = append(cities,"Mumbai")

	fmt.Printf("%q",cities)

	fmt.Println(len(cities))

	var z []int

	fmt.Println(z,len(z),cap(z))

	if z == nil{
		fmt.Println("nil!")
	}

	pow := make([]int, 10)

	for i := range pow{
		pow[i] = 1<<uint(i)
	}

	for _,value := range pow{
		fmt.Printf("%d ",value)
	}

	population := map[string]int{
		"Kanpur" : 230,
		"Lucknow" : 303,
		"Varanasi" : 430,
	}

	for key, value := range population{
		fmt.Printf("%s has %d inhabitants\n",key,value)
	}

}
