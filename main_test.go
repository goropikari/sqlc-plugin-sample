package main

import (
	"os"
	"testing"

	"github.com/goropikari/sqlc_plugin_sample/proto/sqlc/plugin"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func TestGenerate(t *testing.T) {
	t.Run("valid request", func(t *testing.T) {
		data, err := os.ReadFile("testing/gen/dump_proto")
		require.NoError(t, err)

		var req plugin.GenerateRequest
		err = proto.Unmarshal(data, &req)
		require.NoError(t, err)

		_, err = Generate(&req)
		require.NoError(t, err)
	})
}
