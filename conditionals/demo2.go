package conditionals

import "fmt"

func Demo2() {
	var hesap float64 = 1000
	var cekilmekIstene float64 = 900

	if cekilmekIstene > hesap {
		fmt.Println("hesaptaki para yetersiz")
	} else if cekilmekIstene == hesap {
		fmt.Println("paranız hazırlanıyor")
		fmt.Println("hesabınızda para kalmadı")
	} else {
		fmt.Println("paranız hazırlanıyor")
	}

}
