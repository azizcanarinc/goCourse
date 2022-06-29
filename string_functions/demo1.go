package string_functions

import (
	"fmt"
	"strings"
)

func Demo1() {
	isim := "aziz"
	fmt.Println(strings.Count(isim, "n"))
	fmt.Println(strings.Contains(isim, "a"))
	fmt.Println(strings.Index(isim, "n"))
	sonuç := strings.Index(isim, "a")

	if sonuç != -1 {
		fmt.Println("a harfi var")
	} else {
		fmt.Println("a harfi yok")
	}

	fmt.Println(strings.ToLower(isim))

	fmt.Println(strings.ToUpper(isim))

}
