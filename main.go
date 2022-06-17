package main

import (
	"fmt"
	"golesson/arrays"
	"golesson/conditionals"
	"golesson/degisgenler"
	"golesson/functions"
	"golesson/loops"
	"golesson/slices"
)

//komudları çalıştırmak için önündeki işaretleri  "//" sil
func main() {
	degisgenler.Demo1()
	fmt.Println()
	fmt.Println()
	fmt.Println("ŞARTLAR DEMO1")
	fmt.Println()
	conditionals.Demo1()
	fmt.Println()
	fmt.Println()
	fmt.Println("İF ELSE DEMO2")
	fmt.Println()
	conditionals.Demo2()
	fmt.Println()
	fmt.Println()
	fmt.Println("3 değişkenden en büyük olanın değerini yaz")
	fmt.Println()
	conditionals.Demo3()
	fmt.Println()
	fmt.Println()
	fmt.Println("döngüler 1")
	fmt.Println()
	loops.Demo1()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println("döngüler 2")
	fmt.Println()
	fmt.Println()
	loops.Demo2()
	fmt.Println("oyun 1")
	fmt.Println()
	fmt.Println()
	loops.Demo3()
	fmt.Println()
	fmt.Println()
	fmt.Println("oyun 2")
	fmt.Println()
	fmt.Println()
	loops.Demo4()
	fmt.Println()
	fmt.Println()
	fmt.Println("DİZİLİER 1")
	fmt.Println()
	fmt.Println()
	arrays.Demo1()
	fmt.Println()
	fmt.Println()
	fmt.Println("DİZİLER 2")
	fmt.Println()
	fmt.Println()
	arrays.Demo2()
	fmt.Println()
	fmt.Println()
	fmt.Println("DİZİLER3")
	fmt.Println()
	fmt.Println()
	arrays.Demo3()
	fmt.Println()
	fmt.Println()
	fmt.Println("DİZİLER 4")
	fmt.Println()
	fmt.Println()
	arrays.Demo4()
	fmt.Println()
	fmt.Println()
	fmt.Println("slice")
	fmt.Println()
	fmt.Println()
	slices.Demo1()
	fmt.Println()
	fmt.Println()
	fmt.Println("slice2")
	fmt.Println()
	fmt.Println()
	slices.Demo2()
	fmt.Println()
	fmt.Println()
	fmt.Println("functions")
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println("")
	fmt.Println()
	fmt.Println()
	functions.Selamver("eylül")
	functions.Selamver("semih")
	functions.Topla(2, 6)
	var sonuç = functions.Topla(3, 6)
	fmt.Println(sonuç)
	fmt.Println()
	fmt.Println()
	fmt.Println("")
	fmt.Println()
	fmt.Println()
	fmt.Println()

	sonuç1, sonuç2, sonuç3, sonuç4 := functions.Dörtişlem(10, 6)
	fmt.Println("toplam : ", sonuç1)
	fmt.Println("çıkartma : ", sonuç2)
	fmt.Println("çarpım : ", sonuç3)
	fmt.Println("bölme : ", sonuç4)
	fmt.Println()
	fmt.Println("")
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println(functions.Toplavariadic(1, 4, 6, 3, 10))
	fmt.Println(functions.Toplavariadic(1, 10))
	fmt.Println(functions.Toplavariadic())

	sayılar := []int{4, 7, 8}
	fmt.Println(functions.Toplavariadic(sayılar...))
}
