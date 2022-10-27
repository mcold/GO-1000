package main

/*
https://acmp.ru/index.asp?main=task&id_task=19
*/

import (
	"fmt"
	"sort"
	"strconv"
)

type figure struct {
	x int
	y int
}

func getFigure(coords string, m map[rune]int32) figure{
	x := m[[]rune(coords)[0]]
	y, _ := strconv.Atoi(coords[1:])
	f := figure{int(x), y}
	return f
}

func getField(x int, y int) []string {
	sField := make([]string, 0)
	for i := 1; i <= x; i++ {
		for j := 1; j <= y; j++ {
			sField = append(sField, strconv.Itoa(i) + strconv.Itoa(j))
		}
	}
	return sField
}

// a b c d e f g h
func addVertical(f *figure, sFields *[] string) []string{
	sF := *sFields
	sResult := sF[0:]
	for i := 1; i <= 8; i++ {
		sResult = append(sResult, strconv.Itoa(f.x) + strconv.Itoa(i))
	}
	return sResult
}

func addHorizontal(f *figure, sFields *[] string) []string{
	sF := *sFields
	sResult := sF[0:]
	for i := 1; i <= 8; i++ {
		sResult = append(sResult, strconv.Itoa(i) + strconv.Itoa(f.y))
	}
	return sResult
}

func addAcross(f *figure, sFields *[] string) []string{
	sF := *sFields
	sResult := sF[0:]
	xNew := f.x
	yNew := f.y
	for i := 1; i <= 7; i++ {
		xNew = f.x - i
		if xNew > 0 && xNew <=8 {
			yNew = f.y - i
			if yNew > 0 && yNew <= 8 {
				sResult = append(sResult, strconv.Itoa(xNew) + strconv.Itoa(yNew))
			}

			yNew = f.y + i
			if yNew > 0 && yNew <= 8 {
				sResult = append(sResult, strconv.Itoa(xNew) + strconv.Itoa(yNew))
			}
		}

		xNew = f.x + i
		if xNew > 0 && xNew <=8 {
			yNew = f.y - i
			if yNew > 0 && yNew <= 8 {
				sResult = append(sResult, strconv.Itoa(xNew) + strconv.Itoa(yNew))
			}

			yNew = f.y + i
			if yNew > 0 && yNew <= 8 {
				sResult = append(sResult, strconv.Itoa(xNew) + strconv.Itoa(yNew))
			}
		}
	}
	return sResult
}

func addKnightStep(f *figure, sFields *[] string) []string{
	sF := *sFields
	sResult := sF[0:]
	fFigure := *f
	xNew := fFigure.x - 1
	yNew := fFigure.y
	if xNew > 0 && xNew <= 8 {
		yNew = fFigure.y - 2
		if yNew > 0 && yNew <= 8 {
			sResult = append(sResult, strconv.Itoa(xNew) + strconv.Itoa(yNew))
		}

		yNew = fFigure.y + 2
		if yNew > 0 && yNew <= 8 {
			sResult = append(sResult, strconv.Itoa(xNew) + strconv.Itoa(yNew))
		}
	}

	xNew = fFigure.x + 1
	if xNew > 0 && xNew <= 8 {
		yNew = fFigure.y - 2
		if yNew > 0 && yNew <= 8 {
			sResult = append(sResult, strconv.Itoa(xNew) + strconv.Itoa(yNew))
		}

		yNew = fFigure.y + 2
		if yNew > 0 && yNew <= 8 {
			sResult = append(sResult, strconv.Itoa(xNew) + strconv.Itoa(yNew))
		}
	}

	yNew = fFigure.y - 1
	if yNew > 0 && yNew <= 8 {
		xNew = fFigure.x - 2
		if xNew > 0 && xNew <= 8 {
			sResult = append(sResult, strconv.Itoa(xNew) + strconv.Itoa(yNew))
		}

		xNew = fFigure.x + 2
		if xNew > 0 && xNew <= 8 {
			sResult = append(sResult, strconv.Itoa(xNew) + strconv.Itoa(yNew))
		}
	}

	yNew = fFigure.y + 1
	if yNew > 0 && yNew <= 8 {
		xNew = fFigure.x - 2
		if xNew > 0 && xNew <= 8 {
			sResult = append(sResult, strconv.Itoa(xNew) + strconv.Itoa(yNew))
		}

		xNew = fFigure.x + 2
		if xNew > 0 && xNew <= 8 {
			sResult = append(sResult, strconv.Itoa(xNew) + strconv.Itoa(yNew))
		}
	}
	return sResult
}

func getUniqueSliceString(s *[]string) []string{
	sNotUnique := *s
	sUnique := make([]string, 0)
	m := make(map[string] bool)
	for i := range sNotUnique {
		m[sNotUnique[i]] = true
	}

	for k, _ := range m {
		sUnique = append(sUnique, k)
	}

	return sUnique
}

func translateNumToFields(s *[]string) []string{
	sF := *s
	sResult := make([]string, 0)
	m := map[int]string{1: "A", 2: "B", 3: "C", 4: "D", 5: "E", 6: "F", 7: "G", 8: "H"}
	for _, v := range sF {
		newNum, _ := strconv.Atoi(v[0:1])
		sResult = append(sResult, m[newNum] + v[1:])
	}
	return sResult
}

func main() {
	m := make(map[rune]int32)
	sFields := make([]string, 0)
	pF := &sFields

	fdQueen := "D5"
	fdRock := "A7"
	fdKnight := "G2"

	for ch := 'A'; ch <= 'H'; ch++ {
		m[ch] = ch - 64 // 64 - literal before A use as default point
	}
	Q := getFigure(fdQueen, m)
	R := getFigure(fdRock, m)
	K := getFigure(fdKnight, m)
	pQ := &Q
	pR := &R
	pK := &K
	// queen
	sFields = addHorizontal(pQ, pF)
	sFields = addVertical(pQ, pF)
	sFields = addAcross(pQ, pF)

	// rock
	sFields = addHorizontal(pR, pF)
	sFields = addVertical(pR, pF)
	// knight
	sFields = addKnightStep(pK, pF)

	sFields = translateNumToFields(pF)

	sFields = append(sFields, fdQueen, fdRock, fdKnight)
	sFields = getUniqueSliceString(pF)
	fmt.Println(len(sFields))
	sort.Sort(sort.StringSlice(sFields))
	for _, v := range sFields {
		fmt.Println(v)
	}
}
