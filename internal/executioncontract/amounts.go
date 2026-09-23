// Package executioncontract validates only the new execution order wire contract.
package executioncontract

import (
	"encoding/json"
	"fmt"
)

// ValidateAmounts rejects missing actual amounts without conflating null with zero.
// Amounts are strings (including values beyond float64 precision) or explicit null.
func ValidateAmounts(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	for _, key := range []string{"actualBuyAmount", "actualRefundAmount"} {
		raw, present := fields[key]
		if !present {
			return fmt.Errorf("ExecutionOrderResponse requires %s", key)
		}
		var value *string
		if err := json.Unmarshal(raw, &value); err != nil {
			return fmt.Errorf("ExecutionOrderResponse %s must be string or null", key)
		}
	}
	return nil
}
