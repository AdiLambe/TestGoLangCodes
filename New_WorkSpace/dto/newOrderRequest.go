package dto

import (
	"github.com/AdiLambe/TestGoLangCodes/workspace/errs"
)

type NewOrderRequest struct {
	Id          string `json:"order_id"`
	Name        string `json:"order_name"`
	Number      string `json:"order_number"`
	Description string `json:"order_description"`
	Status      string `json:"order_status"`
}

func (r NewOrderRequest) Validate() *errs.AppError {

	if r.Number < "5000" {
		return errs.NewValidationError("To Claim Coupon you need to purchase minimum 5000.00")
	}

	return nil
}
