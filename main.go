package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("https://official-joke-api.appspot.com/random_joke")
	if err != nil {
		fmt.Println("Error fetching joke:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	fmt.Println("Joke data:")
	fmt.Println(string(body))
}
