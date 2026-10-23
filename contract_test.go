package geelark_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"os"
	"reflect"
	"strings"
	"testing"
	"time"

	geelark "github.com/wyosa/geelark-go-sdk"
)

type auditFixture struct {
	Package, Method, Path, URL, Slug string
	Requests, Responses              []json.RawMessage
}
type auditTransport func(*http.Request) (*http.Response, error)

func (f auditTransport) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }
func auditJSON(raw []byte) any {
	var v any
	d := json.NewDecoder(bytes.NewReader(raw))
	d.UseNumber()
	if err := d.Decode(&v); err != nil {
		panic(err)
	}
	return v
}

func auditCompare(t *testing.T, expected any, actual reflect.Value, path string) {
	t.Helper()
	if expected == nil {
		return
	}
	for actual.Kind() == reflect.Ptr || actual.Kind() == reflect.Interface {
		if actual.IsNil() {
			t.Errorf("%s: unexpected nil", path)
			return
		}
		actual = actual.Elem()
	}
	if actual.Type() == reflect.TypeOf(json.RawMessage{}) {
		if !reflect.DeepEqual(expected, auditJSON(actual.Bytes())) {
			t.Errorf("%s: RawMessage differs", path)
		}
		return
	}
	switch want := expected.(type) {
	case map[string]any:
		if actual.Kind() == reflect.Map {
			for key, value := range want {
				got := actual.MapIndex(reflect.ValueOf(key))
				if !got.IsValid() {
					t.Errorf("%s.%s: missing map key", path, key)
				} else {
					auditCompare(t, value, got, path+"."+key)
				}
			}
			return
		}
		if actual.Kind() != reflect.Struct {
			t.Errorf("%s: object decoded into %s", path, actual.Type())
			return
		}
		fields := map[string]int{}
		typ := actual.Type()
		for i := 0; i < typ.NumField(); i++ {
			tag := strings.Split(typ.Field(i).Tag.Get("json"), ",")[0]
			if tag == "" {
				tag = typ.Field(i).Name
			}
			fields[tag] = i
		}
		for key, value := range want {
			index, ok := fields[key]
			if !ok {
				for k, j := range fields {
					if strings.EqualFold(k, key) {
						index = j
						ok = true
						break
					}
				}
			}
			if !ok {
				t.Errorf("%s.%s: documented field absent from %s", path, key, typ)
				continue
			}
			auditCompare(t, value, actual.Field(index), path+"."+key)
		}
	case []any:
		if actual.Kind() != reflect.Slice && actual.Kind() != reflect.Array {
			t.Errorf("%s: array decoded into %s", path, actual.Type())
			return
		}
		if len(want) != actual.Len() {
			t.Errorf("%s: array length %d != %d", path, actual.Len(), len(want))
			return
		}
		for i, value := range want {
			auditCompare(t, value, actual.Index(i), fmt.Sprintf("%s[%d]", path, i))
		}
	case json.Number:
		raw, err := json.Marshal(actual.Interface())
		if err != nil {
			t.Error(err)
			return
		}
		got := string(raw)
		if got != string(want) {
			a, ok1 := new(big.Rat).SetString(got)
			b, ok2 := new(big.Rat).SetString(string(want))
			if !ok1 || !ok2 || a.Cmp(b) != 0 {
				t.Errorf("%s: number %s became %s", path, want, got)
			}
		}
	case string:
		if actual.Kind() == reflect.String {
			if actual.String() != want {
				t.Errorf("%s: string changed", path)
			}
		} else if fmt.Sprint(actual.Interface()) != want {
			t.Errorf("%s: string %q became %v", path, want, actual.Interface())
		}
	case bool:
		if actual.Kind() != reflect.Bool || actual.Bool() != want {
			t.Errorf("%s: boolean changed", path)
		}
	}
}

func auditFixtures(t *testing.T) []auditFixture {
	t.Helper()
	raw, err := os.ReadFile("testdata/official/fixtures.json")
	if err != nil {
		t.Fatal(err)
	}
	var fixtures []auditFixture
	if err := json.Unmarshal(raw, &fixtures); err != nil {
		t.Fatal(err)
	}
	return fixtures
}

