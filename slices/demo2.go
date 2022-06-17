package slices

import "fmt"

func Demo2() {
	şehirler := []string{"ankara", "istambul", "van", "izmir"}
	fmt.Println(şehirler)
	şehirlerkopya := make([]string, len(şehirler))
	fmt.Println(şehirlerkopya)
	copy(şehirlerkopya, şehirler)
	fmt.Println(şehirlerkopya)
	şehirler = append(şehirler, "adana")
	fmt.Println(şehirler)
	fmt.Println(şehirlerkopya)
	fmt.Println(şehirler[2:4])
	fmt.Println(şehirler[2:])
	fmt.Println(şehirler[:4])
}
