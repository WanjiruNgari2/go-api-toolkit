package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Joke struct {
	Setup     string `json:"setup"`
	Punchline string `json:"punchline"`
}

type ChuckNorrisJoke struct {
	Value string `json:"value"`
}

func main() {
	fmt.Println("Choose joke type: 1 for Dad Joke, 2 for Chuck Norris Joke")
	var choice int
	fmt.Scanln(&choice)

	var url string
	if choice == 2 {
		url = "https://api.chucknorris.io/jokes/random"
	} else {
		url = "https://official-joke-api.appspot.com/random_joke"
	}

	resp, err := http.Get(url)
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

	fmt.Println("Time fetched:", time.Now().Format(time.RFC1123))

	if choice == 2 {
		var chuck ChuckNorrisJoke
		err = json.Unmarshal(body, &chuck)
		if err != nil {
			fmt.Println("Error parsing JSON. Raw body:")
			fmt.Println(string(body))
			fmt.Println("Error details:", err)
			return
		}
		fmt.Println("Chuck Norris says:")
		fmt.Println(chuck.Value)
	} else {
		var joke Joke
		err = json.Unmarshal(body, &joke)
		if err != nil {
			fmt.Println("Error parsing JSON. Raw body:")
			fmt.Println(string(body))
			fmt.Println("Error details:", err)
			return
		}
		fmt.Println("Here's your joke:")
		fmt.Println(joke.Setup)
		fmt.Println(joke.Punchline)
	}
}


		

