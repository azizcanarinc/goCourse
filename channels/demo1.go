package channels

func Teksayılar(Teksayıcn chan int) {
	toplam := 0
	for i := 1; i <= 10; i += 2 {
		toplam = toplam + i
	}
	Teksayıcn <- toplam
}

func Çiftsayılar(Çiftsayıcn chan int) {
	toplam := 0
	for i := 0; i <= 10; i += 2 {
		toplam = toplam + i
	}
	Çiftsayıcn <- toplam
}
