package defer_statement

import "fmt"

func A() {
	fmt.Println("A fonksiyonunu çalıştır")
}

func B() {
	defer A()
	defer C()
	defer D()

	fmt.Println("B fonksiyonunu çalıştır")

}

func C() {
	fmt.Println("C fonksiyonunu çalıştır")

}
func D() {
	fmt.Println("D fonksiyonunu çalıştır")

}
