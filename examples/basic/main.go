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
	client, err := geelark.NewClient(os.Getenv("GEELARK_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}
	phones, err := client.CloudPhones.List(context.Background(), cloudphone.ListRequest{Page: 1, PageSize: 10})
	if err != nil {
		log.Fatal(err)
	}
	for _, phone := range phones.Items {
		fmt.Println(phone.ID, phone.SerialName, phone.Status)
	}
}
