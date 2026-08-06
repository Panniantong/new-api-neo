package controller

import (
	"net/http"
	"testing"

	"github.com/QuantumNous/new-api/types"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestShouldRetryStopsAfterSub2ExhaustsInternalCapacityAttempts(t *testing.T) {
	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(nil)
	err := types.WithOpenAIError(types.OpenAIError{
		Type:    "rate_limit_error",
		Code:    "sub2_local_capacity_exhausted",
		Message: "Too many pending requests, please retry later",
	}, http.StatusTooManyRequests)

	require.False(t, shouldRetry(c, err, 2))
}
