package yst2

import (
	"context"
	"time"

	"github.com/bytedance/sonic"
)

type Tq3002Dto struct {
	RespTraceNum string `json:"respTraceNum,omitempty"` // 通联订单号
	ReqTraceNum  string `json:"reqTraceNum,omitempty"`  // 商户订单号
	OriTransDate string `json:"oriTransDate,omitempty"` // 订单创建日期
}

func NewTq3002Dto() *Tq3002Dto {
	return &Tq3002Dto{}
}

func (x *Tq3002Dto) ByRespTraceNum(v string) *Tq3002Dto {
	x.RespTraceNum = v
	return x
}

func (x *Tq3002Dto) ByReqTraceNum(v string, date string) *Tq3002Dto {
	x.ReqTraceNum = v
	x.OriTransDate = date
	return x
}

type Tq3002Result struct {
	ReqTraceNum      string `json:"reqTraceNum"`                // 商户订单号
	RespTraceNum     string `json:"respTraceNum"`               // 通联订单号
	Result           string `json:"result,omitempty"`           // 订单状态
	TxDesc           string `json:"txDesc,omitempty"`           // 订单状态说明
	OrgReqTraceNum   string `json:"orgReqTraceNum,omitempty"`   // 原商户订单号
	OrgRespTraceNum  string `json:"orgRespTraceNum,omitempty"`  // 云商通原订单号
	OrderAmount      int64  `json:"orderAmount"`                // 订单金额
	PayAmount        int64  `json:"payAmount,omitempty"`        // 支付金额
	FinishTime       string `json:"finishTime,omitempty"`       // 订单支付完成时间
	ExtendParams     string `json:"extendParams,omitempty"`     // 扩展参数
	ChannelparamInfo string `json:"channelparamInfo,omitempty"` // 渠道参数信息
	RespCode         string `json:"respCode"`                   // 业务返回码
	RespMsg          string `json:"respMsg"`                    // 业务返回说明
	IsPreConsume     string `json:"isPreConsume,omitempty"`     // 是否微信订单预消费
}

func (x *Yst2) Tq3002(ctx context.Context, dto *Tq3002Dto) (_ *Tq3002Result, err error) {
	now := time.Now()
	var data string
	if data, err = sonic.MarshalString(*dto); err != nil {
		return
	}

	var bizData string
	if bizData, err = x.Request(x.SetNow(ctx, now),
		`/tq/handle`, `3002`, data); err != nil {
		return
	}

	var result Tq3002Result
	if err = sonic.UnmarshalString(bizData, &result); err != nil {
		return
	}
	return &result, nil
}
