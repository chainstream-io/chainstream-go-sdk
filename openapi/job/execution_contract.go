package job

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

// UnmarshalJSON retains an explicit arbitrary JSON null without manufacturing a
// result when the optional key is absent.
func (status *JobStatusResponse) UnmarshalJSON(data []byte) error {
	type wire JobStatusResponse
	var decoded wire
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if raw, ok := fields["result"]; ok {
		decoded.Result = json.RawMessage(append([]byte(nil), raw...))
	}
	*status = JobStatusResponse(decoded)
	return nil
}
