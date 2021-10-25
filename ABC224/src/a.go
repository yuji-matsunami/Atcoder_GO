package main

import "fmt"

func main() {
	var S string
	fmt.Scan(&S)

	if S[len(S)-1] == 'r' {
		fmt.Println("er")
	} else {
		fmt.Println("ist")
	}
}