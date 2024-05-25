package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	url := "https://ai-weather-by-meteosource.p.rapidapi.com/find_places?text=fishermans%20wharf&language=en"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", "")
	req.Header.Add("X-RapidAPI-Host", "ai-weather-by-meteosource.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
