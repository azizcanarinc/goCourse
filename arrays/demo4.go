package arrays

import "fmt"

func Demo4() {
	var sayılar [2][3]int

	sayılar[0][0] = 1
	sayılar[0][1] = 3
	sayılar[0][2] = 5
	sayılar[1][0] = 2
	sayılar[1][1] = 4
	sayılar[1][2] = 6

	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print(sayılar[i][j])
			fmt.Print("")
		}
		fmt.Println()
	}
}
