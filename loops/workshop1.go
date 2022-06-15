package loops

import "fmt"

func Demo3() {
	aklımdakisayı := 80
	tahminedilensayı := 0
	tahminsayısı := 0
	fmt.Println("aklımdan tuttuğum sayıyı tahmin et 1 ila 100 arasında")
	fmt.Scanln(&tahminedilensayı)
	tahminsayısı = tahminsayısı + 1
	for aklımdakisayı != tahminedilensayı {

		if aklımdakisayı < tahminedilensayı {
			fmt.Println("daha küçük bir sayı girin ):")
			fmt.Scanln(&tahminedilensayı)
			tahminsayısı = tahminsayısı + 1
		}

		if aklımdakisayı > tahminedilensayı {
			fmt.Println("daha büyük bir sayı girin ):")
			fmt.Scanln(&tahminedilensayı)
			tahminsayısı = tahminsayısı + 1

		}
	}
	başarıdurumu := ""
	if tahminsayısı > 3 && tahminsayısı <= 5 {
		başarıdurumu = "güzel"
	}

	if tahminsayısı > 0 && tahminsayısı <= 3 {
		başarıdurumu = "süper saladın (;"

	}

	if tahminsayısı > 5 && tahminsayısı <= 7 {
		başarıdurumu = "eh işte"
	}

	if tahminsayısı > 7 {
		başarıdurumu = "kötü.."
	}

	fmt.Printf("Döğru Sayı (; %v tahminde : %v", tahminsayısı, başarıdurumu)
}
