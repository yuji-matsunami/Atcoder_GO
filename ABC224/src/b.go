package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
var (
	scanner = bufio.NewScanner(os.Stdin)
)
func StrStdin() (stringInput string) {

	scanner.Scan()
	stringInput = scanner.Text()

	stringInput = strings.TrimSpace(stringInput)
	// fmt.Println(stringInput)
	return
}

func SplitWithoutEmpty(stringTargeted string, delim string) (stringReturned []string) {
	stringSplited := strings.Split(stringTargeted, delim)

	for _, str := range stringSplited {
		if str != "" {
			stringReturned = append(stringReturned, str)
		}
	}
	return
}

func SplitStrStdin(delim string) (stringReturned []string) {
	stringInput := StrStdin()

	stringReturned = SplitWithoutEmpty(stringInput, delim)

	return
}

func SplitIntStdin(delim string) (intSlices []int) {
	stringSplited := SplitStrStdin(" ")

	for i := range stringSplited {
		var iparam int
		iparam, _ = strconv.Atoi(stringSplited[i])
		intSlices = append(intSlices, iparam)
	}

	return
}
func main() {
	var H, W int
	
	i := SplitIntStdin(" ")
	
	H = i[0]
	W = i[1]
	
	
	// fmt.Scan(&H, &W)
	// fmt.Println(H, W)
	graph := make([][]int, H)
	// fmt.Println(graph)
	for i:=0; i<H; i++ {
		j := SplitIntStdin(" ")
		// fmt.Println(j)
		// fmt.Println("i:", i)
		graph[i] = j
	}
	// fmt.Println(W)
	// fmt.Println(graph)
	var yesno bool = true
	Loop:
	for i1:=0; i1<H; i1++ {
		for j1:=0; j1<W; j1++ {
			for i2:=i1+1; i2<H; i2++ {
				for j2:= j1+1; j2<W; j2++ {
					if graph[i1][j1] + graph[i2][j2] <= graph[i2][j1] + graph[i1][j2]{
						continue
					}else{
						yesno = false
						break Loop
					}
				}
			}

		}
	}

	if yesno {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}