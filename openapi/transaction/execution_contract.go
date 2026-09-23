package transaction

import (
	"encoding/json"
	"github.com/chainstream-io/chainstream-go-sdk/v2/internal/executioncontract"
)

// UnmarshalJSON preserves the required nullable amounts at the HTTP client boundary.
func (order *ExecutionOrderResponse) UnmarshalJSON(data []byte) error {
	if err := executioncontract.ValidateAmounts(data); err != nil {
		return err
	}
	type wire ExecutionOrderResponse
	var decoded wire
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	*order = ExecutionOrderResponse(decoded)
	return nil
}
