// Напишите "функцию голосования", возвращающую то
// значение (0 или 1), которое среди значений ее
// аргументов x, y, z встречается чаще.
// Входные данные
// Вводится 3 числа - x, y и z (x, y и z равны 0 или 1).
// Выходные данные
// Необходимо вернуть значение функции от x, y и z.

// Sample Input:
// 0 0 1
// Sample Output:
// 0

package main

import (
	"fmt"
)

func vote(x int, y int, z int) int {
	if x == y || x == z {
		return x
	} else {
		return y
	}
}

func main() {
	res := vote(1, 0, 1)

	fmt.Println(res)
}
