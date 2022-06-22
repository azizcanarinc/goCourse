package for_range

import "fmt"

func Demo1() {
	şehirler := []string{"ankara", "istambul", "izmir"}
	for i := 0; i < len(şehirler); i++ {
		fmt.Println(şehirler[i])
	}

}
