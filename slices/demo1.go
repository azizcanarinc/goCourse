package slices

import "fmt"

func Demo1() {
	isimler := make([]string, 3)
	fmt.Println(isimler)
	isimler[0] = "engin"
	isimler[1] = "derin"
	isimler[2] = "salih"
	isimler = append(isimler, "büşra", "ahmet")
	fmt.Println(isimler)
	fmt.Println(len(isimler))
}
