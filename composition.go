package main

import "fmt"

type User struct{
	name, location string
	Id int
}

type Player struct{
	User
	GameId int
}

func (u *User) Greeting() string {
	return fmt.Sprintf("Hi %s is from %s", u.name, u.location)
}

func main(){
	p := Player{}
	p.name = "Prateek"
	p.Id=303
	p.GameId=4045
	p.location = "Kanpur"

	fmt.Println(p)
	fmt.Println(p.Greeting())
}


