package error_handling

import (
	"errors"
	"fmt"
)

func tahminet(tahmin int) (string, error) {

	if tahmin < 1 || tahmin > 100 {
		return "", errors.New("1-100 arasında bir sayı giriniz")

	}
	return "bildiniz", nil
}

func Demo2() {
	mesaj, _ := tahminet(50)
	fmt.Println(mesaj)
	fmt.Println(tahminet(102))
	fmt.Println(tahminet(-10))

}
