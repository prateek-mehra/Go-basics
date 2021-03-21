package main

import (
	"fmt"
	"strings"
	"math"
)

type Mystr string

func (s Mystr) Uppercase() string{
	return strings.ToUpper(string(s))
}

type vertex struct{
	X,Y float64
}

func (v *vertex) scale(f float64){
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *vertex) abs() float64{
	return math.Sqrt(v.X*v.X+v.Y*v.Y)
}

func main(){
	var s Mystr = "Prateek"

	fmt.Printf("%s",s.Uppercase())

	v := &vertex{2,3}

	v.scale(5)

	fmt.Printf("%f",v.abs())

}
