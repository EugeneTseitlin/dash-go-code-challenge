package model

type Payment struct {
	ID string `json:"id" validate:"alphanum,max=256,required"`
	Type string `json:"type" validate:"required"`
	FromUserID string `json:"fromUserId" validate:"alphanum,max=256,required"`
	ToMerchantID string `json:"toMerchantId" validate:"omitempty,alphanum,max=256"`
	ToUserID string `json:"toUserId" validate:"omitempty,alphanum,max=256"`
	Amount float64 `json:"amount" validate:"gt=0,required"`
	CreatedAt string `json:"createdAt" validate:"iso8601,required"`
}