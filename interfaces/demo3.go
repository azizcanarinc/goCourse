package interfaces

import "fmt"

func doğrula(i interface{}) {
	sayı, ok := i.(int)
	fmt.Println(sayı, ok)
}

func Demo3() {
	var sayı1 interface{} = "engin"
	doğrula(sayı1)
	var sayı2 interface{} = 5
	doğrula(sayı2)
}
