package chainstream

import (
	"bytes"
	"encoding/json"
	"github.com/chainstream-io/chainstream-go-sdk/v2/openapi/bridgeaggregator"
	"github.com/chainstream-io/chainstream-go-sdk/v2/openapi/dex"
	"github.com/chainstream-io/chainstream-go-sdk/v2/openapi/job"
	"github.com/chainstream-io/chainstream-go-sdk/v2/openapi/swapaggregator"
	"github.com/chainstream-io/chainstream-go-sdk/v2/openapi/transaction"
	"io"
	"net/http"
	"os"
	"reflect"
	"strings"
	"testing"
)

func parseSample(route string, status int, body []byte) (any, error) {
	response := &http.Response{StatusCode: status, Header: http.Header{"Content-Type": []string{"application/json"}}, Body: io.NopCloser(bytes.NewReader(body))}
	switch {
	case strings.HasPrefix(route, "SSE "):
		response.Header.Set("Content-Type", "text/event-stream")
		return job.ParseStreamingResponse(response)
	case strings.Contains(route, "/orders/"):
		return transaction.ParseExecutionGetOrderResponse(response)
	case strings.HasPrefix(route, "GET /v2/job/"):
		return job.ParseGetResponse(response)
	case strings.Contains(route, "/gasless/submit"):
		return transaction.ParseExecutionGaslessSubmitResponse(response)
	case strings.Contains(route, "/gasless/quote"):
		return dex.ParseExecutionGaslessRequestQuoteResponse(response)
	case strings.Contains(route, "/execution/build"):
		return dex.ParseExecutionBuildQuoteResponse(response)
	case strings.Contains(route, "/dex/aggregator/quote"):
		return swapaggregator.ParseExecutionAggregatorQuoteResponse(response)
	case strings.Contains(route, "/dex/aggregator/swap"):
		return swapaggregator.ParseExecutionAggregatorBuildResponse(response)
	case strings.Contains(route, "/bridge/aggregator/route"):
		return bridgeaggregator.ParseExecutionBridgeRouteResponse(response)
	case strings.Contains(route, "/bridge/aggregator/quote"):
		return bridgeaggregator.ParseExecutionBridgeQuoteResponse(response)
	case strings.HasSuffix(route, "/send"):
		return transaction.ParseSendResponse(response)
	case strings.HasSuffix(route, "/route"):
		return dex.ParseRouteResponse(response)
	default:
		return dex.ParseQuoteResponse(response)
	}
}
func TestExecutionActualHTTPSamples(t *testing.T) {
	data, err := os.ReadFile("tests/fixtures/execution-http-samples.json")
	if err != nil {
		t.Fatal(err)
	}
	var samples map[string]struct {
		Status int             `json:"status"`
		Body   json.RawMessage `json:"body"`
	}
	if err = json.Unmarshal(data, &samples); err != nil {
		t.Fatal(err)
	}
	if len(samples) != 20 {
		t.Fatalf("expected 20 HTTP/SSE samples, got %d", len(samples))
	}
	for route, sample := range samples {
		t.Run(route, func(t *testing.T) {
			body := []byte(sample.Body)
			if strings.HasPrefix(route, "SSE ") {
				var stream string
				_ = json.Unmarshal(body, &stream)
				body = []byte(stream)
			}
			parsed, err := parseSample(route, sample.Status, body)
			if err != nil {
				t.Fatal(err)
			}
			result := reflect.ValueOf(parsed).Elem()
			if strings.HasPrefix(route, "SSE ") {
				if !bytes.Equal(result.FieldByName("Body").Bytes(), body) {
					t.Fatal("SSE data changed")
				}
				return
			}
			field := result.FieldByName("JSON200")
			if sample.Status == 201 {
				field = result.FieldByName("JSON201")
			}
			if !field.IsValid() || field.IsNil() {
				t.Fatal("missing typed HTTP response")
			}
			encoded, err := json.Marshal(field.Interface())
			if err != nil {
				t.Fatal(err)
			}
			var actual, expected map[string]any
			_ = json.Unmarshal(encoded, &actual)
			_ = json.Unmarshal(body, &expected)
			for _, key := range []string{"execution", "platformFeePolicy", "routeInfo", "actualBuyAmount", "actualRefundAmount", "status", "success"} {
				if value, exists := expected[key]; exists {
					if !reflect.DeepEqual(actual[key], value) {
						t.Fatalf("lost or changed %s", key)
					}
				}
			}
			if execution, ok := expected["execution"].(map[string]any); ok && execution["orderId"] != nil {
				for _, mutation := range []int{0, 1, 2} {
					var bad map[string]any
					_ = json.Unmarshal(body, &bad)
					order := bad["execution"].(map[string]any)
					if mutation == 0 {
						delete(order, "actualBuyAmount")
					}
					if mutation == 1 {
						order["actualBuyAmount"] = 1
					}
					if mutation == 2 {
						bad["execution"] = nil
					}
					invalid, _ := json.Marshal(bad)
					if _, err := parseSample(route, sample.Status, invalid); err == nil {
						t.Fatal("invalid execution fell back to legacy")
					}
				}
			}
		})
	}
}

func TestExecutionLegacyJobResultJSON(t *testing.T) {
	for _, raw := range []string{"null", "0", "false", `""`, `[]`, `{}`, `"text"`, `[1,null]`, `{"a":0}`} {
		data := []byte(`{"id":"legacy","status":"completed","result":` + raw + `}`)
		var parsed job.JobStatusResponse
		if err := json.Unmarshal(data, &parsed); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(parsed)
		if err != nil {
			t.Fatal(err)
		}
		var fields map[string]json.RawMessage
		_ = json.Unmarshal(encoded, &fields)
		if !bytes.Equal(fields["result"], []byte(raw)) {
			t.Fatalf("result changed: %s => %s", raw, encoded)
		}
	}
}
