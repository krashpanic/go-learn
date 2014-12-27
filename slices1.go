package main

import "fmt"

func main() {
	s := []int{1, 2, 4} //assign elements to a slice of type integer
	fmt.Println(s[1])   //print value in the second element of slice ..s[0] is the first element of the slice
	s[1] = 15           //assign value of 15 to s[1] of slice
	fmt.Println(s[1])

	var s1 []int                                                 //initialize slice somewhere else in the code
	s1 = make([]int, 5, 25)                                      //make slice ; built in fuction - associated with length and capacity that slice points to
	s1 = append(s1, 1)                                           //built-in fuction to append an element to a slice
	s1 = append(s1, s...)                                        //append another slice ..in this case it is slice "s" syntax is "s..."
	fmt.Println(s1, "len is:", len(s1), "capacity is:", cap(s1)) //"len" and "cap" are built-in functions
	/*

	   Output of the above code..note that "1" is appended and "1, 15 and 4" from slice "s" is then appended again

	    go run src/slices1.go
	   2
	   15
	   [0 0 0 0 0 1 1 15 4] len is: 9 capacity is: 25
	*/

	s2 := s1[0:5] //create another slice "s2" based on "s1" and assign elements 0 to 5 (start and end index)
	fmt.Println(s2)
	copy(s2, s1[4:8]) //"copy" is a built-in function that copies elements of an array to another array
	fmt.Println(s2)
}
