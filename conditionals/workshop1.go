package conditionals

import "fmt"

func Demo3() {
	var afab int = 500
	var bfab int = 399
	var cfab int = 700

	var enBuyuk int = afab

	if bfab > enBuyuk {
		enBuyuk = bfab

	}
	if cfab > enBuyuk {
		enBuyuk = cfab

	}

	fmt.Printf("En Büyük Sayı: %v", enBuyuk)
}
