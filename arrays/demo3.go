package arrays

import "fmt"

func Demo3() {
	sayılar := [5]int{20, 30, 10, 50, 2}
	fmt.Println(sayılar)

	enbüyük := sayılar[0]

	for i := 0; i < len(sayılar); i++ {
		if sayılar[i] > enbüyük {
			enbüyük = sayılar[i]
		}
	}
	fmt.Println("enbüyük sayı : ", enbüyük)
}
