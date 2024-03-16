package main

import (
	"fmt"
	"service-payment-gateway/order/razorpay"
	order "service-payment-gateway/order/response"
	"service-payment-gateway/secret"
)

func main() {
	var key string
	var pass string

	key, pass = secret.GetCredentials()

	var result map[string]interface{}
	var err error

	result, err = razorpay.CreateOrder(key, pass)

	// fmt.Println(result)
	fmt.Println(err)

	var response order.OrderResponse = order.ConvertFromRazorPayResponse(result)
	fmt.Println(response)
}