// auditSuccessEnvelope uses the documented data shape for request/path tests.
// Some pages only show a partial-success example; these tests set its envelope
// code to zero without changing the data. Response tests use the original codes.
func auditSuccessEnvelope(t *testing.T, f auditFixture) []byte {
	t.Helper()
	if len(f.Responses) == 0 {
		t.Fatalf("no documented response for %s.%s", f.Package, f.Method)
	}
	var env struct {
		Data json.RawMessage `json:"data,omitempty"`
		Code int             `json:"code"`
	}
	if err := json.Unmarshal(f.Responses[0], &env); err != nil {
		t.Fatal(err)
	}
	env.Code = 0
	raw, err := json.Marshal(env)
	if err != nil {
		t.Fatal(err)
	}
	return raw
}

func TestOfficialContracts(t *testing.T) {
	for _, f := range auditFixtures(t) {
		t.Run(f.Package+"."+f.Method, func(t *testing.T) {
			var response = []byte(`{"code":0}`)
			var sent []byte
			var active *testing.T
			status := 200
			hc := &http.Client{Transport: auditTransport(func(r *http.Request) (*http.Response, error) {
				if r.Method != "POST" || r.URL.Path != f.Path {
					active.Errorf("wrong method/path: %s %s", r.Method, r.URL.Path)
				}
				if r.URL.String() != f.URL {
					if f.Package == "proxygroup" && f.Method == "List" && r.URL.Host == "openapi.geelark.com" {
						active.Log("known source inconsistency: proxy-group list page specifies api.geelark.com")
					} else {
						active.Errorf("host/path differs from live documentation: got %s, expected %s", r.URL.String(), f.URL)
					}
				}
				if r.Header.Get("Authorization") != "Bearer audit-token" || r.Header.Get("traceId") == "" || r.Header.Get("Content-Type") != "application/json" {
					active.Error("authentication/content headers missing")
				}
				defer r.Body.Close()
				var err error
				sent, err = io.ReadAll(r.Body)
				if err != nil {
					return nil, err
				}
				return &http.Response{StatusCode: status, Header: make(http.Header), Body: io.NopCloser(bytes.NewReader(response)), Request: r}, nil
			})}
			method, err := auditMethod(f, hc)
			if err != nil {
				t.Fatal(err)
			}
			requestType := method.Type().In(1)
			responseType := method.Type().Out(0).Elem()
			call := func(body []byte) (reflect.Value, error) {
				// Isolated scenarios must not wait on an earlier Billing request.
				method, err := auditMethod(f, hc)
				if err != nil {
					return reflect.Value{}, err
				}

				request := reflect.New(requestType)
				if err := json.Unmarshal(body, request.Interface()); err != nil {
					return reflect.Value{}, err
				}
				result := method.Call([]reflect.Value{reflect.ValueOf(context.Background()), request.Elem()})
				if !result[1].IsNil() {
					return result[0], result[1].Interface().(error)
				}
				return result[0], nil
			}
			for i, body := range f.Requests {
				t.Run(fmt.Sprintf("request_%d", i), func(t *testing.T) {
					active = t
					response = auditSuccessEnvelope(t, f)
					_, err := call(body)
					if err != nil {
						t.Fatal(err)
					}
					if !reflect.DeepEqual(auditJSON(body), auditJSON(sent)) {
						corrected := auditJSON(body).(map[string]any)
						reason := ""
						if f.Package == "automation" && f.Method == "YouTubePubShort" && i == 0 {
							corrected["sameStyleUrl"] = corrected["SameStyleUrl"]
							delete(corrected, "SameStyleUrl")
							reason = "official example uses SameStyleUrl; field table specifies sameStyleUrl"
						}
						if f.Package == "automation" && f.Method == "RedditWarmup" && i == 0 {
							corrected["duration"] = json.Number("0")
							reason = "official example omits the required duration; the zero-value request includes it"
						}
						if reason != "" && reflect.DeepEqual(corrected, auditJSON(sent)) {
							t.Skip(reason)
						}
						t.Errorf("request differs from live example: expected %s; sent %s", body, sent)
					}
				})
			}
			req := []byte(`{}`)
			if len(f.Requests) > 0 {
				req = f.Requests[0]
			}
			representative := json.RawMessage(`{}`)
			for i, body := range f.Responses {
				t.Run(fmt.Sprintf("response_%d", i), func(t *testing.T) {
					active = t
					response = body
					var envelope struct {
						Code int
						Data json.RawMessage
					}
					if err := json.Unmarshal(body, &envelope); err != nil {
						t.Fatal(err)
					}
					if len(envelope.Data) > 0 && string(envelope.Data) != "null" {
						representative = envelope.Data
					}
					result, err := call(req)
					if envelope.Code != 0 {
						var apiErr *geelark.APIError
						if !errors.As(err, &apiErr) || apiErr.Code != envelope.Code {
							t.Fatalf("expected APIError %d; got %v", envelope.Code, err)
						}
						result = reflect.New(responseType)
						if len(envelope.Data) > 0 && string(envelope.Data) != "null" {
							if err := apiErr.DecodeData(result.Interface()); err != nil {
								t.Fatal(err)
							}
						}
					} else if err != nil {
						t.Fatal(err)
					}
					if len(envelope.Data) > 0 {
						auditCompare(t, auditJSON(envelope.Data), result, "data")
					}
				})
			}
			t.Run("partial_typed", func(t *testing.T) {
				active = t
				response, _ = json.Marshal(map[string]any{"code": 40006, "data": representative, "traceId": "partial-trace", "msg": "partial success"})
				_, err := call(req)
				var apiErr *geelark.APIError
				if !errors.As(err, &apiErr) || apiErr.Code != 40006 || apiErr.HTTPStatus != 200 || apiErr.TraceID != "partial-trace" {
					t.Fatalf("partial error details lost: %v", err)
				}
				result := reflect.New(responseType)
				if err := apiErr.DecodeData(result.Interface()); err != nil {
					t.Fatal(err)
				}
				auditCompare(t, auditJSON(representative), result, "partial.data")
			})
			t.Run("http_error", func(t *testing.T) {
				active = t
				status = 503
				response = []byte(`{"code":50000,"msg":"unavailable"}`)
				_, err := call(req)
				var apiErr *geelark.APIError
				if !errors.As(err, &apiErr) || apiErr.HTTPStatus != 503 || apiErr.Code != 50000 {
					t.Fatalf("HTTP failure lost: %v", err)
				}
			})
		})
	}
}

