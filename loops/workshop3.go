package loops

import "fmt"

func Demo5() {
	sayı1 := 220
	sayı2 := 280
	toplam1 := 0
	toplam2 := 0

	for i := 1; i < sayı1; i++ {
		if sayı1%i == 0 {
			toplam1 = toplam1 + i
		}

	}

	for i := 1; i < sayı2; i++ {

		if sayı2%i == 0 {
			toplam2 = toplam2 + i
		}
	}

	if toplam1 == toplam2 && toplam2 == toplam1 {
		fmt.Printf("%v ve %v arkadaş sayılardır ", sayı1, sayı2)

	} else {
		fmt.Printf("%v ve %v arkadaş sayı deyiller ", sayı1, sayı2)
	}

}
