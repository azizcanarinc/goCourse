package project

type product struct {
	Id          int     `json:"id"`
	ProductName string  `json:"productName"`
	Categoryıd  int     `json:"categoryıd"`
	Uniyprice   float64 `json:"unitprice"`
}
type Category struct {
	Id           int `json:"id"`
	Categoryname int `json:"categoryname"`
}

func Get() {

}
