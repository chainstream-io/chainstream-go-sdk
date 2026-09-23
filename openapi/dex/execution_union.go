package dex

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// Raw anyOf wrappers preserve all fields, but execution markers require validation.
func validateExecutionUnion(data []byte, model any) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if execution, exists := fields["execution"]; exists {
		if bytes.Equal(bytes.TrimSpace(execution), []byte("null")) {
			return fmt.Errorf("execution must not be null")
		}
		return json.Unmarshal(data, model)
	}
	return nil
}
