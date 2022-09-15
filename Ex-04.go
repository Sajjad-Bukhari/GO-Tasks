// Sajjad Bukhari
// 19I-1686
// CySec-M
// Example # 04

package main

import "fmt"

func main() {
	elements := []int{999, 99, 9}

	// Loop over the int slice and Print its elements.
	// ... No newline is inserted after Print.

	for i := 0; i < len(elements); i++ {
		fmt.Print(elements[i] + 1)
		fmt.Print(" ")
	}
	fmt.Println("... DONE!")
}