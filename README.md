# GeeLark Go SDK

[Русский](README.ru.md)

An unofficial Go SDK for the [GeeLark API](https://open.geelark.com/api). It follows the package layout and explicit request/response style of [duoplus-go-sdk](https://github.com/wyosa/duoplus-go-sdk).

The SDK covers 206 documented endpoint paths across cloud phone, app, automation, file, library, proxy, group, tag, analytics, billing, webhook, OEM, live stream, and local browser APIs. See the [endpoint map](docs/coverage.md). The documentation snapshot used for this release is 2026-10-23, including the [latest changelog](https://open.geelark.com/api/change-log).

## Requirements

- Go 1.22+
- A GeeLark API token, or an app ID and API key for the cloud APIs
- The GeeLark desktop client, open and logged in, for local browser APIs

The module path is `github.com/wyosa/geelark-go-sdk`. Install the SDK in your project:

```bash
go get github.com/wyosa/geelark-go-sdk@latest
```

Go 1.22 is the minimum language version. Use a supported, patched Go release for production. The 2026-10-23 audit found [GO-2026-6809](https://pkg.go.dev/vuln/GO-2026-6809) in the installed Go 1.27.2 standard library; the scan passed with Go 1.27.3.

## Quick start

```go
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
	phones, err := client.CloudPhones.List(context.Background(), cloudphone.ListRequest{
		Page: 1, PageSize: 10,
	})
	if err != nil {
		log.Fatal(err)
	}
	for _, phone := range phones.Items {
		fmt.Println(phone.ID, phone.SerialName, phone.Status)
	}
}
```

Run the examples in this repository:

```bash
GEELARK_TOKEN=your-token go run ./examples/basic
GEELARK_TOKEN=your-token go run ./examples/browser
GEELARK_APP_ID=your-app-id GEELARK_API_KEY=your-api-key go run ./examples/key-auth
```

## Clients and authentication

```go
cloud, err := geelark.NewClient(token) // Bearer token
cloud, err := geelark.NewKeyClient(appID, apiKey) // signed headers
browser, err := geelark.NewBrowserClient(token) // local desktop API
```

Cloud requests use `https://openapi.geelark.com/open/v1` (and `/open/v2` where documented). Browser requests use `http://localhost:40185/api/v1`. The browser extension category method is a remote cloud endpoint and uses the browser client's token.

Each request uses `POST` with JSON and a fresh UUIDv4 `traceId`. Key authentication signs `appId + traceId + ts + nonce + apiKey` with uppercase SHA-256. API success is `code: 0`; other codes return `*geelark.APIError`. The error includes the HTTP status, GeeLark code, message, trace ID, and any partial-success `data`.

The API limit is 200 requests per minute and 24,000 per hour **per interface**. Each SDK client spaces calls to the same path by 310 ms by default; separate client instances do not share a limiter. It never retries a request automatically. Configure the client with:

- `geelark.WithBaseURL(url)` — cloud or local browser host for the constructor being called
- `geelark.WithRemoteBaseURL(url)` — remote host for `NewBrowserClient`'s extension category method
- `geelark.WithHTTPClient(client)` — custom HTTP client, copied before redirects are disabled
- `geelark.WithMinInterval(duration)` — per-path spacing

The API host must use HTTPS, except loopback hosts used for local browser access and tests. Redirects are refused so credentials are never forwarded to another host.

## Services

Cloud client services: `CloudPhones`, `ADB`, `Shell`, `Apps`, `Files`, `Library`, `Automation`, `Analytics`, `LiveStream`, `Webhooks`, `OEM`, `Proxies`, `ProxyGroups`, `Groups`, `Tags`, and `Billing`.

Browser client services: `Browsers` and `Automation`. Request and response types live in matching packages (`cloudphone`, `browser`, `browserautomation`, etc.); shared helpers live in `common`.

Optional numeric and boolean fields that need to carry a zero value are pointers. For example:

```go
_, err := cloud.CloudPhones.Start(ctx, cloudphone.StartRequest{
	IDs: []string{"phone-id"},
	EnableEnergySavingModeAfterMinutes: common.Ptr(0),
	DarkMode: common.Ptr(false),
})
```

Some task payloads vary by task type. `automation.AddTaskItem` contains the union of documented fields; set only those relevant to the chosen `TaskType`. A few response bodies without a stable documented shape remain `json.RawMessage`.

Custom cloud and browser flows accept arbitrary input names in `ParamMap`, including boolean and numeric values:

```go
req := automation.CreateCustomTaskRequest{
	ID: "phone-id", FlowID: "flow-id", ScheduleAt: int(time.Now().Unix()),
	ParamMap: automation.CreateCustomTaskRequestParamMap{
		"Caption": "Example", "Count": 0, "Enabled": false,
		"Files": []string{"https://example.com/video.mp4"},
	},
}
```

In update requests, cloud phone, browser, and analytics tag lists and the cloud phone network blacklist distinguish `nil` (omit the field) from `[]string{}` (send an explicit empty array). For example, `browser.UpdateRequest{ID: "browser-id", TagIDs: []string{}}` clears browser tags.

## Errors and partial success

```go
var apiErr *geelark.APIError
if errors.As(err, &apiErr) {
	fmt.Println(apiErr.Code, apiErr.Message, apiErr.TraceID)
	if apiErr.Code == 40006 { // partial success
		var partial cloudphone.StartResponse
		if decodeErr := apiErr.DecodeData(&partial); decodeErr == nil {
			fmt.Println(partial.SuccessAmount, partial.FailAmount)
		}
	}
}
```

## Uploads and callbacks

`Files.Upload` / `Files.UploadFile` and `Library.Upload` / `Library.UploadFile` fetch a signed URL and PUT file bytes to storage without GeeLark credentials. GeeLark's examples currently return HTTP signed storage URLs, which are accepted for this separate upload step. Call `Library.Create` with the returned resource URL to retain a Library file.

API requests have a default 30-second HTTP timeout. Storage uploads use the caller's context without a default HTTP timeout; set a context deadline for them. A custom `WithHTTPClient` timeout is honored by both API requests and storage uploads.

The `webhook` package defines callback event types and a common `Event` struct. `webhook.VerifyKeySignature` and `webhook.VerifyToken` can check callback headers before processing the body.

## Development

```bash
go test -race -coverpkg=./... ./...
go vet ./...
```

Tests verify every endpoint path and method, authorization, signed uploads, errors, and 356 request/response examples from the official documentation. Two inconsistent documentation request examples are excluded from example round-trip tests: one uses `SameStyleUrl` instead of the listed `sameStyleUrl`, and one omits a documented required `duration`.

Every endpoint is exercised against an HTTP test server for success, partial success, and HTTP failure. Regression tests cover fractional billing amounts, custom flow parameters, empty update lists, task failure details, cancellation, rate limiting, and upload timeouts. See the [audit findings and compatibility notes](docs/coverage.md#audit-2026-10-23). Authenticated cloud operations and the GeeLark desktop client were not exercised during this audit because API credentials were unavailable.

The [official Markdown bundle](https://material.geelark.com/openapi-doc/en-doc.zip) and [official GitHub repository](https://github.com/GeeLark/geelark-openapi) help with bulk review. The live API pages take precedence when they differ from the bundle or the partial OpenAPI specification.
