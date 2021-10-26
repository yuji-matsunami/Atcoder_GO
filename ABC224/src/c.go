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

func IntStdin() (i int) {
	stringInput := StrStdin()

	i, _ = strconv.Atoi(stringInput)
	return
}

func main() {
	var N int
	N = IntStdin()
	var points [][]int = make([][]int, N)
	for i:=0; i<N; i++ {
		sl_int := SplitIntStdin(" ")
		points[i] = sl_int
	}
	var ans int = 0
	for i:=0; i<N; i++ {
		for j:=i+1; j<N; j++ {
			for k:=j+1; k<N; k++ {
				var h1, h2 float64
				if points[i][0] < points[j][0]{
					h1 = (float64(points[j][1] - points[i][1])) / float64((points[j][0] - points[i][0]))
				}else {
					h1 = (float64(points[i][1] - points[j][1])) / float64((points[i][0] - points[j][0]))
				}

				if points[k][0] < points[j][0]{
					h2 = (float64(points[j][1] - points[k][1])) / float64((points[j][0] - points[k][0]))
				}else {
					h2 = (float64(points[k][1] - points[j][1])) / float64((points[k][0] - points[j][0]))
				}

				

				if h1 == h2 {
					continue
				}else if points[i][0] == points[j][0] && points[j][0] == points[k][0] {
					continue
				}
				ans++
				//h1 = (points[i][1] - points[j][1])
			}
		}
	}
	fmt.Println(ans)
}