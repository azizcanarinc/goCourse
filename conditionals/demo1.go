package conditionals

import "fmt"

func Demo1() {
	var hesap float64 = 1000
	var cekilmekIstene float64 = 900

	if cekilmekIstene > hesap {
		fmt.Print("hesaptaki para yetersiz")
	}

	if cekilmekIstene <= hesap {
		fmt.Println("paranız hazırlanıyor")
		hesap = hesap - cekilmekIstene

	}
	fmt.Println("Bitti. hesaptaki para : " + fmt.Sprintf("%v", hesap))
}
