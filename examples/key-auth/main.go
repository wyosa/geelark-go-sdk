package main

import (
	"context"
	"fmt"
	"log"
	"os"

	geelark "github.com/wyosa/geelark-go-sdk"
	"github.com/wyosa/geelark-go-sdk/cloudphone"
)

func main() {
	client, err := geelark.NewKeyClient(os.Getenv("GEELARK_APP_ID"), os.Getenv("GEELARK_API_KEY"))
	if err != nil {
		log.Fatal(err)
	}
	phones, err := client.CloudPhones.List(context.Background(), cloudphone.ListRequest{Page: 1, PageSize: 10})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("phones:", phones.Total)
}
