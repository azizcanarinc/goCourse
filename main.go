package main

import (
	"fmt"
	"golesson/conditionals"
	"golesson/degisgenler"
	"golesson/loops"
)

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
}
