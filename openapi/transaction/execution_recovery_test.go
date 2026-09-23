package transaction

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestExecutionRecoveryGET(t *testing.T) {
	raw, err := os.ReadFile("../../tests/fixtures/execution-recovery.json")
	if err != nil {
		t.Fatal(err)
	}
	var fixture struct {
		Body json.RawMessage `json:"body"`
	}
	if err := json.Unmarshal(raw, &fixture); err != nil {
		t.Fatal(err)
	}
	for _, status := range []int{200, 404} {
		requests := 0
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requests++
			if r.Method != "GET" || r.URL.Path != "/v2/transaction/9007199254740993/orders/by-idempotency" {
				t.Errorf("unexpected request %s %s", r.Method, r.URL.Path)
			}
			if r.URL.Query().Get("idempotencyKey") != "recovery:/?&=+%encoded" {
				t.Error("key altered")
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(status)
			_, _ = w.Write(fixture.Body)
		}))
		client, err := NewClientWithResponses(server.URL)
		if err != nil {
			t.Fatal(err)
		}
		result, err := client.ExecutionFindOrderWithResponse(context.Background(), "9007199254740993", &ExecutionFindOrderParams{IdempotencyKey: "recovery:/?&=+%encoded"})
		server.Close()
		if err != nil {
			t.Fatal(err)
		}
		if requests != 1 || result.StatusCode() != status {
			t.Fatal("unexpected request count/status")
		}
		if status == 404 {
			if result.JSON200 != nil {
				t.Fatal("404 became order")
			}
			continue
		}
		encoded, err := json.Marshal(result.JSON200)
		if err != nil {
			t.Fatal(err)
		}
		var got, want interface{}
		_ = json.Unmarshal(encoded, &got)
		_ = json.Unmarshal(fixture.Body, &want)
		if stringMustJSON(t, got) != stringMustJSON(t, want) {
			t.Fatal("recovered response changed")
		}
	}
}
func stringMustJSON(t *testing.T, value interface{}) string {
	t.Helper()
	b, err := json.Marshal(value)
	if err != nil {
		t.Fatal(err)
	}
	return string(b)
}
