package main

import "fmt"

type Artist struct{
	name, genre string
	songs int
}

func NewReleaseByVal(A Artist) int {
	A.songs++
	return A.songs
}

func NewReleaseByRef(A *Artist) int {
	A.songs++
	return A.songs
}

func main(){

	art1:= Artist{"Yellow Diary", "pop", 5}
	fmt.Printf("The genre of %s is %s and they've released their %dth song \n", art1.name, art1.genre, NewReleaseByVal(art1))
	fmt.Printf("%s had released %d songs till now\n", art1.name, art1.songs)

	art2:= &Artist{"Yellow Diary", "pop", 5}
	fmt.Printf("The genre of %s is %s and they've released their %dth song \n", art2.name, art2.genre, NewReleaseByRef(art2))
	fmt.Printf("%s had released %d songs till now\n", art2.name, art2.songs)

}
