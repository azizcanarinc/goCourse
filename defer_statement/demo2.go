package defer_statement

import "fmt"

func Demo2(sayı int) {
	if sayı%2 == 0 {
		fmt.Println("Çift sayıdır")
	}
	if sayı%2 != 0 {
		fmt.Println("Tek sayıdır")
	}

}
