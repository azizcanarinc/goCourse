package error_handling

import (
	"fmt"
	"os"
)

func Demo1() {
	f, err := os.Open("demo11.txt")

	if err != nil {
		if pErr, ok := err.(*os.PathError); ok {
			fmt.Println("dosya bulunamad─▒ : ", pErr.Path)
			return
		} else {
			fmt.Println("dosya bulunamad─▒")
			return
		}

	} else {
		fmt.Println(f.Name())
	}

}
