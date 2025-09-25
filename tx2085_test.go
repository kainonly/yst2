package yst2_test

import (
	"context"
	"testing"

	"github.com/kainonly/yst2"
	"github.com/stretchr/testify/assert"
)

func TestYst2_Tx2085(t *testing.T) {
	ctx := context.TODO()
	orderNo := OrderNumber(`00001`, `0`)
	amount := `1`
	dto := yst2.NewTx2085Dto(orderNo, amount, `https://pay.kainonly.com:8443/notify/yst2`).
		SetSummary(`测试生成`).
		SetExtendParams(yst2.M{"code": `00001`}). // 调试时需要设置下
		SetPayAmount(amount).
		SetPayMode(yst2.M{
			//"SCAN_WEIXIN": yst2.M{
			//	"limitPay": "no_credit",
			//},
			"SCAN_ALIPAY": yst2.M{
				"limitPay": "no_credit",
			},
			//"SCAN_UNIONPAY": yst2.M{
			//	"limitPay": "no_credit",
			//},
		}).
		SetGoodsName(`账号充值`).
		SetGoodsDesc(`客户-00001`)

	r, err := x.Tx2085(ctx, dto)
	assert.NoError(t, err)

	t.Log(r.RespTraceNum)
	assert.Equal(t, orderNo, r.ReqTraceNum)
	t.Log(`code:`, r.RespCode)
	t.Log(`msg:`, r.RespMsg)
	t.Log(`extend:`, r.ExtendParams)
	t.Log(`channelParamInfo:`, r.ChannelParamInfo)
	t.Log(`chnlFrontParamInfo:`, r.ChnlFrontParamInfo)
}
