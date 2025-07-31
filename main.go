package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"io/ioutil"
)

// Define a struct to match the JSON response
type Joke struct {
	ID       int    `json:"id"`
	Type     string `json:"type"`
	Setup    string `json:"setup"`
	Punchline string `json:"punchline"`
}

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

	var joke Joke
	err = json.Unmarshal(body, &joke)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	// Print only the joke setup and punchline
	fmt.Println("Here's your joke:")
	fmt.Println(joke.Setup)
	fmt.Println(joke.Punchline)
}
