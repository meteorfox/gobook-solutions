package main

import (
        "fmt"
        "math"
)

type Shape interface {
        area() float64
        perimeter() float64
}

type Circle struct {
        x, y, r float64
}

func (c *Circle) area() float64 {
        return math.Pi * c.r * c.r
}

func (c *Circle) perimeter() float64 {
        return math.Pi * 2 * c.r
}

type Rectangle struct {
        x1, y1, x2, y2 float64
}

func (r *Rectangle) area() float64 {
        l := distance(r.x1, r.y1, r.x1, r.x2)
        w := distance(r.x1, r.y1, r.x2, r.y1)
        return l * w
}

func (r *Rectangle) perimeter() float64 {
        l := distance(r.x1, r.y1, r.x1, r.x2)
        w := distance(r.x1, r.y1, r.x2, r.y1)
        return 2 * (l + w)
}

func distance(x1, y1, x2, y2 float64) float64 {
        a := x2 - x1
        b := y2 - y1
        return math.Sqrt(a*a + b*b)
}

func main() {
        c := Circle{0, 0, 5}
        r := Rectangle{0, 0, 10, 10}

        shapes := []Shape{&c, &r}
        var area float64
        var perimeter float64
        for _, s := range shapes {
                area += s.area()
                perimeter += s.perimeter()
        }

        fmt.Println("Total area of shapes:", area)
        fmt.Println("Total perimeter of shapes:", perimeter)

}
