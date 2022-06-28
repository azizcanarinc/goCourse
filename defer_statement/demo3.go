package defer_statement

import "fmt"

type product struct {
	productname string
	unitprice   int
}

func (p product) save() {
	fmt.Println("ürün kaydedildi : ", p.productname)
	defer log()

}
func log() {
	fmt.Println("loglandı")

}

func Demo3() {
	p := product{productname: "laptop", unitprice: 9000}
	p.save()
	fmt.Println("işlem başarılı")

}
