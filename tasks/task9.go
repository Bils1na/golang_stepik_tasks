// Заданы три числа - a,b,c(a<b<c) - длины сторон треугольника.
// Нужно проверить, является ли треугольник прямоугольным.
// Если является, вывести "Прямоугольный".
// Иначе вывести "Непрямоугольный"

// Sample Input:
// 6 8 10
// Sample Output:
// Прямоугольный

package main

import (
	"fmt"
)

func main() {
	var (
		a int
		b int
		c int
	)
	fmt.Scan(&a, &b, &c)

	switch {
	case a*a+b*b == c*c:
		fmt.Print("Прямоугольный")
	default:
		fmt.Print("Непрямоугольный")
	}
}
