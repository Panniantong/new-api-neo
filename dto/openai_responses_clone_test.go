package dto

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOpenAIResponsesRequestCloneForRelayKeepsOriginalMutableFieldsIsolated(t *testing.T) {
	original := &OpenAIResponsesRequest{
		Model:     "gpt-5.5",
		Input:     json.RawMessage(`[{"role":"user","content":"hello"}]`),
		Reasoning: &Reasoning{Effort: "medium"},
	}

	clone := original.CloneForRelay()
	clone.Model = "gpt-5.6-sol"
	clone.Reasoning.Effort = "high"

	require.Equal(t, "gpt-5.5", original.Model)
	require.Equal(t, "medium", original.Reasoning.Effort)
	require.JSONEq(t, string(original.Input), string(clone.Input))
}
