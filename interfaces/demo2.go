package interfaces

import "fmt"

type mortgage struct {
	creditpaymenttotal float64
	address            string
	rate               float64
}

type car struct {
	creditpaymenttotal float64
	carinfo            string
	rate               float64
}

type creditcalculator interface {
	calculate() float64
}

func (m mortgage) calculate() float64 {
	return m.creditpaymenttotal * m.rate / 100 / 12
}

func (c car) calculate() float64 {
	return c.creditpaymenttotal * c.rate / 100 / 12
}

func CalculateMonthlypayment(credits []creditcalculator) float64 {
	paymenttotal := 0.0
	for i := 0; i < len(credits); i++ {
		paymenttotal = paymenttotal + credits[i].calculate()
	}
	return paymenttotal
}

func Demo2() {
	credit1 := mortgage{rate: 10, creditpaymenttotal: 100000, address: "tokat"}
	credit2 := mortgage{rate: 12, creditpaymenttotal: 50000, address: "istambul"}
	credit3 := car{rate: 15, creditpaymenttotal: 200000, carinfo: "BMW"}
	credits := []creditcalculator{credit1, credit2, credit3}
	total := CalculateMonthlypayment(credits)
	fmt.Println("toplam ödeme : ", total)
}
