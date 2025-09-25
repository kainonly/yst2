package yst2_test

import (
	"context"
	"testing"

	"github.com/kainonly/yst2"
	"github.com/stretchr/testify/assert"
)

func TestYst2_Tx3001(t *testing.T) {
	ctx := context.TODO()
	respTraceNum := `20250925135844208501094930`
	dto := yst2.NewTq3001Dto().
		ByRespTraceNum(respTraceNum)

	r, err := x.Tq3001(ctx, dto)
	assert.NoError(t, err)

	t.Log(`code:`, r.RespCode)
	t.Log(`msg:`, r.RespMsg)
	t.Log(`reqTraceNum:`, r.ReqTraceNum)
	t.Log(`respTraceNum:`, r.RespTraceNum)
	t.Log(`result:`, r.Result)
	t.Log(`txDesc:`, r.TxDesc)
	t.Log(`orderAmount:`, r.OrderAmount)
	t.Log(`payAmount:`, r.PayAmount)
	t.Log(`finishTime:`, r.FinishTime)
}

func TestYst2_Tx3001ByOrderNo(t *testing.T) {
	ctx := context.TODO()
	orderNo := `W00001-202509251358446600`
	dto := yst2.NewTq3001Dto().
		ByReqTraceNum(orderNo, `20250925`)

	r, err := x.Tq3001(ctx, dto)
	assert.NoError(t, err)

	t.Log(`code:`, r.RespCode)
	t.Log(`msg:`, r.RespMsg)
	t.Log(`reqTraceNum:`, r.ReqTraceNum)
	t.Log(`respTraceNum:`, r.RespTraceNum)
	t.Log(`result:`, r.Result)
	t.Log(`txDesc:`, r.TxDesc)
	t.Log(`orderAmount:`, r.OrderAmount)
	t.Log(`payAmount:`, r.PayAmount)
	t.Log(`finishTime:`, r.FinishTime)
}
