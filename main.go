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
		log.Fatal("failed to get url:", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Fatalf("got status %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal("failed to read data:", err)
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
