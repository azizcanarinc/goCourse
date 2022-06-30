package restful

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type todo struct {
	UserId    int  `json:"userid"`
	Id        int  `json:"id"`
	Title     int  `json:"title"`
	Completed bool `json:"completed"`
}

func Demo1() {
	response, err := http.Get("http://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(response.Body)
	bodystring := string(bodyBytes)
	fmt.Println(bodystring)

	var todo todo
	json.Unmarshal(bodyBytes, &todo)
	fmt.Println(todo)
}
func Demo2() {
	todo := todo{1, 2, "alışvediş yapılacak", folse}
}
