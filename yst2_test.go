package yst2_test

import (
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/kainonly/yst2"
)

var x *yst2.Yst2
var privateKeyStr = `MIGTAgEAMBMGByqGSM49AgEGCCqBHM9VAYItBHkwdwIBAQQgxL/uTSkBSGPtJWV7O+KTBPds8wZ9iTxondaH7NneHw6gCgYIKoEcz1UBgi2hRANCAATnOBrmh+4x5RBW2hJR+PI/3UVTxDHfn1rI9CZuwA0MbdYVIQi/1jOLIxgYpWmQ1TjIMXDBONioztQDH7fGlbQu`
var allinpayPublicKeyStr = `MFkwEwYHKoZIzj0CAQYIKoEcz1UBgi0DQgAEu9LNkJlyLtjJxtQWIGlcZ/hyHt5eZ7LEH1nfOiK1H9HsE1cMPu5KK5jZVTtAyc7lPMXixUMirf6A3tMbuMbgqg==`

func TestMain(m *testing.M) {
	var err error
	if x, err = yst2.NewYst2(yst2.Option{
		BaseURL:           `https://ibstest.allinpay.com/yst/yst-service-api`,
		PrivateKey:        privateKeyStr,
		AllinpayPublicKey: allinpayPublicKeyStr,
		AppID:             "21774685401604222977",
		SecretKey:         "878427523d3525e070298d44481b8d2e",
	}); err != nil {
		panic(err)
	}
	os.Exit(m.Run())
}

func OrderNumber(code string, kind string) string {
	now := time.Now()
	formatter := now.Format("20060102150405")
	rand.New(rand.NewSource(time.Now().UnixNano()))
	num := rand.Intn(999) + 1
	paddedStr := fmt.Sprintf("%03d", num)
	return fmt.Sprintf("W%s-%s%s%s", code, formatter, paddedStr, kind)
}
