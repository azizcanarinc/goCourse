package loops

import "fmt"

func Demo1() {
	var metin string = "merhaba dünya!"

	i := 1

	for i <= 5 {
		fmt.Println(metin)
		i = i + 1
	}
	fmt.Println("Bitti")
}
