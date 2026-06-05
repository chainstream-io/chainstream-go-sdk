package prediction

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestPredictionActivitiesClientWithResponse(t *testing.T) {
	var observedPath string
	var observedQuery string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		observedPath = r.URL.Path
		observedQuery = r.URL.RawQuery
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(PredictionActivitiesResponse{
			Activities: []PredictionActivity{
				{
					ActivityId:     "activity-1",
					Amount:         "1",
					AssetIds:       []string{"token-1"},
					BlockNumber:    1,
					ConditionId:    "condition-1",
					EventSlug:      "world-cup-winner",
					LogIndex:       2,
					MarketIcon:     "",
					MarketId:       "market-1",
					MarketQuestion: "question",
					Outcome:        "Yes",
					Outcomes:       []string{"Yes", "No"},
					Price:          "0.1",
					Quantity:       "10",
					SeqIndex:       3,
					Source:         "chainstream",
					Taker:          "0x1",
					TakerAge:       0,
					TakerImage:     "",
					TakerName:      "",
					TakerOrderHash: "",
					TakerPseudonym: "",
					TakerTags:      []string{},
					Timestamp:      4,
					TokenId:        "token-1",
					TxHash:         "0x2",
					Type:           PredictionActivityTypeBuy,
				},
			},
			EventSlug:     "world-cup-winner",
			Limit:         1,
			Order:         PredictionActivityOrderDesc,
			RetentionDays: 3,
		})
	}))
	defer server.Close()

	client, err := NewClientWithResponses(server.URL)
	if err != nil {
		t.Fatalf("NewClientWithResponses returned error: %v", err)
	}
	limit := int64(1)
	tokenID := "token-1"
	activityType := PredictionActivityTypeBuy
	order := PredictionActivityOrderDesc
	resp, err := client.GetPredictionMarketActivitiesWithResponse(context.Background(), "condition-1", &GetPredictionMarketActivitiesParams{
		Limit:        &limit,
		TokenId:      &tokenID,
		ActivityType: &activityType,
		Order:        &order,
	})
	if err != nil {
		t.Fatalf("GetPredictionMarketActivitiesWithResponse returned error: %v", err)
	}
	if resp.JSON200 == nil {
		t.Fatal("expected JSON200 response")
	}
	if observedPath != "/v1/prediction/markets/condition-1/activities" {
		t.Fatalf("unexpected path: %s", observedPath)
	}
	for _, want := range []string{"limit=1", "token_id=token-1", "activity_type=buy", "order=desc"} {
		if !containsQueryFragment(observedQuery, want) {
			t.Fatalf("query %q missing %q", observedQuery, want)
		}
	}
	if got := resp.JSON200.Activities[0].Type; got != PredictionActivityTypeBuy {
		t.Fatalf("unexpected activity type: %s", got)
	}
}

func TestPredictionActivityEnums(t *testing.T) {
	if !PredictionActivityTypeInventoryAdjust.Valid() {
		t.Fatal("inventory_adjust should be a valid activity type")
	}
	if PredictionActivityType("unknown").Valid() {
		t.Fatal("unknown activity type should be invalid")
	}
	if !PredictionActivityOrderAsc.Valid() || !PredictionActivityOrderDesc.Valid() {
		t.Fatal("asc and desc should be valid orders")
	}
}

func containsQueryFragment(query string, fragment string) bool {
	for _, part := range strings.Split(query, "&") {
		if part == fragment {
			return true
		}
	}
	return false
}
