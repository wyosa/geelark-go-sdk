# GeeLark Go SDK

[Official API documentation](https://open.geelark.com/api)

## Requirements

- Go 1.22+ (repository toolchain: 1.27.3).
- GeeLark API token or app ID + API key.
- Running, logged-in GeeLark desktop client for browser APIs.

```sh
go get github.com/wyosa/geelark-go-sdk
```

## Examples

```go
client, err := geelark.NewClient(os.Getenv("GEELARK_TOKEN"))
if err != nil {
	log.Fatal(err)
}
phones, err := client.CloudPhones.List(context.Background(), cloudphone.ListRequest{
	Page: 1, PageSize: 10,
})
if err != nil {
	log.Fatal(err)
}
fmt.Println(phones.Items)
```

Run from the repository root:

```sh
GEELARK_TOKEN=your-token go run ./examples/basic
GEELARK_TOKEN=your-token go run ./examples/browser
GEELARK_APP_ID=your-app-id GEELARK_API_KEY=your-api-key go run ./examples/key-auth
```

[Cloud phones](examples/basic/main.go) · [Browser](examples/browser/main.go) · [Key auth](examples/key-auth/main.go) · [Uploads and errors](example_test.go) · [Polling and pagination](workflow_examples_test.go) · [Webhooks](webhook/example_test.go)

## API coverage

211/211 endpoint paths from the 2026-10-29 documentation snapshot.
[Endpoint list](testdata/endpoints.json).

| Package | Endpoints |
| --- | ---: |
| `cloudphone` | 30 |
| `automation` | 65 |
| `browser` | 18 |
| `browserautomation` | 21 |
| `app` | 19 |
| `analytics` | 9 |
| `library` | 8 |
| `billing` | 7 |
| `files` | 5 |
| `proxy` | 5 |
| `group` | 4 |
| `proxygroup` | 4 |
| `tag` | 4 |
| `livestream` | 3 |
| `patch` | 3 |
| `adb` | 2 |
| `webhook` | 2 |
| `shell` | 1 |
| `oem` | 1 |
