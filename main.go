package main

import (
	"fmt"
	"golesson/arrays"
	"golesson/channels"
	"golesson/conditionals"
	"golesson/defer_statement"
	"golesson/degisgenler"
	"golesson/error_handling"
	"golesson/for_range"
	"golesson/functions"
	"golesson/goroutines"
	"golesson/interfaces"
	"golesson/loops"
	"golesson/maps"
	"golesson/pointers"
	"golesson/slices"
	"golesson/string_functions"
	"golesson/structs"
	"time"
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
	// loops.Demo3()
	fmt.Println()
	fmt.Println()
	fmt.Println("oyun 2")
	fmt.Println()
	fmt.Println()
	// loops.Demo4()
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
	fmt.Println()
	fmt.Println()
	fmt.Println("")
	fmt.Println()
	fmt.Println()
	fmt.Println()
	maps.Demo1()
	fmt.Println()
	fmt.Println()
	fmt.Println("")
	fmt.Println()
	fmt.Println()
	fmt.Println()
	for_range.Demo1()
	fmt.Println()
	fmt.Println("")
	fmt.Println()
	fmt.Println()
	for_range.Demo2()
	fmt.Println()
	fmt.Println("")
	fmt.Println()
	fmt.Println()
	for_range.Demo3()
	fmt.Println()
	fmt.Println("")
	fmt.Println()
	fmt.Println()
	sayı := 20
	pointers.Demo1(sayı)
	fmt.Println("maindeki sayı", sayı)
	fmt.Println()
	fmt.Println("")
	fmt.Println()
	fmt.Println()
	sayılar2 := []int{1, 2, 3}
	pointers.Demo2(sayılar2)
	fmt.Println("maindeki sayı : ", sayılar2[0])
	fmt.Println()
	fmt.Println("")
	fmt.Println()
	fmt.Println()
	structs.Demo1()
	fmt.Println()
	fmt.Println("")
	fmt.Println()
	fmt.Println()
	structs.Demo2()
	fmt.Println()
	fmt.Println("")
	fmt.Println()
	fmt.Println()
	go goroutines.Çiftsayılar()
	go goroutines.Teksayılar()

	time.Sleep(5 * time.Second)
	fmt.Println("main bitti")
	fmt.Println()
	fmt.Println("")
	fmt.Println()
	fmt.Println()
	Çiftsayıcn := make(chan int)
	teksayıcn := make(chan int)
	go channels.Çiftsayılar(Çiftsayıcn)
	go channels.Teksayılar(teksayıcn)
	Çiftsayıtoplam, teksayıtoplam := <-Çiftsayıcn, <-teksayıcn
	çarpım := Çiftsayıtoplam * teksayıtoplam
	fmt.Println("çarpım : ", çarpım)
	fmt.Println()
	fmt.Println("şekilin alanını hesaplama")
	fmt.Println()
	fmt.Println()
	interfaces.Demo1()
	fmt.Println()
	fmt.Println("kredi")
	fmt.Println()
	fmt.Println()
	interfaces.Demo2()
	fmt.Println()
	fmt.Println("")
	fmt.Println()
	fmt.Println()
	defer_statement.B()
	fmt.Println()
	fmt.Println("")
	fmt.Println()
	fmt.Println()
	defer_statement.Demo2(10)
	fmt.Println()
	fmt.Println("")
	fmt.Println()
	fmt.Println()
	defer_statement.Demo3()
	fmt.Println()
	fmt.Println("")
	fmt.Println()
	fmt.Println()
	error_handling.Demo1()
	fmt.Println()
	fmt.Println("")
	fmt.Println()
	fmt.Println()
	interfaces.Demo3()
	fmt.Println()
	fmt.Println("")
	fmt.Println()
	fmt.Println()
	error_handling.Demo2()
	fmt.Println()
	fmt.Println("")
	fmt.Println()
	fmt.Println()
	string_functions.Demo1()

}
