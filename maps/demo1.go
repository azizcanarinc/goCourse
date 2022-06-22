package maps

import "fmt"

func Demo1() {
	sözlük := make(map[string]string)
	sözlük["table"] = "masa"
	sözlük["book"] = "kitap"
	sözlük["pencil"] = "kalem"

	fmt.Println(sözlük["book"])
	fmt.Println("Eleman sayısı : ", len(sözlük))
	fmt.Println("sözlük : ", sözlük)
	fmt.Println()
	delete(sözlük, "book")
	fmt.Println("Eleman sayısı : ", len(sözlük))
	fmt.Println("sözlük : ", sözlük)
	fmt.Println()

	değer, varmı := sözlük["pencil"]
	fmt.Println(değer)
	fmt.Println("listede olma durumu : ", varmı)
	sözlük2 := map[string]string{"glass": "bardak", "ston": "taş"}
	fmt.Println(sözlük2)
}
