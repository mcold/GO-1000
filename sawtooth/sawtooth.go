package main

/* Пилообразная последовательность
 (Время: 2 сек. Память: 128 Мб Сложность: 38%)
 Последовательность a1, a2, a3, … , an-1, an называется пилообразной, если она удовлетворяет одному из следующих условий:
1) a1 < a2 > a3 < … > an-1 < an
2) a1 < a2 > a3 < … < an-1 > an
3) a1 > a2 < a3 > … < an-1 > an
4) a1 > a2 < a3 > … > an-1 < an
Дана числовая последовательность. Требуется определить длину самой длинной ее пилообразной непрерывной подпоследовательности.
*/

import (
	"fmt"
	"math"
)

func main() {
	s := []int{5, 7, 5, 5, 4, 6, 7, 1, 8, 9, 4, 5} // 3
	// s := []int{5, 7, 12, 5, 4, 6, 7, 1, 8, 9, 4, 5} // 0
	// s := []int{5, 7, 4, 5, 4, 6, 7, 1, 8, 9, 4, 5} // 5
	// s := []int{5, 7, 4, 5, 4, 6, 3, 12, 8, 9, 4, 5} // 12
	s1 := s[1:]
	nRes := 0

	// save the items of slice in map
	m := make(map[int]bool)
	for i := 0; i < len(s); i++ {
		m[i] = true
	}

	// если длина слайса <= 1 то возвращаем длину слайса
	if len(s) <= 1 {
		fmt.Println(len(s))
		return
	}

	for order, num := range s1 {
		bLeft := math.Signbit(float64(s[order] - num))
		if _, ok := m[order+2]; ok {
			bRight := math.Signbit(float64(s[order+2] - num))
			if bLeft == bRight {
				nRes = order + 2
			} else {
				break
			}

		} else {
			nRes = order + 2
			break
		}
	}

	fmt.Println(nRes)
}