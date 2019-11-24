package main

import (
	"fmt"
	"net/http"
)


// CheckWebsite returns true if the URL returns a 200 status code, false otherwise
func CheckWebsite(url string) bool {
	response, err := http.Head(url)
	if err != nil {
		return false
	}

	if response.StatusCode != http.StatusOK {
		return false
	}

	return true
}

func main(){
	testUrl := "http://www.baidu.com"
	if CheckWebsite(testUrl){
		fmt.Println("Now you have the access to baidu")
	}else {
		fmt.Println("Now you do not have the access to baidu")
	}
}
