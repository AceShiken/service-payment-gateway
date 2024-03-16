package order

import (
	"time"
)

type OrderResponse struct {
	id          string
	amount      float64
	entity      string
	amount_paid float64
	amount_due  float64
	currency    string
	receipt     string
	status      string
	attempts    float64
	notes       []interface{}
	created_at  time.Time
}

func ConvertFromRazorPayResponse(result map[string]interface{}) OrderResponse {
	var response OrderResponse
	response.id = result["id"].(string)
	response.amount = result["amount"].(float64)
	response.entity = result["entity"].(string)
	response.amount_paid = result["amount_paid"].(float64)
	response.amount_due = result["amount_due"].(float64)
	response.currency = result["currency"].(string)
	response.receipt = result["receipt"].(string)
	response.status = result["status"].(string)
	response.attempts = result["attempts"].(float64)
	response.notes = result["notes"].([]interface{})
	response.created_at = time.Unix(int64(result["created_at"].(float64)), 0)
	return response
}
