package restful

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type todo struct {
	UserId    int    `json:"userid"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
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

//func Demo2() {
//	todo := todo{1, 2, "Alışvediş yapılacak", false}
//	jsonTodo, err := json.Marshal(todo)
//
//	response, err := http.Post("http://jsonplaceholder.typicode.com/todos", "aplication/json;charset=utf-8",
//		bytes.NewBuffer(jsonTodo))

//	if err != nil {
//		fmt.Println(err)
//	}
//	defer response.Body.Close()

//	bodyBytes, _ := ioutil.ReadAll(response.Body)

//	bodystring := string(bodyBytes)
//	fmt.Println(bodystring)

//	var todoResponse Todo
//json.Unmarshal(bodyBytes, &todoResponse)
//	fmt.Println(todoResponse)
//	}
