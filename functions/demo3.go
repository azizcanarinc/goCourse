package functions

func Toplavariadic(sayılar ...int) int {
	toplam := 0
	for i := 0; i < len(sayılar); i++ {
		toplam = toplam + sayılar[i]

	}
	return toplam

}
