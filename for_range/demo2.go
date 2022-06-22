package for_range

import "fmt"

func Demo2() {

	sayılar := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	toplam := 0
	for _, sayı := range sayılar {
		if sayı%2 != 0 {
			toplam = toplam + sayı
		}

	}
	fmt.Println("toplamı : ", toplam)
}
