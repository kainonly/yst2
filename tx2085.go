package yst2

import (
	"context"
	"time"

	"github.com/bytedance/sonic"
)

type Tx2085Dto struct {
	ReqTraceNum    string `json:"reqTraceNum"`              // 商户订单号
	OrderAmount    string `json:"orderAmount"`              // 订单金额
	PayAmount      string `json:"payAmount,omitempty"`      // 支付金额
	PayMode        M      `json:"payMode,omitempty"`        // 支付模式
	ReqsUrl        string `json:"reqsUrl,omitempty"`        // 前台通知地址
	RespUrl        string `json:"respUrl"`                  // 后台通知地址
	OrderValidTime string `json:"orderValidTime,omitempty"` // 订单过期时间
	GoodsName      string `json:"goodsName,omitempty"`      // 商品名称
	Summary        string `json:"summary,omitempty"`        // 摘要
	ExtendParams   M      `json:"extendParams,omitempty"`   // 扩展参数
	TxDistrictCode string `json:"txDistrictCode,omitempty"` // 交易所在省市
	GoodsDesc      string `json:"goodsDesc,omitempty"`      // 商品描述
}

func NewTx2085Dto(reqTraceNum string, orderAmount string, respUrl string) *Tx2085Dto {
	return &Tx2085Dto{
		ReqTraceNum: reqTraceNum,
		OrderAmount: orderAmount,
		RespUrl:     respUrl,
	}
}

func (x *Tx2085Dto) SetPayAmount(v string) *Tx2085Dto {
	x.PayAmount = v
	return x
}

func (x *Tx2085Dto) SetPayMode(v M) *Tx2085Dto {
	x.PayMode = v
	return x
}

func (x *Tx2085Dto) SetReqsUrl(v string) *Tx2085Dto {
	x.ReqsUrl = v
	return x
}

func (x *Tx2085Dto) SetOrderValidTime(v string) *Tx2085Dto {
	x.OrderValidTime = v
	return x
}

func (x *Tx2085Dto) SetGoodsName(v string) *Tx2085Dto {
	x.GoodsName = v
	return x
}

func (x *Tx2085Dto) SetSummary(v string) *Tx2085Dto {
	x.Summary = v
	return x
}

func (x *Tx2085Dto) SetExtendParams(v M) *Tx2085Dto {
	x.ExtendParams = v
	return x
}

func (x *Tx2085Dto) SetTxDistrictCode(v string) *Tx2085Dto {
	x.TxDistrictCode = v
	return x
}

func (x *Tx2085Dto) SetGoodsDesc(v string) *Tx2085Dto {
	x.GoodsDesc = v
	return x
}

type Tx2085Result struct {
	Result             string `json:"result,omitempty"`
	RespTraceNum       string `json:"respTraceNum"`
	ReqTraceNum        string `json:"reqTraceNum,omitempty"`
	ExtendParams       string `json:"extendParams,omitempty"`
	ChannelParamInfo   string `json:"channelParamInfo,omitempty"`
	ChnlFrontParamInfo string `json:"chnlFrontParamInfo,omitempty"`
	RespCode           string `json:"respCode"`
	RespMsg            string `json:"respMsg"`
	IsPreConsume       string `json:"isPreConsume,omitempty"`
}

func (x *Yst2) Tx2085(ctx context.Context, dto *Tx2085Dto) (_ *Tx2085Result, err error) {
	now := time.Now()
	var data string
	if data, err = sonic.MarshalString(*dto); err != nil {
		return
	}

	var bizData string
	if bizData, err = x.Request(x.SetNow(ctx, now),
		`/tx/handle`, `2085`, data); err != nil {
		return
	}

	var result Tx2085Result
	if err = sonic.UnmarshalString(bizData, &result); err != nil {
		return
	}
	return &result, nil
}
