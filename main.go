package main

import (
	"fmt"

	"github.com/Mitmadhu/commons/handle_http"
)

func main(){
	resp, err := handle_http.HandleHttpCall("https://jsonplaceholder.typicode.com/todos", "GET", nil)
	if(err != nil){
		print(err.Error())
		return
	}
	fmt.Printf("%v", resp)
}