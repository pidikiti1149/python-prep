// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

func main() {
	// 1. uncomment the code below
	png, header := []byte{'P', 'N', 'G'}, []byte{}

	header = append(header, png...)

	// 2. append elements to header to make it equal with the png slice

	// 3. compare the slices using the bytes.Equal function

	// 4. print whether they're equal or not

	if len(png) == len(header) {
		for i := range png {
			if png[i] != header[i] {
				return
			}
		}
		fmt.Println("They are equal.")
	}
}
