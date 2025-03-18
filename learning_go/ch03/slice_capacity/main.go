// スライスのキャパシティについて
package main

import "fmt"

func main() {
	// スライスのキャパシティは、スライスが参照する配列の長さを超えた場合に自動的に増加する
	var x []int
	fmt.Println(x, len(x), cap(x)) // [] 0 0
	x = append(x, 10)
	fmt.Println(x, len(x), cap(x)) // [10] 1 1
	x = append(x, 20)
	fmt.Println(x, len(x), cap(x)) // [10 20] 2 2
	x = append(x, 30)
	fmt.Println(x, len(x), cap(x)) // [10 20 30] 3 4
	x = append(x, 40)
	fmt.Println(x, len(x), cap(x)) // [10 20 30 40] 4 4
	x = append(x, 50)
	fmt.Println(x, len(x), cap(x)) // [10 20 30 40 50] 5 8

	x1 := make([]int, 5) // 長さ5、キャパシティ5のスライスを作成
	x1 = append(x1, 10)
	fmt.Println(x1, len(x1), cap(x1)) // [0 0 0 0 0 10] 6 10

	x2 := make([]int, 0, 10) // 長さ0、キャパシティ10のスライスを作成
	x2 = append(x2, 5, 6, 7, 8)
	fmt.Println(x2, len(x2), cap(x2)) // [5 6 7 8] 4 10

}
