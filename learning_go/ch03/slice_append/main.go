package main

import "fmt"

// main 関数はスライスの操作を示すサンプルプログラムです。
// スライスに要素を追加するために append 関数を使用しています。
// 初めに空のスライスに要素を追加し、その後既存のスライスに複数の要素を追加します。
// 最後に別のスライスの要素を追加する例も示しています。
func main() {

	var x []int
	x = append(x, 10)
	fmt.Println(x) // Output: [10]

	x = []int{1, 2, 3}
	x = append(x, 4)
	fmt.Println(x) // Output: [1 2 3 4]
	x = append(x, 5, 6, 7)
	fmt.Println(x) // Output: [1 2 3 4 5 6 7]
	y := []int{20, 30, 40}
	x = append(x, y...)
	fmt.Println(x) // Output: [1 2 3 4 5 6 7 20 30 40]
}
