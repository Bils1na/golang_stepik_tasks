// Напишите функцию, которая умножает
// значения на которые ссылаются два
// указателя и выводит получившееся
// произведение в консоль. Ввод данных
// уже написан за вас.

// Sample Input:
// 2 2
// Sample Output:
// 4

package main

import (
	"fmt"
)

func test(x1 *int, x2 *int) {
	fmt.Print(*x1 * *x2)
}

func main() {
	var a int = 3
	var b int = 2

	test(&a, &b)
}
