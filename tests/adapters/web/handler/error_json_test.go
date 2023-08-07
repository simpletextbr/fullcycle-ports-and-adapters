package handler_test

import (
	"testing"

	"github.com/simpletextbr/fullcycle-ports-and-adapters/adapters/web/handler"
	"github.com/stretchr/testify/require"
)

func TestJsonError(t *testing.T) {
	msg := "Hello Json"
	result := handler.JsonError(msg)

	require.Equal(t, []byte(`{"message":"Hello Json"}`), result)
}
