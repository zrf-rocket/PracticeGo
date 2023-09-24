package main

import (
	"fmt"
	"math"
)

func main() {
	var radius float64 = 5.0
	var area float64 = math.Pi * radius * radius
	var circumference float64 = 2 * math.Pi * radius

	fmt.Printf("半径为 %.2f 的圆的面积为 %.2f，周长为 %.2f\n", radius, area, circumference)
}
