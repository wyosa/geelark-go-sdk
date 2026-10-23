package main

import (
	"context"
	"fmt"
	"log"
	"os"

	geelark "github.com/wyosa/geelark-go-sdk"
	"github.com/wyosa/geelark-go-sdk/browser"
)

func main() {
	client, err := geelark.NewBrowserClient(os.Getenv("GEELARK_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}
	profiles, err := client.Browsers.List(context.Background(), browser.ListRequest{Page: 1, PageSize: 10})
	if err != nil {
		log.Fatal(err)
	}
	for _, profile := range profiles.Items {
		fmt.Println(profile.ID, profile.SerialName)
	}
}
