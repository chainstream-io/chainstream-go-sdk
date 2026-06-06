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
					ActivityId:      "activity-1",
					Amount:          "1",
					AmountInUsd:     "5074.5",
					AssetIds:        []string{"token-1"},
					BlockNumber:     1,
					ConditionId:     "condition-1",
					Counterparty:    "0xmaker",
					EventSlug:       "world-cup-winner",
					FeeAmountInUsd:  "1.77",
					LogIndex:        2,
					MarketIcon:      "",
					MarketId:        "market-1",
					MarketQuestion:  "question",
					Outcome:         "Yes",
					Outcomes:        []string{"Yes", "No"},
					ParticipantRole: PredictionParticipantRoleTaker,
					Price:           "0.1",
					PricePerShare:   "0.995",
					Quantity:        "10",
					RawTokenId:      "raw-token-yes",
					SeqIndex:        3,
					Source:          "chainstream",
					Taker:           "0x1",
					TakerAge:        0,
					TakerImage:      "",
					TakerName:       "",
					TakerOrderHash:  "",
					TakerPseudonym:  "",
					TakerTags:       []string{},
					Timestamp:       4,
					TokenId:         "token-1",
					TxHash:          "0x2",
					Type:            PredictionActivityTypeBuy,
					Wallet:          "0x1",
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
	if got := resp.JSON200.Activities[0].ParticipantRole; got != PredictionParticipantRoleTaker {
		t.Fatalf("unexpected participant role: %s", got)
	}
	if got := resp.JSON200.Activities[0].AmountInUsd; got != "5074.5" {
		t.Fatalf("unexpected amountInUsd: %s", got)
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
	if !PredictionParticipantRoleMaker.Valid() || !PredictionParticipantRoleTaker.Valid() || !PredictionParticipantRoleUnknown.Valid() {
		t.Fatal("maker/taker/unknown should be valid participant roles")
	}
}

