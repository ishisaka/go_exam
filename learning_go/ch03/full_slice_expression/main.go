// サブスライスの容量を変更する
package main

import "fmt"

func main() {
	x := make([]string, 0, 5)
	x = append(x, "a", "b", "c", "d")
	y := x[:2:2]                        // yはxの最初の2つの要素を参照し、容量は2
	z := x[2:4:4]                       // zはxの3番目と4番目の要素を参照し、容量は4
	fmt.Println(cap(x), cap(y), cap(z)) // 5 2 2
	y = append(y, "i", "j", "k")
	x = append(x, "x")
	z = append(z, "y")
	fmt.Println("x:", x) // x: [a b c d x]
	fmt.Println("y:", y) // y: [a b i j k]
	fmt.Println("z:", z) // z: [c d y]
}
