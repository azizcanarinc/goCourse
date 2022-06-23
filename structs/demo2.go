package structs

import "fmt"

type customer struct {
	firsname string
	lastname string
	age      int
}

func (c customer) update() {
	fmt.Println("güncellendi : ", c.firsname)
}

func (c customer) save() {
	fmt.Println("Eklendi : ", c.firsname)
}

func Demo2() {
	c := customer{firsname: "Aziz Can", lastname: "Arinç", age: 13}
	c.save()
	c.update()
}
