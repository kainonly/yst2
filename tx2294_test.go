package yst2_test

import (
	"context"
	"testing"

	"github.com/kainonly/yst2"
	"github.com/stretchr/testify/assert"
)

func TestYst2_Tx2294(t *testing.T) {
	ctx := context.TODO()
	dto := yst2.NewTx2294Dto(
		`TK-W00001-202509251358446600`,
		`20250925135844208501094930`,
		`1`,
	)

	r, err := x.Tx2294(ctx, dto)
	assert.NoError(t, err)

	t.Log(r.RespTraceNum)
	t.Log(`code:`, r.RespCode)
	t.Log(`msg:`, r.RespMsg)
	t.Log(`reqTraceNum:`, r.ReqTraceNum)
	t.Log(`respTraceNum:`, r.RespTraceNum)
	t.Log(`extendParams:`, r.ExtendParams)
	t.Log(`channelParamInfo:`, r.ChannelParamInfo)
	t.Log(`result:`, r.Result)
}
