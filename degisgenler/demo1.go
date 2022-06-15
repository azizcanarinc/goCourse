package degisgenler

import "fmt"

func Demo1() {
	fmt.Print("merhaba ")
	fmt.Println("dünya")
	var metin string = "merhaba türkiye  ___ "
	fmt.Print(metin)
	fmt.Print(metin)
	fmt.Print(metin)
	fmt.Println(metin)

	var kdv int = 10
	fmt.Println(kdv)
	fmt.Println(100 + (100 * 10 / 100))

	var kdv2 float32 = 0.1
	fmt.Println(kdv2)
	fmt.Println(100 + 100*kdv2)

	kdv3 := 25
	fmt.Println(kdv3)
	fmt.Printf("veri tipi : %T\n", kdv3)
	var durum bool
	var metin1 string = "aziz"
	var metin2 string = "alper"
	durum = metin2 == metin1
	fmt.Println(durum)

	var durum2 bool
	var metin3 string = "aziz"
	var metin4 string = "aziz"
	durum2 = metin3 == metin4
	fmt.Println(durum2)
}
