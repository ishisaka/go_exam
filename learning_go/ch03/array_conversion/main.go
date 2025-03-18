// 配列の変換
package main

import "fmt"

func main() {
	arrayConversions()
	arrayPointerConversions()
	panicArrayConversions()
}

func arrayConversions() {
	xSlice := []int{1, 2, 3, 4}
	xArray := [4]int(xSlice)
	smallArray := [2]int(xSlice)
	xSlice[0] = 10
	fmt.Println(xSlice)     // [10 2 3 4]
	fmt.Println(xArray)     // [1 2 3 4]
	fmt.Println(smallArray) // [1 2]
}

func arrayPointerConversions() {
	xSlice := []int{1, 2, 3, 4}
	xArrayPointer := (*[4]int)(xSlice)
	xSlice[0] = 10
	xArrayPointer[1] = 20
	fmt.Println(xSlice)        // [10 20 3 4]
	fmt.Println(xArrayPointer) // [10 20 3 4]
}

func panicArrayConversions() {
	xSlice := []int{1, 2, 3, 4}
	panicArray := [5]int(xSlice)
	fmt.Println(panicArray) // Panic!
}
