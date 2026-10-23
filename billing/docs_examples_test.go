package billing

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
		{"BalanceResponse", "{\"balance\":1549.77,\"giftMoney\":0,\"availableTimeAddOn\":10}", func() any { return new(BalanceResponse) }},
		{"PlansResponse", "[{\"id\":\"497540679501610040\",\"price\":5,\"level\":0,\"envNum\":5,\"freeTime\":60,\"openEnvNumOneDay\":1000,\"createEnvNumOneDay\":25},{\"id\":\"512719311391949750\",\"price\":19,\"level\":1,\"envNum\":20,\"freeTime\":60,\"openEnvNumOneDay\":10000,\"createEnvNumOneDay\":200}]", func() any { return new(PlansResponse) }},
		{"ChangePlanRequest", "{\"profilesId\":\"497540679501610040\",\"parallelsNum\":1,\"monthlyRentalNum\":1}", func() any { return new(ChangePlanRequest) }},
		{"RenewPlanRequest", "{\"days\":30,\"promoCode\":\"GeeLark666\"}", func() any { return new(RenewPlanRequest) }},
		{"PlanInfoResponse", "{\"plan\":1,\"profiles\":1000,\"monthlyRental\":13,\"parallels\":3,\"expirationTime\":1774596449,\"monthlyFee\":777.4,\"availableProfiles\":160,\"availableMonthlyRentals\":0}", func() any { return new(PlanInfoResponse) }},
		{"TransactionsRequest", "{\"id\":\"612451567282427943\",\"limit\":10,\"lastFlowId\":\"612476158453291064\",\"startAt\":1774593838,\"endAt\":1774593840}", func() any { return new(TransactionsRequest) }},
		{"TransactionsResponse", "{\"total\":1,\"page\":1,\"pageSize\":10,\"list\":[{\"id\":\"612451567282427943\",\"envId\":\"612451567282427942\",\"amount\":0,\"usedTime\":2,\"type\":1,\"chargeType\":5,\"createdTime\":1774593838}],\"lastFlowId\":\"612451567282427943\"}", func() any { return new(TransactionsResponse) }},
		{"BuyTimeRequest", "{\"minutes\":2000,\"promoCode\":\"GeeLark666\"}", func() any { return new(BuyTimeRequest) }},
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
