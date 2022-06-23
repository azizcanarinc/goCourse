package structs

import "fmt"

func Demo1() {
	fmt.Println(product{"bilgiseyar", 500000, "xyz"})
	fmt.Println(product{name: "bilgiseyar", unitprice: 50000})
}

type product struct {
	name      string
	unitprice float32
	brand     string
}
