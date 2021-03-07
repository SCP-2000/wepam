package main

import (
	"context"
	"github.com/SCP-2000/wepam/pkg/adapter"
	"log"
)

func main() {
	config := &adapter.Config{
		// the client id is left intentionally as it is not a secret and can be used for testing
		ClientID: "111109bc0316c05e23aa",
		Adapter:  &adapter.Github{},
	}
	user, err := config.Auth(context.Background(), func(url string, code string) error {
		log.Printf("please navigate to %s, and input %s\n", url, code)
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(user)
}
