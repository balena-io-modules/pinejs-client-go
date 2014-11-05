package main

import (
	"fmt"

	"bitbucket.com/rulemotion/pinejs-client-go"
)

func main() {
	resinApi := &pinejs.Client{}
	resinApi.Init("http://api.resindev.io", "petros")

	coupon, _ := coupon.New(&stripe.CouponParams{
		Percent:        25,
		Duration:       "repeating",
		DurationPeriod: 3,
		ID:             "250FF",
	})

	fmt.Println(coupon)
}
