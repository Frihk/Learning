package main

import "fmt"

type ptr struct {
	x int
	y int
}

func setPoint(pointer *ptr, newpointx int, newpointy int) {
	pointer.x = newpointx
	pointer.y = newpointy
}

func main() {
	points := ptr{}
	setPoint(&points, 42, 21)

	fmt.Printf("x = %d, y = %d\n", points.x, points.y)
}
