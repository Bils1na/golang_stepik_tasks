// На вход подается строка. Нужно определить,
// \является ли она правильной или нет.
// Правильная строка начинается с заглавной
// буквы и заканчивается точкой. Если строка
// правильная - вывести Right иначе - вывести Wrong

// Маленькая подсказка: fmt.Scan() считывает
// строку до первого пробела, вы можете считать
// полностью строку за раз с помощью bufio:
// text, _ := bufio.NewReader(os.Stdin).ReadString('\n')

// Sample Input:
// Быть или не быть.
// Sample Output:
// Right

package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	runeText := []rune(text)
	if unicode.IsUpper(runeText[0]) && runeText[len(runeText)-1] == '.' {
		fmt.Print("Right")
	} else {
		fmt.Print("Wrong")
	}
}
