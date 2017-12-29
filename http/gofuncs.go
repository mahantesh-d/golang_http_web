package main

import (
	"math"
	"fmt"
)

func distance(x1, y1, x2, y2 float64) float64{
	a := x2-x1
	b := y2-y1
	return math.Sqrt(a*a + b*b)
}

func rectarea(x1, y1, x2, y2 float64) float64  {
	l := distance(x1,y1,x2,y2)
	w := distance(x1,y1,x2,y2)
	return w*l
}
func areacircle(x,y,r float64) float64  {
	return math.Pi*r*r
}

func main()  {
	var rx1, ry1 float64 = 0,0
	var rx2, ry2 float64 = 10,10
	var cx, cy, cr float64  = 0,0,5

	fmt.Println("RectArea: ",rectarea(rx1,rx2,ry1,ry2))
	fmt.Println("CircleArea: ",areacircle(cx,cy,cr))
}