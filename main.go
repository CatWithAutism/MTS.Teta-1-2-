package main

import (
	"fmt"
	"sort"
	"strings"
)

func archiveString(s string) string {
	//Счетчик символов в строке
	symbolCounter := make(map[rune]int)
	for _, r := range s {
		symbolCounter[r] += 1
	}

	//Вытаскиваем ключи и конвертим руну к инту для сортировки по аски
	runes := make([]int, 0, len(symbolCounter))
	for r := range symbolCounter {
		runes = append(runes, int(r))
	}
	sort.Ints(runes)

	//Сб наш друг на длинные строки
	sb := strings.Builder{}
	for _, r := range runes {
		countOfSymbol := symbolCounter[rune(r)]
		if countOfSymbol > 1 {
			sb.WriteString(fmt.Sprintf("%s%d", string(r), countOfSymbol))
		} else {
			sb.WriteString(fmt.Sprintf("%s", string(r)))
		}

	}

	return sb.String()
}

func main() {
	println(archiveString("ябд"))
	println(archiveString("ЯЯЯБББддда"))
	println(archiveString("aaabbbcccccaaaaa"))
	println(archiveString("aaabbbccccc"))
	println(archiveString("zzzzcccUUUuu"))
} 