// auditMethod creates independent clients for a single fixture exchange. This
// keeps Billing's real minimum intervals enabled without delaying offline tests.
func auditMethod(f auditFixture, hc *http.Client) (reflect.Value, error) {
	cloud, err := geelark.NewClient("audit-token", geelark.WithHTTPClient(hc), geelark.WithMinInterval(time.Nanosecond))
	if err != nil {
		return reflect.Value{}, err
	}
	local, err := geelark.NewBrowserClient("audit-token", geelark.WithHTTPClient(hc), geelark.WithMinInterval(time.Nanosecond))
	if err != nil {
		return reflect.Value{}, err
	}
	services := map[string]any{
		"cloudphone": cloud.CloudPhones, "adb": cloud.ADB, "shell": cloud.Shell,
		"app": cloud.Apps, "files": cloud.Files, "library": cloud.Library,
		"automation": cloud.Automation, "analytics": cloud.Analytics,
		"livestream": cloud.LiveStream, "webhook": cloud.Webhooks, "oem": cloud.OEM,
		"proxy": cloud.Proxies, "proxygroup": cloud.ProxyGroups, "group": cloud.Groups,
		"tag": cloud.Tags, "billing": cloud.Billing,
		"browser": local.Browsers, "browserautomation": local.Automation,
	}
	service, ok := services[f.Package]
	if !ok {
		return reflect.Value{}, fmt.Errorf("unknown fixture package %s", f.Package)
	}
	method := reflect.ValueOf(service).MethodByName(f.Method)
	if !method.IsValid() {
		return reflect.Value{}, fmt.Errorf("missing SDK method %s.%s", f.Package, f.Method)
	}
	return method, nil
}
