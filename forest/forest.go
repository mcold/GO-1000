package main

// https://acmp.ru/index.asp?main=task&id_task=24

import "fmt"

var n, m int

func main() {
	n = 5
	m = 2
	var nDelta, nVars int

	for delta := 1; delta <= n-m+1; delta++ {
		if delta * (m - 1) + 1 <= n {
			// number of elements using this delta
			nDelta = (m - 1) * delta + 1
			// number of variants
			nVars = nVars + (n - nDelta) + 1
			}
		}
	fmt.Println(nVars)
	}
