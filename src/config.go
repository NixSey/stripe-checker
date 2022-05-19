package src

import (
	"strconv"

	"github.com/ovh/configstore"
)

var (
	cfg Cfg
	err error
)

type Cfg struct {
	StripeSdkKey     string
	StripePublishKey string
	Amount           int64
}

func LoadCfg(path string) Cfg {
	configstore.LogInfoFunc = func(format string, v ...any) {}
	configstore.File(path)

	cfg.StripeSdkKey, err = configstore.GetItemValue("stripe-private-api-key")
	HandleError(err)

	cfg.StripePublishKey, err = configstore.GetItemValue("stripe-publish-api-key")
	HandleError(err)

	rawAmount, err := configstore.GetItemValue("amount")
	HandleError(err)

	cfg.Amount, _ = strconv.ParseInt(rawAmount, 10, 64)
	return cfg
}
