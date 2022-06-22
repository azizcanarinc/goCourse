package for_range

import "fmt"

func Demo3() {
	sözlük := map[string]string{"book : ": "kitap", "table : ": "masa"}
	for k, v := range sözlük {
		fmt.Print(k)
		fmt.Println(v)
	}
}
