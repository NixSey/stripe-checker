# stripe-checker
Credit card checker using stripe payment gateway.

## how it works
__checker works by creating a Payment Ident right after it reads a credit card then creates a payment, if payment is approved, then refunds the amount charged (avoiding future problems with the card) and is given as: "live". If it is not approved, depending on the error code, the card is given as live, or not.__

### how the gateway is handled 
![](https://github.com/J4c5/stripe-checker/blob/assets/2022-05-18%20(2).png)


### formats and how a list of cards should be loaded
- the card format should be: `card_number|card_exp_month|card_exp_year|cvc`
- the separator: "|" can be modified, but you will have to pass the value of the separator to the tool using: `--separator` or `-s`

## responsibility and terms
> ⚠️ By using this tool you agree that the author of the tool and the tool are not to blame for misuse, this was created only for the purpose of studying carding. It must not be used for evil purposes.
