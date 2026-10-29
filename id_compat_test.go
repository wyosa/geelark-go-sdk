package geelark_test

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	geelark "github.com/wyosa/geelark-go-sdk"
	"github.com/wyosa/geelark-go-sdk/app"
	"github.com/wyosa/geelark-go-sdk/cloudphone"
	"github.com/wyosa/geelark-go-sdk/library"
)

func TestResponseIdentifierCompatibility(t *testing.T) {
	types := []any{
		cloudphone.StopResponseFailDetailsItem{},
		cloudphone.StopResponseSuccessDetailsItem{},
		cloudphone.DeleteResponseFailDetailsItem{},
		cloudphone.UpdateResponseFailDetailsItem{},
		cloudphone.HideAccessibilityResponseFailDetailsItem{},
		app.BatchOperateResponseItemsItem{},
		library.SearchResponseListItem{},
		library.SearchResponseListItemTagsItem{},
		library.SearchTagsResponseListItem{},
	}
	for _, sample := range types {
		t.Run(reflect.TypeOf(sample).String(), func(t *testing.T) {
			for _, tc := range []struct{ raw, want string }{
				{`"phone-id"`, "phone-id"},
				{`""`, ""},
				{`0`, "0"},
				{`9007199254740993`, "9007199254740993"},
				{`"9007199254740993"`, "9007199254740993"},
				{`18446744073709551616`, "18446744073709551616"},
			} {
				value := reflect.New(reflect.TypeOf(sample))
				if err := json.Unmarshal([]byte(`{"id":`+tc.raw+`}`), value.Interface()); err != nil {
					t.Fatal(err)
				}
				if got := value.Elem().FieldByName("ID").String(); got != tc.want {
					t.Fatalf("ID %s became %q, want %q", tc.raw, got, tc.want)
				}
				encoded, err := json.Marshal(value.Interface())
				if err != nil {
					t.Fatal(err)
				}
				var fields map[string]json.RawMessage
				if err := json.Unmarshal(encoded, &fields); err != nil {
					t.Fatal(err)
				}
				var id string
				if err := json.Unmarshal(fields["id"], &id); err != nil || id != tc.want {
					t.Fatalf("public string encoding changed: %s (%v)", encoded, err)
				}
				// Preserve encoding/json's existing string semantics on reused values.
				for _, raw := range []string{`{}`, `{"id":null}`} {
					if err := json.Unmarshal([]byte(raw), value.Interface()); err != nil || value.Elem().FieldByName("ID").String() != tc.want {
						t.Fatalf("%s changed an existing ID: %v", raw, err)
					}
				}
			}
			for _, raw := range []string{`1.25`, `1e3`, `true`, `[]`, `{}`} {
				value := reflect.New(reflect.TypeOf(sample))
				if err := json.Unmarshal([]byte(`{"id":`+raw+`}`), value.Interface()); err == nil {
					t.Errorf("accepted invalid identifier %s", raw)
				}
			}
		})
	}
}

func TestNumericIdentifierInPartialResponse(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"code":40006,"data":{"failDetails":[{"id":9007199254740993,"code":42001,"msg":"missing"}]}}`)
	}))
	defer srv.Close()
	c, err := geelark.NewClient("token", geelark.WithBaseURL(srv.URL))
	if err != nil {
		t.Fatal(err)
	}
	_, err = c.CloudPhones.Stop(context.Background(), cloudphone.StopRequest{IDs: []string{"9007199254740993"}})
	var apiErr *geelark.APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("expected partial API error, got %v", err)
	}
	var partial cloudphone.StopResponse
	if err := apiErr.DecodeData(&partial); err != nil {
		t.Fatal(err)
	}
	if len(partial.FailDetails) != 1 || partial.FailDetails[0].ID != "9007199254740993" || partial.FailDetails[0].Code != 42001 || partial.FailDetails[0].Msg != "missing" {
		t.Fatalf("partial response lost fields: %+v", partial)
	}
}
