package stream

import (
	"encoding/json"
	"testing"
)

func predictionActivityFrame(t *testing.T) map[string]interface{} {
	t.Helper()

	var frame map[string]interface{}
	err := json.Unmarshal([]byte(`{
		"seq": 42,
		"a": {
			"id": "activity-1",
			"amt": "1.23",
			"as": ["asset-yes", "asset-no"],
			"bn": 123,
			"cid": "0xcondition",
			"es": "world-cup-winner",
			"li": 7,
			"mi": "",
			"mid": "558936",
			"mq": "Will France win the 2026 FIFA World Cup?",
			"oc": "Yes",
			"ocs": ["Yes", "No"],
			"p": "0.17",
			"q": "6",
			"src": "chainstream",
			"tk": "0xtaker",
			"ta": 1780174594000,
			"ti": "",
			"tn": "",
			"toh": "0xorder",
			"tp": "",
			"tt": ["fresh_wallet"],
			"ts": 1780411421000,
			"tid": "108233603819467706476318984012158651931658302669301887462181073562758483842092",
			"tx": "0xtx",
			"ty": "buy"
		}
	}`), &frame)
	if err != nil {
		t.Fatalf("unmarshal frame: %v", err)
	}

	return frame
}

func TestPredictionActivityChannel(t *testing.T) {
	if got := predictionActivityChannel(PredictionActivityChannelKindEvent, "world-cup-winner"); got != "pred:evt:world-cup-winner:act" {
		t.Fatalf("event channel = %q", got)
	}
	if got := predictionActivityChannel(PredictionActivityChannelKindToken, "token-1"); got != "pred:tok:token-1:act" {
		t.Fatalf("token channel = %q", got)
	}
}

func TestParsePredictionActivity(t *testing.T) {
	activity := parsePredictionActivity(predictionActivityFrame(t))

	if activity.ActivityID != "activity-1" {
		t.Fatalf("activity id = %q", activity.ActivityID)
	}
	if activity.SeqIndex != 42 {
		t.Fatalf("seq index = %d", activity.SeqIndex)
	}
	if activity.Type != PredictionActivityTypeBuy {
		t.Fatalf("type = %q", activity.Type)
	}
	if len(activity.AssetIDs) != 2 || activity.AssetIDs[0] != "asset-yes" {
		t.Fatalf("asset ids = %#v", activity.AssetIDs)
	}
	if len(activity.TakerTags) != 1 || activity.TakerTags[0] != "fresh_wallet" {
		t.Fatalf("taker tags = %#v", activity.TakerTags)
	}
}

func TestPredictionActivityFilterFields(t *testing.T) {
	filter := "conditionId == '0xcondition' && type == 'buy' && tokenId != ''"
	got := ReplaceFilterFields(filter, "subscribePredictionEventActivities")
	want := "meta.a.cid == '0xcondition' && meta.a.ty == 'buy' && meta.a.tid != ''"
	if got != want {
		t.Fatalf("filter = %q, want %q", got, want)
	}
}
