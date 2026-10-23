package tag

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestDocumentationExamples(t *testing.T) {
	tests := []struct {
		name     string
		raw      string
		newValue func() any
	}{
		{"CreateRequest", "{\"list\":[{\"name\":\"tagEmpty\"},{\"name\":\"tagRed\",\"color\":\"red\"},{\"name\":\"tagBlue\",\"color\":\"blue\"},{\"name\":\"tagGreen\",\"color\":\"green\"},{\"name\":\"tagYellow\",\"color\":\"yellow\"},{\"name\":\"tagPurple\",\"color\":\"purple\"},{\"name\":\"tagWhite\",\"color\":\"white\"},{\"name\":\"tagWhite2\",\"color\":\"\"},{\"name\":\"tagInvalid\",\"color\":\"abc\"}]}", func() any { return new(CreateRequest) }},
		{"CreateResponse", "{\"totalAmount\":9,\"successAmount\":8,\"failAmount\":1,\"successDetails\":[{\"id\":\"528989565877224448\",\"name\":\"tagWhite2\",\"color\":\"white\"},{\"id\":\"528989565877289984\",\"name\":\"tagBlue\",\"color\":\"blue\"},{\"id\":\"528989565877355520\",\"name\":\"tagGreen\",\"color\":\"green\"},{\"id\":\"528989565894001664\",\"name\":\"tagYellow\",\"color\":\"yellow\"},{\"id\":\"528989565894067200\",\"name\":\"tagPurple\",\"color\":\"purple\"},{\"id\":\"528989565894132736\",\"name\":\"tagWhite\",\"color\":\"white\"},{\"id\":\"528989565894198272\",\"name\":\"tagEmpty\",\"color\":\"white\"},{\"id\":\"528989565910778880\",\"name\":\"tagRed\",\"color\":\"red\"}],\"failDetails\":[{\"code\":43023,\"name\":\"tagInvalid\",\"msg\":\"tag color not found\"}]}", func() any { return new(CreateResponse) }},
		{"DeleteRequest", "{\"ids\":[\"528989565877224448\",\"528953724308030464\"]}", func() any { return new(DeleteRequest) }},
		{"DeleteResponse", "{\"totalAmount\":2,\"successAmount\":1,\"failAmount\":1,\"failDetails\":[{\"code\":43022,\"id\":\"528953724308030464\",\"msg\":\"tag not found\"}]}", func() any { return new(DeleteResponse) }},
		{"ListRequest", "{\"page\":1,\"pageSize\":5}", func() any { return new(ListRequest) }},
		{"ListResponse", "{\"total\":24,\"page\":1,\"pageSize\":5,\"list\":[{\"id\":\"528989565877355520\",\"name\":\"tagGreen\",\"color\":\"green\"},{\"id\":\"528989565877289984\",\"name\":\"tagBlue\",\"color\":\"blue\"},{\"id\":\"528989565894067200\",\"name\":\"tagPurple\",\"color\":\"purple\"},{\"id\":\"528989565910778880\",\"name\":\"tagRed\",\"color\":\"red\"},{\"id\":\"528989565894198272\",\"name\":\"tagEmpty\",\"color\":\"white\"}]}", func() any { return new(ListResponse) }},
		{"UpdateRequest", "{\"list\":[{\"id\":\"528989565894198272\",\"name\":\"tagEmptyUpdate\",\"color\":\"red\"},{\"id\":\"528994200482481152\",\"name\":\"tagUpdate\",\"color\":\"red\"}]}", func() any { return new(UpdateRequest) }},
		{"UpdateResponse", "{\"totalAmount\":2,\"successAmount\":2,\"failAmount\":0}", func() any { return new(UpdateResponse) }},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			value := tc.newValue()
			if err := json.Unmarshal([]byte(tc.raw), value); err != nil {
				t.Fatal(err)
			}
			got, err := json.Marshal(value)
			if err != nil {
				t.Fatal(err)
			}
			var a, b any
			if err := json.Unmarshal([]byte(tc.raw), &a); err != nil {
				t.Fatal(err)
			}
			if err := json.Unmarshal(got, &b); err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(a, b) {
				t.Errorf("roundtrip mismatch: want %s, got %s", tc.raw, got)
			}
		})
	}
}
