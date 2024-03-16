package razorpay

import (
	razorpay "github.com/razorpay/razorpay-go"
)

func CreateOrder(key string, pass string) (map[string]interface{}, error) {
	client := razorpay.NewClient(key, pass)

	data := map[string]interface{}{
		"amount":   50000,
		"currency": "INR",
		"receipt":  "some_receipt_id",
	}
	body, err := client.Order.Create(data, nil)

	return body, err
}
