package geelark_test

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	geelark "github.com/wyosa/geelark-go-sdk"
	"github.com/wyosa/geelark-go-sdk/cloudphone"
	"github.com/wyosa/geelark-go-sdk/files"
)

func ExampleNewClient() {
	// Read the token from your application's secret configuration.
	client, err := geelark.NewClient("example-token")
	if err != nil {
		panic(err)
	}
	fmt.Println(client.CloudPhones != nil)
	// Output: true
}

func ExampleNewKeyClient() {
	client, err := geelark.NewKeyClient("example-app-id", "example-api-key")
	if err != nil {
		panic(err)
	}
	fmt.Println(client.CloudPhones != nil)
	// Output: true
}

// exampleTransport keeps the executable examples offline. Real applications
// normally use the default HTTP client or their own shared transport.
type exampleTransport func(*http.Request) (*http.Response, error)

func (f exampleTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	if r.Body != nil {
		defer r.Body.Close()
	}
	return f(r)
}

func exampleReply(body string) *http.Response {
	return &http.Response{StatusCode: http.StatusOK, Header: make(http.Header), Body: io.NopCloser(strings.NewReader(body))}
}

func ExampleClient() {
	httpClient := &http.Client{Transport: exampleTransport(func(r *http.Request) (*http.Response, error) {
		return exampleReply(`{"code":0,"data":{"total":1,"page":1,"pageSize":10,"items":[{"id":"phone-id","serialName":"Demo"}]}}`), nil
	})}
	client, err := geelark.NewClient("example-token", geelark.WithHTTPClient(httpClient))
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := client.CloudPhones.List(ctx, cloudphone.ListRequest{Page: 1, PageSize: 10})
	if err != nil {
		panic(err)
	}
	for _, phone := range result.Items {
		fmt.Println(phone.ID, phone.SerialName)
	}
	// Output: phone-id Demo
}

func ExampleAPIError() {
	httpClient := &http.Client{Transport: exampleTransport(func(r *http.Request) (*http.Response, error) {
		return exampleReply(`{"code":40006,"msg":"partial success","data":{"successAmount":1,"failAmount":1,"failDetails":[{"id":9007199254740993,"code":42001,"msg":"missing"}]}}`), nil
	})}
	client, err := geelark.NewClient("example-token", geelark.WithHTTPClient(httpClient))
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err = client.CloudPhones.Stop(ctx, cloudphone.StopRequest{IDs: []string{"phone-id", "9007199254740993"}})
	var apiErr *geelark.APIError
	if !errors.As(err, &apiErr) || apiErr.Code != 40006 {
		panic(fmt.Sprintf("expected partial success, got %v", err))
	}
	var partial cloudphone.StopResponse
	if err := apiErr.DecodeData(&partial); err != nil {
		panic(err)
	}
	fmt.Println(partial.SuccessAmount, partial.FailAmount, partial.FailDetails[0].ID)
	// Output: 1 1 9007199254740993
}

func ExampleClient_upload() {
	httpClient := &http.Client{Transport: exampleTransport(func(r *http.Request) (*http.Response, error) {
		if r.Method == http.MethodPost {
			return exampleReply(`{"code":0,"data":{"uploadUrl":"https://storage.example/upload","resourceUrl":"https://storage.example/file.xml"}}`), nil
		}
		body, err := io.ReadAll(r.Body)
		if err != nil {
			return nil, err
		}
		fmt.Println("storage bytes:", len(body), "API authorization:", r.Header.Get("Authorization") != "")
		return exampleReply(""), nil
	})}
	client, err := geelark.NewClient("example-token", geelark.WithHTTPClient(httpClient))
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	content := "<root/>"
	result, err := client.Files.Upload(ctx, files.GetUploadURLRequest{FileType: "xml"}, strings.NewReader(content), int64(len(content)))
	if err != nil {
		panic(err)
	}
	fmt.Println(result.ResourceURL)
	// Output:
	// storage bytes: 7 API authorization: false
	// https://storage.example/file.xml
}