func TestPredictionWalletPnlClientWithResponse(t *testing.T) {
	var observedPath string
	var observedQuery string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		observedPath = r.URL.Path
		observedQuery = r.URL.RawQuery
		w.Header().Set("Content-Type", "application/json")
		tag := "worldcup_2026"
		_ = json.NewEncoder(w).Encode(PredictionWalletPnlResponse{
			Wallet: "0xwallet",
			Tag:    &tag,
			Limit:  1,
			SortBy: PredictionWalletPnlSortByTotalPnl,
			Order:  PredictionWalletPnlOrderDesc,
			Summary: PredictionWalletPnlSummary{
				Wallet:                   "0xwallet",
				Tag:                      &tag,
				TokenCount:               1,
				MarketCount:              1,
				OpenPositionCount:        1,
				CurrentValue:             "100",
				RealizedPnl:              "12",
				UnrealizedPnl:            "3",
				TotalPnl:                 "15",
				TotalPnlRatio:            "0.15",
				TotalVolume:              "200",
				TotalBuyAmount:           "120",
				TotalSellAmount:          "80",
				TotalRedeemAmount:        "0",
				TodayRealizedPnl:         "2",
				TodayVolume:              "50",
				SevenDayRealizedPnl:      "12",
				SevenDayVolume:           "200",
				SevenDayActivityCount:    4,
				SevenDayMarketCount:      1,
				AvgInitialCost:           "120",
				AvgHoldingSeconds:        "3600",
				AvgEntryCount:            "2",
				LastActivityTs:           "2026-06-05T00:00:00Z",
				BestTradeTokenId:         "token-1",
				BestTradeConditionId:     "condition-1",
				BestTradeMarketQuestion:  "question",
				BestTradeOutcome:         "Yes",
				BestTradePnl:             "15",
				BestTradePnlShare:        "1",
				WorstTradeTokenId:        "token-1",
				WorstTradeConditionId:    "condition-1",
				WorstTradeMarketQuestion: "question",
				WorstTradeOutcome:        "Yes",
				WorstTradePnl:            "0",
				WorstTradePnlShare:       "0",
				ProfitFactor:             "2",
				SettlementRatio:          "1",
				SettlementWinRate:        "1",
				WinCount:                 1,
				LossCount:                0,
				WinRate:                  "1",
				StateQuality:             "current",
			},
			DailyPnls: []PredictionWalletDailyPnl{
				{
					Day:           "2026-06-05",
					RealizedPnl:   "2",
					Volume:        "50",
					ActivityCount: 1,
					WinCount:      1,
					LossCount:     0,
				},
			},
			Tokens: []PredictionWalletTokenPnl{
				{
					Wallet:            "0xwallet",
					TokenId:           "token-1",
					ConditionId:       "condition-1",
					EventSlug:         "world-cup-winner",
					MarketId:          "market-1",
					MarketQuestion:    "question",
					Outcome:           "Yes",
					Tags:              []string{"worldcup_2026"},
					OpenQuantity:      "10",
					CostBasis:         "120",
					AvgEntryPrice:     "12",
					AvgSellPrice:      "15",
					LastPrice:         "13",
					CurrentValue:      "130",
					RealizedPnl:       "12",
					UnrealizedPnl:     "3",
					TotalPnl:          "15",
					TotalPnlRatio:     "0.15",
					TotalBuyAmount:    "120",
					TotalSellAmount:   "80",
					TotalRedeemAmount: "0",
					TotalVolume:       "200",
					WinCount:          1,
					LossCount:         0,
					FirstActivityTs:   "2026-06-01T00:00:00Z",
					LastActivityTs:    "2026-06-05T00:00:00Z",
					LastActivityId:    "activity-1",
					LastSeqIndex:      10,
					StateQuality:      "current",
				},
			},
		})
	}))
	defer server.Close()

	client, err := NewClientWithResponses(server.URL)
	if err != nil {
		t.Fatalf("NewClientWithResponses returned error: %v", err)
	}
	limit := int64(1)
	tag := "worldcup_2026"
	sortBy := PredictionWalletPnlSortByTotalPnl
	order := PredictionWalletPnlOrderDesc
	resp, err := client.GetPredictionWalletPnlWithResponse(context.Background(), "0xwallet", &GetPredictionWalletPnlParams{
		Limit:  &limit,
		Tag:    &tag,
		SortBy: &sortBy,
		Order:  &order,
	})
	if err != nil {
		t.Fatalf("GetPredictionWalletPnlWithResponse returned error: %v", err)
	}
	if resp.JSON200 == nil {
		t.Fatal("expected JSON200 response")
	}
	if observedPath != "/v1/prediction/wallets/0xwallet/pnl" {
		t.Fatalf("unexpected path: %s", observedPath)
	}
	for _, want := range []string{"limit=1", "tag=worldcup_2026", "sort_by=totalPnl", "order=desc"} {
		if !containsQueryFragment(observedQuery, want) {
			t.Fatalf("query %q missing %q", observedQuery, want)
		}
	}
	if got := resp.JSON200.Summary.TotalPnl; got != "15" {
		t.Fatalf("unexpected total pnl: %s", got)
	}
	if got := resp.JSON200.Tokens[0].Tags[0]; got != "worldcup_2026" {
		t.Fatalf("unexpected tag: %s", got)
	}
}

func TestPredictionWalletPnlEnums(t *testing.T) {
	if !PredictionWalletPnlSortByTotalPnl.Valid() || !PredictionWalletPnlSortByLastActive.Valid() {
		t.Fatal("expected totalPnl and lastActive to be valid sort fields")
	}
	if PredictionWalletPnlSortBy("eventSlug").Valid() {
		t.Fatal("eventSlug should not be a valid PNL sort field")
	}
	if !PredictionWalletPnlOrderAsc.Valid() || !PredictionWalletPnlOrderDesc.Valid() {
		t.Fatal("asc and desc should be valid PNL orders")
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
