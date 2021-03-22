package main

import "fmt"

type User struct{
	firstname, lastname string
}

type Customer struct{
	Id int
	fullname string
}

type Namer interface {
	Name() string
}

func (u *User) Name() string{
	return fmt.Sprintf("%s %s",u.firstname,u.lastname)
}

func (c *Customer) Name() string {
	return fmt.Sprintf("%s",c.fullname)
}

func Greet(n Namer) string {
	return fmt.Sprintf("Hi %s", n.Name())

}


func main(){
	u := &User{"Matt","Damon"}
	fmt.Println(Greet(u))

	c := &Customer{31,"Jagdish"}
	fmt.Println(Greet(c))

}
