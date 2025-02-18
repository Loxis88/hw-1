package models

import (
	"errors"
	"fmt"
	"time"
)

type OrderStatus string

const (
	StatusNew       OrderStatus = "new"
	StatusExpired   OrderStatus = "expired"
	StatusReturned  OrderStatus = "returned"
	StatusDelivered OrderStatus = "delivered"
)

var (
	ErrOrderAlreadyExists       = errors.New("order already exists")
	ErrInvalidStorageDate       = errors.New("invalid storage date")
	ErrOrderCannotBeReturned    = errors.New("order cannot be returned")
	ErrOrderCannotBeDelivered   = errors.New("order cannot be delivered")
	ErrStoragePeriodExpired     = errors.New("storage period expired")
	ErrStoragePeriodNotExpired  = errors.New("storage period not expired yet")
	ErrReturnPeriodExpired      = errors.New("return period expired")
	ErrOrderNotBelongToCustomer = errors.New("order does not belong to customer")
)

type Order struct {
	ID           uint        `json:"id"`
	CustomerID   uint        `json:"customer_id"`
	StorageUntil time.Time   `json:"storage_until"`
	Status       OrderStatus `json:"status"`
	DeliveredAt  time.Time   `json:"delivered_at,omitempty"`
	UpdatedAt    time.Time   `json:"updated_at"`
}

func (o Order) String() string {
	storageUntil := o.StorageUntil.Format(time.ANSIC)
	updatedAt := o.UpdatedAt.Format(time.ANSIC)
	deliveredAt := o.DeliveredAt.Format(time.ANSIC)

	// Если DeliveredAt нулевое, выводим "N/A"
	if o.DeliveredAt.IsZero() {
		deliveredAt = "N/A"
	}

	return fmt.Sprintf(
		"\nOrder ID: %d\nCustomer ID: %d\nStatus: %s\nStorage Until: %s\nUpdated At: %s\nDelivered At: %s\n",
		o.ID,
		o.CustomerID,
		o.Status,
		storageUntil,
		updatedAt,
		deliveredAt,
	)
}
