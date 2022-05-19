package src

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/Jeffail/gabs"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/paymentintent"
	"github.com/stripe/stripe-go/v72/refund"
)

type Result struct {
	Code             string
	DeclinedReason   string
	DeclineCodeValid bool
	Valid            bool
}

var (
	result Result
)

// set payment ident parameters
func DefineParams() *stripe.PaymentIntentParams {
	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(cfg.Amount),
		Currency: stripe.String(string(stripe.CurrencyEUR)),
		AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
			Enabled: stripe.Bool(true),
		},
	}

	return params
}

// refund payment using payment ident id
func RefundPi(pi string) {
	params := &stripe.RefundParams{
		PaymentIntent: stripe.String(pi),
	}
	refund.New(params)
}

// create a payment indent
func CreatePi(cfg Cfg) (string, string) {
	params := DefineParams()
	pi, err := paymentintent.New(params)
	HandleError(err)
	return pi.ClientSecret, pi.ID
}

// check the cards, if live refund the amount charged, if not just ignore
func CheckCard(card Card, cfg Cfg) Result {
	result.Valid = false
	result.DeclineCodeValid = false
	stripe.Key = cfg.StripeSdkKey
	clientSecret, Pid := CreatePi(cfg)

	url := fmt.Sprintf("https://api.stripe.com/v1/payment_intents/%s/confirm", Pid)
	method := "POST"

	payload := strings.NewReader(fmt.Sprintf("return_url=https://stripe.com&payment_method_data[type]=card&payment_method_data[card][number]=%s&payment_method_data[card][cvc]=%s&payment_method_data[card][exp_year]=%s&payment_method_data[card][exp_month]=%s&payment_method_data[billing_details][address][country]=BR&payment_method_data[payment_user_agent]=stripe.js/eb14574ae;+stripe-js-v3/eb14574ae;+payment-element&payment_method_data[time_on_page]=145828&payment_method_data[guid]=a6cb16e6-1eee-420c-bdd7-49271c53ee9537ac30&payment_method_data[muid]=6a5417b9-177e-4293-af9b-200ef3fdac60ef8bc5&payment_method_data[sid]=c0d2f286-3f58-4e9b-95a5-d5bef895e91b377d41&expected_payment_method_type=card&use_stripe_sdk=true&key=%s&client_secret=%s", card.CardNumber, card.Cvv, card.ExpYear, card.ExpMonth, cfg.StripePublishKey, clientSecret))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	HandleError(err)

	req.Header.Add("authority", "api.stripe.com")
	req.Header.Add("accept", "application/json")
	req.Header.Add("accept-language", "en-US,en;q=0.9,pt;q=0.8")
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("origin", "https://js.stripe.com")
	req.Header.Add("referer", "https://js.stripe.com/")
	req.Header.Add("sec-ch-ua", "\" Not A;Brand\";v=\"99\", \"Chromium\";v=\"101\", \"Google Chrome\";v=\"101\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "\"Windows\"")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-site", "same-site")
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.54 Safari/537.36")

	res, err := client.Do(req)
	HandleError(err)

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	HandleError(err)

	jsonParsed, err := gabs.ParseJSON(body)
	HandleError(err)

	exists := jsonParsed.ExistsP("error")
	if exists {
		ErrorCode, ok := jsonParsed.Path("error.code").Data().(string)
		if ok {
			result.Code = ErrorCode
			if ErrorCode == "incorrect-cvc" {
				result.Valid = true
			}
			if ErrorCode == "not_permitted" {
				result.Valid = true
			}
			if ErrorCode == "card_declined" {
				declined_reason, ok := jsonParsed.Path("error.decline_code").Data().(string)
				if ok {
					result.DeclineCodeValid = true
					result.DeclinedReason = declined_reason

					if declined_reason == "currency_not_supported" {
						result.Valid = true
					}

				}
			}
			if ErrorCode == "insufficient-funds" {
				result.Valid = true
			}
			if ErrorCode == "amount-too-large" {
				result.Valid = true
			}
			if ErrorCode == "balance-insufficient" {
				result.Valid = true
			}
		}
	} else {
		result.Valid = true
		RefundPi(Pid)
	}

	return result
}
