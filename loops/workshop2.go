package loops

import (
	"fmt"
)

//kullanıcıdan sayı iste
// 23 : asaldır

func Demo4() {
	istenensayı := 0
	fmt.Println("bir sayı söyle asalmı deyilmi soyliyim")
	fmt.Scanln(&istenensayı)

	asalmı := true
	for i := 2; i < istenensayı; i++ {
		if istenensayı%i == 0 {
			asalmı = false
		}
	}
	if asalmı == true {
		fmt.Println("asal sayı")

	}
	if asalmı == false {
		fmt.Println("asal sayı deyil")

	}
}
