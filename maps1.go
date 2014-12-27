//MAPs - used to store sets of key value pairs
// all keys have to be unique ;

package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	m := map[int]int{1: 2, 2: 3} //declaring a mapping variable - initialize with a type for key and another type for value with some starting values
	fmt.Println(m)

	/*
		go run src/maps1.go
		map[1:2 2:3]
	*/

	var m2 map[int]string         //allocate memory using "var" and create a map with make function
	m2 = make(map[int]string, 10) //type of mapping and initial size
	fmt.Println(m2)

	fmt.Println(m[1]) // Similar to slices, we can add elements to the map, but instead of the index, use the key to access the value of that key

	/*
		go run src/maps1.go
		map[1:2 2:3]
		map[]
		2 <--value of key "1"
	*/
	m[1] = 15
	fmt.Println(m)
	/*
		go run src/maps1.go
		map[1:2 2:3]
		map[]
		2
		map[1:15 2:3] <--value of key 1 got modified to 15
	*/
	delete(m, 1) //delete key "1"

	fmt.Println(m)

	// Maps are used for serialization...way to represent in-memory data and to write to hdd or network socket
	// using MAPs to serialize to JSON
	//unmarshalling - from JSON to in-memory
	//http://golang.org/pkg/encoding/json/

	m3 := map[string]string{"name": "Jack"} //serialize to JSON
	res, _ := json.Marshal(m3) //marshal function to recieve a slice of bytes to export data to the application
	fmt.Println(res)

	//example for Ummarshalling?? 

}
