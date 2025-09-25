package yst2_test

import (
	"context"
	"testing"

	"github.com/kainonly/yst2"
	"github.com/stretchr/testify/assert"
)

func TestYst2_Tx2295(t *testing.T) {
	ctx := context.TODO()
	// 注意：测试环境不支持订单关闭，仅测试互通
	// 填写 tx2085 生成的通联订单号，
	respTraceNum := `20250627151006208501596062`
	dto := yst2.NewTx2295Dto(respTraceNum, `测试`)

	r, err := x.Tx2295(ctx, dto)
	assert.NoError(t, err)

	t.Log(r.RespTraceNum)
	t.Log(`code:`, r.RespCode)
	t.Log(`msg:`, r.RespMsg)
	t.Log(`closeResult:`, r.CloseResult)
	t.Log(`closeFinishTime:`, r.CloseFinishTime)
	t.Log(`result:`, r.Result)
}
