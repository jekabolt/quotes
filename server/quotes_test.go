package server

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

func TestHowMuchIsXGreaterThanYt(t *testing.T) {
	basket := ResponseGetQuote{
		Quotes: []Quote{
			Quote{Id: "kek",
				Category: []string{"cat1", "cat2"}},
		},
	}
	jsonData, err := json.Marshal(basket)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("%s", jsonData)
}
