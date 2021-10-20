package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	url := "https://sameer-kumar-aztro-v1.p.rapidapi.com/?sign=aries&day=tomorrow"

	req, _ := http.NewRequest("POST", url, nil)

	req.Header.Add("x-rapidapi-host", "sameer-kumar-aztro-v1.p.rapidapi.com")
	req.Header.Add("x-rapidapi-key", "d78c3f1461msh9d8b3f5fb7d0e5cp1612abjsn9032a3e600c8")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))

}