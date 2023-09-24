package main

import "fmt"

func main() {
	var mapLit map[string]int
	var mapCreated01 map[string]float32
	var mapAssigned map[string]int

	mapLit = map[string]int{"one": 1, "two": 2}
	mapCreated := make(map[string]float32)
	mapAssigned = mapLit

	mapCreated["key1"] = 4.5
	mapCreated["key2"] = 3.14159
	mapAssigned["two"] = 3
	fmt.Println(mapLit)  // map[one:1 two:3]
	fmt.Println(mapCreated) // map[key1:4.5 key2:3.14159]
	fmt.Println(mapAssigned) // map[one:1 two:3]
	fmt.Println(mapCreated01, len(mapCreated01))   // map[] 0
	//fmt.Println(cap(mapCreated01)) //invalid argument mapCreated01 (type map[string]float32) for cap

	fmt.Printf("Map literal at \"one\" is: %d\n", mapLit["one"])  // Map literal at "one" is: 1
	fmt.Printf("Map created at \"key2\" is: %f\n", mapCreated["key2"]) //  Map created at "key2" is: 3.141590
	fmt.Printf("Map assigned at \"two\" is: %d\n", mapLit["two"]) // Map assigned at "two" is: 3
	fmt.Printf("Map literal at \"ten\" is: %d\n", mapLit["ten"]) // Map literal at "ten" is: 0
}
