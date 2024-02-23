package dto

type OrderResponse struct {
	Id          string `json:"order_id"`
	Name        string `json:"order_name"`
	Number      string `json:"order_number"`
	Description string `json:"order_description"`
	Status      string `json:"order_status"`
}
