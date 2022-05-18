package src

import (
	"os"

	"github.com/gookit/color"
)

func SuccessCheckLog(cc Card, cfg Cfg) {
	color.Green.Printf("[live] (%s, %s/%s, %s) [%d]\n", cc.CardNumber, cc.ExpMonth, cc.ExpYear, cc.Cvv, cfg.Amount)
}

func FailCheckLog(cc Card, cfg Cfg, r Result) {
	color.Red.Printf("[die] (%s, %s/%s, %s) [%d] (Reason: %s)\n", cc.CardNumber, cc.ExpMonth, cc.ExpYear, cc.Cvv, cfg.Amount, r.Code)
}

func FatalWithoutExit(e error) {
	color.Error.Prompt("%s", e.Error())

}

func Fatal(e error) {
	color.Error.Prompt("%s", e.Error())
	os.Exit(1)
}
