package string_functions

import (
	"fmt"
	s "strings"
)

func Demo2() {
	isim := "Aziz"
	//başlıyormu HasPrefix
	fmt.Println(s.HasPrefix(isim, "Azi"))
	//bitiyormu
	fmt.Println(s.HasSuffix(isim, "az"))

	fmt.Println(s.Index(isim, "i"))
	harfler := []string{"a", "z", "i", "z"}
	sonuç := s.Join(harfler, "*")
	fmt.Println(sonuç)

	fmt.Println(s.Replace(sonuç, "*", "+", 2))
	fmt.Println(s.Split(sonuç, "*"))
	fmt.Println(s.Repeat(sonuç, 5))
}
