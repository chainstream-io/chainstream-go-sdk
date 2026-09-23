package executioncontract

import "testing"

func TestRequiredNullableAmounts(t *testing.T) {
	for _, raw := range []string{
		`{}`, `{"actualBuyAmount":null}`, `{"actualRefundAmount":null}`, `null`, `[]`,
		`{"actualBuyAmount":1,"actualRefundAmount":null}`,
		`{"actualBuyAmount":null,"actualRefundAmount":false}`,
	} {
		if ValidateAmounts([]byte(raw)) == nil {
			t.Errorf("expected missing/invalid amount rejection")
		}
	}
	for _, raw := range []string{
		`{"actualBuyAmount":null,"actualRefundAmount":null}`,
		`{"actualBuyAmount":"115792089237316195423570985008687907853269984665640564039457584007913129639935","actualRefundAmount":"0"}`,
	} {
		if err := ValidateAmounts([]byte(raw)); err != nil {
			t.Fatal(err)
		}
	}
}
