package inventory

import (
	"time"
)

type OrderCancelled struct{
	EventType string	`json:"eventType" type:"string"`
	TimeStamp string 	`json:"timeStamp" type:"string"`
	Id int `json:"id" type:"int"` 
	ProductId string `json:"productId" type:"string"` 
	Qty int `json:"qty" type:"int"` 
	CustomerId string `json:"customerId" type:"string"` 
	Amount float64 `json:"amount" type:"float64"` 
	Status string `json:"status" type:"string"` 
	Address string `json:"address" type:"string"` 
	
}

func NewOrderCancelled() *OrderCancelled{
	event := &OrderCancelled{EventType:"OrderCancelled", TimeStamp:time.Now().String()}

	return event
}
