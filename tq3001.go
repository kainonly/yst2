package yst2

import (
	"context"
	"time"

	"github.com/bytedance/sonic"
)

type Tq3001Dto struct {
	RespTraceNum string `json:"respTraceNum,omitempty"` // 通联订单号
	ReqTraceNum  string `json:"reqTraceNum,omitempty"`  // 商户订单号
	OriTransDate string `json:"oriTransDate,omitempty"` // 订单创建日期
}

func NewTq3001Dto() *Tq3001Dto {
	return &Tq3001Dto{}
}

func (x *Tq3001Dto) ByRespTraceNum(v string) *Tq3001Dto {
	x.RespTraceNum = v
	return x
}

func (x *Tq3001Dto) ByReqTraceNum(v string, date string) *Tq3001Dto {
	x.ReqTraceNum = v
	x.OriTransDate = date
	return x
}

type Tq3001Result struct {
	ReqTraceNum  string `json:"reqTraceNum"`            // 商户订单号
	RespTraceNum string `json:"respTraceNum"`           // 通联订单号
	Result       string `json:"result,omitempty"`       // 订单状态
	TxDesc       string `json:"txDesc,omitempty"`       // 订单状态说明
	OrderAmount  int64  `json:"orderAmount"`            // 订单金额
	PayAmount    int64  `json:"payAmount,omitempty"`    // 支付金额
	FinishTime   string `json:"finishTime,omitempty"`   // 订单支付完成时间
	RespCode     string `json:"respCode"`               // 业务返回码
	RespMsg      string `json:"respMsg"`                // 业务返回说明
	IsPreConsume string `json:"isPreConsume,omitempty"` // 是否微信订单预消费
}

func (x *Yst2) Tq3001(ctx context.Context, dto *Tq3001Dto) (_ *Tq3001Result, err error) {
	now := time.Now()
	var data string
	if data, err = sonic.MarshalString(*dto); err != nil {
		return
	}

	var bizData string
	if bizData, err = x.Request(x.SetNow(ctx, now),
		`/tq/handle`, `3001`, data); err != nil {
		return
	}

	var result Tq3001Result
	if err = sonic.UnmarshalString(bizData, &result); err != nil {
		return
	}
	return &result, nil
}
