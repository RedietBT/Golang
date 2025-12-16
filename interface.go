package main
import "fmt"

type Shape interface{
	Area() float64
}

type Circle struct{
	Radius float64
}

type Rectangle struct{
	Width, Height float64
}

func (c Circle) Area() float64{
	return 3.14 * c.Radius * c.Radius
}

func (r Rectangle) Area() float64{
	return r.Width * r.Height
}

func PrintArea(s Shape){
	fmt.Println("Area:", s.Area())
}

func interfaces(){
	c := Circle{Radius: 5}
	r := Rectangle{Width: 4, Height: 6}

	PrintArea(c)
	PrintArea(r)
}