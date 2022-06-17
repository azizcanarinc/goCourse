package functions

func Dörtişlem(sayı1 int, sayı2 int) (int, int, int, float32) {
	toplam := sayı1 + sayı2
	çıkartma := sayı1 - sayı2
	çarpım := sayı1 * sayı2
	bölüm := float32(sayı1 / sayı2)

	return toplam, çıkartma, çarpım, bölüm
}
