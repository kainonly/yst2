package yst2

import (
	"context"
	"time"

	"github.com/bytedance/sonic"
)

type Tx2294Dto struct {
	ReqTraceNum     string `json:"reqTraceNum"`            // 商户订单号
	OrgRespTraceNum string `json:"orgRespTraceNum"`        // 原通联订单号
	OrderAmount     string `json:"orderAmount"`            // 退款总金额
	RespUrl         string `json:"respUrl,omitempty"`      // 后台通知地址
	ChnlDiscAmt     M      `json:"chnlDiscAmt,omitempty"`  // 优惠信息
	ExtendParams    string `json:"extendParams,omitempty"` // 扩展信息
}

func NewTx2294Dto(reqTraceNum string, orgRespTraceNum string, orderAmount string) *Tx2294Dto {
	return &Tx2294Dto{
		ReqTraceNum:     reqTraceNum,
		OrgRespTraceNum: orgRespTraceNum,
		OrderAmount:     orderAmount,
	}
}

func (x *Tx2294Dto) SetRespUrl(v string) *Tx2294Dto {
	x.RespUrl = v
	return x
}

func (x *Tx2294Dto) SetChnlDiscAmt(v M) *Tx2294Dto {
	x.ChnlDiscAmt = v
	return x
}

func (x *Tx2294Dto) SetExtendParams(v string) *Tx2294Dto {
	x.ExtendParams = v
	return x
}

type Tx2294Result struct {
	Result           string `json:"result,omitempty"`           // 订单状态
	RespMsg          string `json:"respMsg,omitempty"`          // 业务返回说明
	ReqTraceNum      string `json:"reqTraceNum"`                // 商户订单号
	RespTraceNum     string `json:"respTraceNum"`               // 通联订单号
	ExtendParams     string `json:"extendParams,omitempty"`     // 渠道参数信息
	ChannelParamInfo string `json:"channelParamInfo,omitempty"` // 扩展信息
	RespCode         string `json:"respCode"`                   // 业务返回码
}

func (x *Yst2) Tx2294(ctx context.Context, dto *Tx2294Dto) (_ *Tx2294Result, err error) {
	now := time.Now()
	var data string
	if data, err = sonic.MarshalString(*dto); err != nil {
		return
	}

	var bizData string
	if bizData, err = x.Request(x.SetNow(ctx, now),
		`/tx/handle`, `2294`, data); err != nil {
		return
	}

	var result Tx2294Result
	if err = sonic.UnmarshalString(bizData, &result); err != nil {
		return
	}
	return &result, nil
}
