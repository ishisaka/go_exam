// スライスのコピー
package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4}
	d := [4]int{5, 6, 7, 8}
	y := make([]int, 2)
	copy(y, d[:])  // 配列からスライスのコピー
	fmt.Println(y) // [5 6]
	copy(d[:], x)  // スライスから配列のコピー
	fmt.Println(d) // [1 2 3 4]
}
