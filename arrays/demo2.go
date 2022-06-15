package arrays

import "fmt"

func Demo2() {
	var şehirler [5]string
	şehirler[0] = "Ankara"
	şehirler[1] = "İstambul"
	şehirler[2] = "İzmir"
	şehirler[3] = "Adana"
	şehirler[4] = "Diyarbakır"
	fmt.Println(şehirler)
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	for i := 0; i < 5; i++ {
		fmt.Println(şehirler[i])
	}

}
