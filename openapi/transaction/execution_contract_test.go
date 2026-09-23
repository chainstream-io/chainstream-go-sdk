package transaction

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestExecutionOrderHTTPContract(t *testing.T) {
	const large = "115792089237316195423570985008687907853269984665640564039457584007913129639935"
	const prefix = `{"orderId":"11111111-1111-4111-8111-111111111111","quoteId":"22222222-2222-4222-8222-222222222222","chainId":"9007199254740993","kind":"crossChain","state":"source_confirmed","updatedAt":1,`
	for _, tc := range []struct {
		name, amounts string
		valid         bool
	}{
		{"null", `"actualBuyAmount":null,"actualRefundAmount":null}`, true},
		{"large", `"actualBuyAmount":"` + large + `","actualRefundAmount":"0"}`, true},
		{"missing_buy", `"actualRefundAmount":null}`, false},
		{"missing_refund", `"actualBuyAmount":null}`, false},
		{"number", `"actualBuyAmount":1,"actualRefundAmount":null}`, false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if !strings.Contains(r.URL.Path, "9007199254740993") {
					t.Errorf("chain precision lost: %s", r.URL.Path)
				}
				w.Header().Set("Content-Type", "application/json")
				_, _ = w.Write([]byte(prefix + tc.amounts))
			}))
			defer server.Close()
			client, err := NewClientWithResponses(server.URL)
			if err != nil {
				t.Fatal(err)
			}
			response, err := client.ExecutionGetOrderWithResponse(context.Background(), "9007199254740993", uuid.MustParse("11111111-1111-4111-8111-111111111111"))
			if !tc.valid {
				if err == nil {
					t.Fatal("accepted missing or non-string actual amount")
				}
				return
			}
			if err != nil {
				t.Fatal(err)
			}
			encoded, err := json.Marshal(response.JSON200)
			if err != nil {
				t.Fatal(err)
			}
			var fields map[string]json.RawMessage
			_ = json.Unmarshal(encoded, &fields)
			for _, key := range []string{"actualBuyAmount", "actualRefundAmount"} {
				if _, ok := fields[key]; !ok {
					t.Fatalf("omitted %s", key)
				}
			}
			if tc.name == "large" && *response.JSON200.ActualBuyAmount != large {
				t.Fatal("amount precision lost")
			}
			if response.JSON200.State != "source_confirmed" {
				t.Fatal("source confirmation changed")
			}
		})
	}
}
