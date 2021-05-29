package main

import (
	"fmt"
)

type arr []int

func GetArr(n int) []int {
	var arr []int
	for i := 1; i < n+1; i++ {
		arr = append(arr, i)
	}
	return arr
}

func Count(arrMinus *arr, arrRes *arr, cnt *int, iter int) arr {
	var NewArr arr
	var n int
	if iter == 1 {
		for j := 0; j < len(*arrMinus); j++ {
			n = len(*arrRes) - (*arrMinus)[j]
			if n == 0 {
				*cnt = *cnt + 1
			}
			if n > 0 {
				NewArr = append(NewArr, n)
			}
		}
	} else {
		for i := 0; i < len(*arrRes); i++ {
			for j := 0; j < len(*arrMinus); j++ {
				n = (*arrRes)[i] - (*arrMinus)[j]
				if n == 0 {
					*cnt = *cnt + 1
				}
				if n > 0 {
					NewArr = append(NewArr, n)
				}
			}
		}
	}
	return NewArr
}

func CountVariants(n, k int) int {
	var cnt int
	var arrMinus, arrRes arr
	iter := 1
	arrMinus = GetArr(n)
	arrRes = GetArr(k)
	for len(arrRes) > 0 {
		arrRes = Count(&arrMinus, &arrRes, &cnt, iter)
		iter = iter + 1
	}
	return cnt
}

func printArr(pArr *arr) {
	for i := 0; i < len(*pArr); i++ {
		fmt.Println((*pArr)[i])
	}
}

func changeNumber(pNum *int) {
	*pNum = *pNum + 1
}

func main() {
	n := 3
	k := 5
	var res int
	res = CountVariants(n, k)
	fmt.Println(res)

}
