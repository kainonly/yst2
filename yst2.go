package yst2

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"time"

	"github.com/bytedance/sonic"
	"github.com/emmansun/gmsm/sm2"
	"github.com/weplanx/go/help"
	"resty.dev/v3"
)

type Yst2 struct {
	Option *Option
	Client *resty.Client
	priKey *sm2.PrivateKey
	pubKey *ecdsa.PublicKey
}

type Option struct {
	BaseURL           string `yaml:"base_url" env:"BASE_URL"`
	PrivateKey        string `yaml:"private_key" env:"PRIVATE_KEY"`
	AllinpayPublicKey string `yaml:"allinpay_public_key" env:"ALLINPAY_PUBLIC_KEY"`
	AppID             string `yaml:"app_id" env:"APP_ID"`
	SecretKey         string `yaml:"secret_key" env:"SECRET_KEY"`
}

func NewYst2(opt Option) (x *Yst2, err error) {
	x = &Yst2{
		Option: &opt,
		Client: resty.New().SetBaseURL(opt.BaseURL),
	}
	if x.priKey, err = help.PrivKeySM2FromBase64(opt.PrivateKey); err != nil {
		return
	}
	if x.pubKey, err = help.PubKeySM2FromBase64(opt.AllinpayPublicKey); err != nil {
		return
	}
	return
}

func (x *Yst2) Debug() {
	x.Client.EnableTrace()
}

func (x *Yst2) GetPublicKey() *ecdsa.PublicKey {
	return x.pubKey
}

func (x *Yst2) GetPrivateKey() *sm2.PrivateKey {
	return x.priKey
}

type M map[string]any

func (x *Yst2) SetNow(ctx context.Context, ts time.Time) context.Context {
	return context.WithValue(ctx, "now", ts)
}

func (x *Yst2) GetNow(ctx context.Context) time.Time {
	return ctx.Value("now").(time.Time)
}

type ResponseBody struct {
	Code    string `json:"code"`    // 调用结果返回码
	Msg     string `json:"msg"`     // 调用结果返回码描述
	Sign    string `json:"sign"`    // 商户请求参数的签名串
	BizData string `json:"bizData"` // 返回参数的集合
}

func (x *Yst2) Request(ctx context.Context, path string, code string, data string) (_ string, err error) {
	now := x.GetNow(ctx)
	body := M{
		"appId":     x.Option.AppID,
		"bizData":   data,
		"charset":   "UTF-8",
		"format":    "json",
		"transCode": code,
		"transDate": now.Format(`20060102`),
		"transTime": now.Format(`150405`),
		"version":   "1.0",
	}

	var signature string
	if signature, err = help.Sm2Sign(x.priKey, help.MapToSignText(body)); err != nil {
		return
	}
	body["sign"] = signature
	body["signType"] = "SM3withSM2"

	var resp *resty.Response
	if resp, err = x.Client.R().
		SetContext(ctx).
		SetBody(body).
		Post(path); err != nil {
		return
	}

	if resp.StatusCode() != 200 {
		err = help.E(0, `第三方接口响应失败!`)
		return
	}
	var content M
	if err = sonic.Unmarshal(resp.Bytes(), &content); err != nil {
		return
	}

	if content["code"] != "00000" {
		err = help.E(0, fmt.Sprintf(`第三方请求失败![%s]: %s`, content["code"], content["msg"]))
		return
	}

	sign := content["sign"].(string)
	delete(content, "sign")
	delete(content, "signType")

	var verify bool
	if verify, err = help.Sm2Verify(x.pubKey, help.MapToSignText(content), sign); err != nil {
		return
	}
	if !verify {
		err = help.E(0, `第三方响应内容签名存在不一致!`)
		return
	}

	return content["bizData"].(string), nil
}
