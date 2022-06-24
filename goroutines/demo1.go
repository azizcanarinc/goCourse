package goroutines

import (
	"fmt"
	"time"
)

func Teksayılar() {
	for i := 1; i <= 10; i += 2 {
		fmt.Println("tek sayılar : ", i)
		time.Sleep(1 * time.Second)
	}
}

func Çiftsayılar() {
	for i := 0; i <= 10; i += 2 {
		fmt.Println("çift sayılar : ", i)
		time.Sleep(1 * time.Second)
	}
}
