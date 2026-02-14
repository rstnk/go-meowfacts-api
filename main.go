package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const Endpoint = "https://meowfacts.herokuapp.com/"

type CatFact struct {
	Fact []string `json:"data"`
}

func main() {
	res, err := http.Get(Endpoint)
	if err != nil {
		log.Fatal("failed to get url:", Endpoint)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal("failed to read data:", err)
	}

	if res.StatusCode != 200 {
		log.Fatalf("got status %d", res.StatusCode)
	}

	var cf CatFact
	if err = json.Unmarshal(body, &cf); err != nil {
		log.Fatal("could not parse data:", err)
	}

	if len(cf.Fact) == 0 {
		log.Fatal("no cat facts!")
	}
	fmt.Println(cf.Fact[0])

}
